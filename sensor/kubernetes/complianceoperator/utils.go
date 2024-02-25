package complianceoperator

import (
	stdErrors "errors"

	"github.com/ComplianceAsCode/compliance-operator/pkg/apis/compliance/v1alpha1"
	"github.com/adhocore/gronx"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/pkg/complianceoperator"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/pointers"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type scanNameGetter interface {
	GetScanName() string
}

func validateScanName(req scanNameGetter) error {
	if req == nil {
		return errors.New("apply scan configuration request is empty")
	}
	if req.GetScanName() == "" {
		return errors.New("no name provided for the scan")
	}
	return nil
}

func convertCentralRequestToScanSetting(namespace string, request *central.ApplyComplianceScanConfigRequest_BaseScanSettings, cron string) *v1alpha1.ScanSetting {
	return &v1alpha1.ScanSetting{
		TypeMeta: v1.TypeMeta{
			Kind:       complianceoperator.ScanSetting.Kind,
			APIVersion: complianceoperator.GetGroupVersion().String(),
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      request.GetScanName(),
			Namespace: namespace,
			Labels: map[string]string{
				"app.kubernetes.io/name": "stackrox",
			},
			Annotations: map[string]string{
				"owner": "stackrox",
			},
		},
		Roles: []string{masterRole, workerRole},
		ComplianceSuiteSettings: v1alpha1.ComplianceSuiteSettings{
			AutoApplyRemediations:  false,
			AutoUpdateRemediations: false,
			Schedule:               cron,
		},
		ComplianceScanSettings: v1alpha1.ComplianceScanSettings{
			StrictNodeScan:    pointers.Bool(false),
			ShowNotApplicable: false,
			Timeout:           env.ComplianceScanTimeout.Setting(),
			MaxRetryOnTimeout: env.ComplianceScanRetries.IntegerSetting(),
		},
	}
}

func convertCentralRequestToScanSettingBinding(namespace string, request *central.ApplyComplianceScanConfigRequest_BaseScanSettings) *v1alpha1.ScanSettingBinding {
	profileRefs := make([]v1alpha1.NamedObjectReference, 0, len(request.GetProfiles()))
	for _, profile := range request.GetProfiles() {
		profileRefs = append(profileRefs, v1alpha1.NamedObjectReference{
			Name:     profile,
			Kind:     complianceoperator.Profile.Kind,
			APIGroup: complianceoperator.GetGroupVersion().String(),
		})
	}

	return &v1alpha1.ScanSettingBinding{
		TypeMeta: v1.TypeMeta{
			Kind:       complianceoperator.ScanSettingBinding.Kind,
			APIVersion: complianceoperator.GetGroupVersion().String(),
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      request.GetScanName(),
			Namespace: namespace,
			Labels: map[string]string{
				"app.kubernetes.io/name": "stackrox",
			},
			Annotations: map[string]string{
				"owner": "stackrox",
			},
		},
		Profiles: profileRefs,
		SettingsRef: &v1alpha1.NamedObjectReference{
			Name:     request.GetScanName(),
			Kind:     complianceoperator.ScanSetting.Kind,
			APIGroup: complianceoperator.GetGroupVersion().String(),
		},
	}
}

func updateScanSettingFromCentralRequest(scanSetting *v1alpha1.ScanSetting, request *central.ApplyComplianceScanConfigRequest_UpdateScheduledScan) *v1alpha1.ScanSetting {
	// TODO:  Update additional fields as ACS capability expands
	scanSetting.Roles = []string{masterRole, workerRole}
	scanSetting.ComplianceSuiteSettings = v1alpha1.ComplianceSuiteSettings{
		AutoApplyRemediations:  false,
		AutoUpdateRemediations: false,
		Schedule:               request.GetCron(),
	}
	scanSetting.ComplianceScanSettings = v1alpha1.ComplianceScanSettings{
		StrictNodeScan:    pointers.Bool(false),
		ShowNotApplicable: false,
		Timeout:           env.ComplianceScanTimeout.Setting(),
		MaxRetryOnTimeout: env.ComplianceScanRetries.IntegerSetting(),
	}

	return scanSetting
}

func updateScanSettingBindingFromCentralRequest(scanSettingBinding *v1alpha1.ScanSettingBinding, request *central.ApplyComplianceScanConfigRequest_BaseScanSettings) *v1alpha1.ScanSettingBinding {
	profileRefs := make([]v1alpha1.NamedObjectReference, 0, len(request.GetProfiles()))
	for _, profile := range request.GetProfiles() {
		profileRefs = append(profileRefs, v1alpha1.NamedObjectReference{
			Name:     profile,
			Kind:     complianceoperator.Profile.Kind,
			APIGroup: complianceoperator.GetGroupVersion().String(),
		})
	}

	// TODO:  Update additional fields as ACS capability expands
	scanSettingBinding.Profiles = profileRefs

	return scanSettingBinding
}

func validateApplyScheduledScanConfigRequest(req *central.ApplyComplianceScanConfigRequest_ScheduledScan) error {
	if req == nil {
		return errors.New("apply scan configuration request is empty")
	}
	var validationErrs error
	if req.GetScanSettings().GetScanName() == "" {
		validationErrs = stdErrors.Join(validationErrs, errors.New("no name provided for the scan"))
	}
	if len(req.GetScanSettings().GetProfiles()) == 0 {
		validationErrs = stdErrors.Join(validationErrs, errors.New("compliance profiles not specified"))
	}
	if req.GetCron() != "" {
		cron := gronx.New()
		if !cron.IsValid(req.GetCron()) {
			validationErrs = stdErrors.Join(validationErrs, errors.New("schedule is not valid"))
		}
	}
	return validationErrs
}

func validateUpdateScheduledScanConfigRequest(req *central.ApplyComplianceScanConfigRequest_UpdateScheduledScan) error {
	if req == nil {
		return errors.New("update scan configuration request is empty")
	}
	var validationErrs error
	if req.GetScanSettings().GetScanName() == "" {
		validationErrs = stdErrors.Join(validationErrs, errors.New("no name provided for the scan"))
	}
	if len(req.GetScanSettings().GetProfiles()) == 0 {
		validationErrs = stdErrors.Join(validationErrs, errors.New("compliance profiles not specified"))
	}
	if req.GetCron() != "" {
		cron := gronx.New()
		if !cron.IsValid(req.GetCron()) {
			validationErrs = stdErrors.Join(validationErrs, errors.New("schedule is not valid"))
		}
	}
	return validationErrs
}

func validateApplySuspendScheduledScanRequest(req *central.ApplyComplianceScanConfigRequest_SuspendScheduledScan) error {
	return validateScanName(req)
}

func validateApplyResumeScheduledScanRequest(req *central.ApplyComplianceScanConfigRequest_ResumeScheduledScan) error {
	return validateScanName(req)
}

func validateApplyRerunScheduledScanRequest(req *central.ApplyComplianceScanConfigRequest_RerunScheduledScan) error {
	if req == nil {
		return errors.New("apply scan configuration request is empty")
	}
	if req.GetScanName() == "" {
		return errors.New("no name provided for the scan")
	}
	return nil
}

func runtimeObjToUnstructured(obj runtime.Object) (*unstructured.Unstructured, error) {
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	return &unstructured.Unstructured{
		Object: unstructuredObj,
	}, nil
}
