package scannerv4

import (
	"fmt"
	"strings"

	v4 "github.com/stackrox/rox/generated/internalapi/scanner/v4"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/clair"
	"github.com/stackrox/rox/pkg/cvss"
	"github.com/stackrox/rox/pkg/cvss/cvssv2"
	"github.com/stackrox/rox/pkg/cvss/cvssv3"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/protocompat"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stackrox/rox/pkg/utils"
)

func imageScan(metadata *storage.ImageMetadata, report *v4.VulnerabilityReport) *storage.ImageScan {
	scan := &storage.ImageScan{
		// TODO(ROX-21362): Get ScannerVersion from ScannerV4 matcher API
		// ScannerVersion: ,
		ScanTime:        protocompat.TimestampNow(),
		OperatingSystem: os(report),
		Components:      components(metadata, report),
		Notes:           notes(report),
	}

	return scan
}

func components(metadata *storage.ImageMetadata, report *v4.VulnerabilityReport) []*storage.EmbeddedImageScanComponent {
	layerSHAToIndex := clair.BuildSHAToIndexMap(metadata)

	log.Debugf("ROSS METADATA: %+v", metadata)
	log.Debugf("ROSS METADATA layerSHAToIndex: %v", layerSHAToIndex)

	pkgs := report.GetContents().GetPackages()
	components := make([]*storage.EmbeddedImageScanComponent, 0, len(pkgs))
	for _, pkg := range pkgs {
		id := pkg.GetId()
		vulnIDs := report.GetPackageVulnerabilities()[id].GetValues()

		var (
			source   storage.SourceType
			location string
			layerIdx *storage.EmbeddedImageScanComponent_LayerIndex
		)
		env := environment(report, id)
		if env != nil {
			source, location = ParsePackageDB(env.GetPackageDb())
			layerIdx = layerIndex(layerSHAToIndex, env)
		}

		component := &storage.EmbeddedImageScanComponent{
			Name:          pkg.GetName(),
			Version:       pkg.GetVersion(),
			Vulns:         vulnerabilities(report.GetVulnerabilities(), vulnIDs),
			FixedBy:       pkg.GetFixedInVersion(),
			Source:        source,
			Location:      location,
			HasLayerIndex: layerIdx,
		}

		components = append(components, component)
	}

	return components
}

func environment(report *v4.VulnerabilityReport, id string) *v4.Environment {
	envList, ok := report.GetContents().GetEnvironments()[id]
	log.Debugf("ROSS ENVIRONMENT: %+v - %v", envList.GetEnvironments(), ok)
	if !ok {
		return nil
	}

	envs := envList.GetEnvironments()
	if len(envs) > 0 {
		// Just use the first environment.
		// It is possible there are multiple environments associated with this package;
		// however, for our purposes, we only need the first one,
		// as the layer index and package DB will always be the same between different environments.
		// Quay does this, too: https://github.com/quay/quay/blob/v3.10.3/data/secscan_model/secscan_v4_model.py#L583
		return envs[0]
	}

	return nil
}

// ParsePackageDB parses the given packageDB into its source type + filepath.
func ParsePackageDB(packageDB string) (storage.SourceType, string) {
	prefix, path, found := strings.Cut(packageDB, ":")
	if !found {
		// All currently know language packages have a prefix, so this must be an OS package.
		return storage.SourceType_OS, packageDB
	}

	switch prefix {
	case "go":
		return storage.SourceType_GO, path
	case "file", "jar", "maven":
		return storage.SourceType_JAVA, path
	case "nodejs":
		return storage.SourceType_NODEJS, path
	case "python":
		return storage.SourceType_PYTHON, path
	case "ruby":
		return storage.SourceType_RUBY, path
	case "bdb", "sqlite", "ndb":
		// RPM databases are prefixed with the DB kind.
		return storage.SourceType_OS, path
	default:
		// ":" is a valid character in a file path.
		// We could not identify a known prefix, so just return the entire path.
		return storage.SourceType_OS, packageDB
	}
}

func layerIndex(layerSHAToIndex map[string]int32, env *v4.Environment) *storage.EmbeddedImageScanComponent_LayerIndex {
	idx, ok := layerSHAToIndex[env.GetIntroducedIn()]
	if !ok {
		return nil
	}

	return &storage.EmbeddedImageScanComponent_LayerIndex{
		LayerIndex: idx,
	}
}

func vulnerabilities(vulnerabilities map[string]*v4.VulnerabilityReport_Vulnerability, ids []string) []*storage.EmbeddedVulnerability {
	if len(vulnerabilities) == 0 || len(ids) == 0 {
		return nil
	}

	vulns := make([]*storage.EmbeddedVulnerability, 0, len(ids))
	uniqueVulns := set.NewStringSet()
	for _, id := range ids {
		if !uniqueVulns.Add(id) {
			// Already saw this vulnerability, so ignore it.
			continue
		}

		ccVuln, ok := vulnerabilities[id]
		if !ok {
			log.Debugf("vuln ID %q from PackageVulnerabilities not found in Vulnerabilities, skipping", id)
			continue
		}

		// TODO(ROX-20355): Populate last modified once the API is available.
		vuln := &storage.EmbeddedVulnerability{
			Cve:         ccVuln.GetName(),
			Summary:     ccVuln.GetDescription(),
			Link:        link(ccVuln.GetLink()),
			PublishedOn: ccVuln.GetIssued(),
			// LastModified: ,
			VulnerabilityType: storage.EmbeddedVulnerability_IMAGE_VULNERABILITY,
			Severity:          normalizedSeverity(ccVuln.GetNormalizedSeverity()),
		}
		if err := setScoresAndScoreVersion(vuln, ccVuln.GetCvss()); err != nil {
			utils.Should(err)
		}
		maybeOverwriteSeverity(vuln)
		if ccVuln.GetFixedInVersion() != "" {
			vuln.SetFixedBy = &storage.EmbeddedVulnerability_FixedBy{
				FixedBy: ccVuln.GetFixedInVersion(),
			}
		}

		vulns = append(vulns, vuln)
	}

	return vulns
}

