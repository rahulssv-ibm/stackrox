package dispatchers

import (
	"github.com/ComplianceAsCode/compliance-operator/pkg/apis/compliance/v1alpha1"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/centralsensor"
	"github.com/stackrox/rox/sensor/common/centralcaps"
	"github.com/stackrox/rox/sensor/kubernetes/eventpipeline/component"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// SuitesDispatcher handles compliance operator suites
type SuitesDispatcher struct{}

// NewSuitesDispatcher creates and returns a new compliance suite dispatcher.
func NewSuitesDispatcher() *SuitesDispatcher {
	return &SuitesDispatcher{}
}

// ProcessEvent processes a suite event
func (c *SuitesDispatcher) ProcessEvent(obj, _ interface{}, action central.ResourceAction) *component.ResourceEvent {
	var complianceSuite v1alpha1.ComplianceSuite

	unstructuredObject, ok := obj.(*unstructured.Unstructured)
	if !ok {
		log.Errorf("Not of type 'unstructured': %T", obj)
		return nil
	}

	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObject.Object, &complianceSuite); err != nil {
		log.Errorf("error converting unstructured to compliance suite: %v", err)
		return nil
	}
	id := string(complianceSuite.UID)
	// We are pulling additional data for suites and using the storage object even in an internal api
	// is a bad practice, so we will make that split now.  V1 and  compliance will both need to work for a period
	// of time.  However, we should not need to send the same suite twice, the pipeline can convert the  sensor message
	// so V1 and  objects can both be stored.

	events := []*central.SensorEvent{
		{
			Id:     id,
			Action: action,
			Resource: &central.SensorEvent_Comp.SensorEvent_ComplianceOperatorSuite{
				ComplianceOperatorSuite: &storage.ComplianceOperatorSuite{
					Id:          id,
					SuiteId:     complianceSuite.ID,
					Name:        complianceSuite.Name,
					Title:       complianceSuite.Title,
					Labels:      complianceSuite.Labels,
					Annotations: complianceSuite.Annotations,
					Description: complianceSuite.Description,
					Rationale:   complianceSuite.Rationale,
				},
			},
		},
	}

	if centralcaps.Has(centralsensor.ComplianceIntegrations) {
		fixes := make([]*central.ComplianceOperatorSuite_Fix, 0, len(complianceSuite.AvailableFixes))
		for _, r := range complianceSuite.AvailableFixes {
			fixes = append(fixes, &central.ComplianceOperatorSuite_Fix{
				Platform:   r.Platform,
				Disruption: r.Disruption,
			})
		}

		events = append(events, &central.SensorEvent{
			Id:     id,
			Action: action,
			Resource: &central.SensorEvent_ComplianceOperatorSuite{
				ComplianceOperatorSuite: &central.ComplianceOperatorSuite{
					SuiteId:     complianceSuite.ID,
					Id:          id,
					Name:        complianceSuite.Name,
					SuiteType:   complianceSuite.CheckType,
					Severity:    suiteSeverityToSeverity(complianceSuite.Severity),
					Labels:      complianceSuite.Labels,
					Annotations: complianceSuite.Annotations,
					Title:       complianceSuite.Title,
					Description: complianceSuite.Description,
					Rationale:   complianceSuite.Rationale,
					Fixes:       fixes,
					Warning:     complianceSuite.Warning,
				},
			},
		})
	}

	return component.NewEvent(events...)
}
