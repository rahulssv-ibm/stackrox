package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	v4 "github.com/stackrox/rox/generated/internalapi/scanner/v4"
	"github.com/stackrox/rox/pkg/scannerv4/client"
)

func nodeScanCmd(ctx context.Context) *cobra.Command {
	cmd := cobra.Command{
		Use:   "nodescan",
		Short: "Triggers a node scan for the node the target node indexer runs on",
	}
	flags := cmd.PersistentFlags()
	withMatching := flags.Bool(
		"match",
		false,
		"Additionally match vulnerabilities after scanning")
	withSummary := flags.Bool(
		"summary",
		false,
		"Print report summary in addition to full report")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		// Create scanner client.
		scanner, err := factory.Create(ctx)
		if err != nil {
			return fmt.Errorf("create client: %w", err)
		}

		var report any
		report, err = createReport(ctx, scanner, *withMatching)
		if err != nil {
			return fmt.Errorf("scanning: %w", err)
		}

		reportJSON, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			return fmt.Errorf("decoding report: %w", err)
		}

		fmt.Println(string(reportJSON))

		if !*withSummary {
			return nil
		}

		switch r := report.(type) {
		case *v4.VulnerabilityReport:
			fmt.Printf("Vulnerability Report contains %d packages with %d vulnerabilities\n",
				len(r.GetContents().GetPackages()),
				len(r.GetVulnerabilities()))
		case *v4.IndexReport:
			fmt.Printf("Index Report contains %d packages\n", len(r.GetContents().GetPackages()))
		}

		return nil
	}
	return &cmd
}

func createReport(ctx context.Context, scanner client.Scanner, withMatching bool) (report any, err error) {
	if withMatching {
		var vr *v4.VulnerabilityReport
		vr, err = scanner.IndexAndScanNode(ctx)
		report = vr
	} else {
		var ir *v4.IndexReport
		ir, err = scanner.CreateNodeIndexReport(ctx)
		report = ir
	}
	if err != nil {
		return nil, err
	}

	return report, nil
}