func setScoresAndScoreVersion(vuln *storage.EmbeddedVulnerability, vulnCVSS *v4.VulnerabilityReport_Vulnerability_CVSS) error {
	if vulnCVSS == nil {
		return nil
	}
	errList := errorhelpers.NewErrorList("failed to parse vector")
	if v2 := vulnCVSS.GetV2(); v2 != nil {
		if c, err := cvssv2.ParseCVSSV2(v2.GetVector()); err == nil {
			err = cvssv2.CalculateScores(c)
			if err != nil {
				errList.AddError(fmt.Errorf("calculating CVSS v2 scores: %w", err))
			}
			// Use the report's score if it exists.
			if baseScore := v2.GetBaseScore(); baseScore != 0.0 && baseScore != c.Score {
				log.Debugf("Calculated CVSSv2 score does not match given base score (%f != %f) for %s. Using given score...", c.Score, baseScore, vuln.GetCve())
				c.Score = baseScore
			}
			c.Severity = cvssv2.Severity(c.Score)
			vuln.CvssV2 = c
			// This sets the top level score for use in policies.
			// It will be overwritten if v3 exists.
			vuln.ScoreVersion = storage.EmbeddedVulnerability_V2
			vuln.Cvss = c.Score
		} else {
			errList.AddError(fmt.Errorf("v2: %w", err))
		}
	}
	if v3 := vulnCVSS.GetV3(); v3 != nil {
		if c, err := cvssv3.ParseCVSSV3(v3.GetVector()); err == nil {
			err = cvssv3.CalculateScores(c)
			if err != nil {
				errList.AddError(fmt.Errorf("calculating CVSS v3 scores: %w", err))
			}
			// Use the report's score if it exists.
			if baseScore := v3.GetBaseScore(); baseScore != 0.0 && baseScore != c.Score {
				log.Debugf("Calculated CVSSv3 score does not match given base score (%f != %f) for %s. Using given score...", c.Score, baseScore, vuln.GetCve())
				c.Score = baseScore
			}
			c.Severity = cvssv3.Severity(c.Score)
			vuln.CvssV3 = c
			// Overwrite v2, if set.
			vuln.ScoreVersion = storage.EmbeddedVulnerability_V3
			vuln.Cvss = c.Score
		} else {
			errList.AddError(fmt.Errorf("v3: %w", err))
		}
	}
	return errList.ToError()
}

// link returns the first link from space separated list of links (which is how ClairCore provides links).
// The ACS UI will fail to show a vulnerability's link if it is an invalid URL.
func link(links string) string {
	link, _, _ := strings.Cut(links, " ")
	return link
}

// maybeOverwriteSeverity overwrites vuln.Severity with one derived from the CVSS scores
// if vuln.Severity == storage.VulnerabilitySeverity_UNKNOWN_VULNERABILITY_SEVERITY.
func maybeOverwriteSeverity(vuln *storage.EmbeddedVulnerability) {
	vuln.Severity = cvss.VulnToSeverity(cvss.NewFromEmbeddedVulnerability(vuln))
}

func normalizedSeverity(severity v4.VulnerabilityReport_Vulnerability_Severity) storage.VulnerabilitySeverity {
	switch severity {
	case v4.VulnerabilityReport_Vulnerability_SEVERITY_LOW:
		return storage.VulnerabilitySeverity_LOW_VULNERABILITY_SEVERITY
	case v4.VulnerabilityReport_Vulnerability_SEVERITY_MODERATE:
		return storage.VulnerabilitySeverity_MODERATE_VULNERABILITY_SEVERITY
	case v4.VulnerabilityReport_Vulnerability_SEVERITY_IMPORTANT:
		return storage.VulnerabilitySeverity_IMPORTANT_VULNERABILITY_SEVERITY
	case v4.VulnerabilityReport_Vulnerability_SEVERITY_CRITICAL:
		return storage.VulnerabilitySeverity_CRITICAL_VULNERABILITY_SEVERITY
	default:
		return storage.VulnerabilitySeverity_UNKNOWN_VULNERABILITY_SEVERITY
	}
}

// os retrieves the OS name:version for the image represented by the given
// vulnerability report.
// If there are zero known distributions for the image or if there are multiple distributions,
// return "unknown", as StackRox only supports a single base-OS at this time.
func os(report *v4.VulnerabilityReport) string {
	dists := report.GetContents().GetDistributions()
	if len(dists) != 1 {
		return "unknown"
	}

	dist := dists[0]
	return dist.GetDid() + ":" + dist.GetVersionId()
}

func notes(report *v4.VulnerabilityReport) []storage.ImageScan_Note {
	notes := make([]storage.ImageScan_Note, 0, len(v4.VulnerabilityReport_Note_value))

	for _, note := range report.GetNotes() {
		switch note {
		case v4.VulnerabilityReport_NOTE_OS_UNSUPPORTED:
			notes = append(notes, storage.ImageScan_OS_CVES_UNAVAILABLE)
		case v4.VulnerabilityReport_NOTE_OS_UNKNOWN:
			notes = append(notes, storage.ImageScan_OS_UNAVAILABLE)
		default:
			// Ignore unknown/unsupported note.
		}
	}

	if len(notes) > 0 {
		notes = append(notes, storage.ImageScan_PARTIAL_SCAN_DATA)
	}

	return notes
}
