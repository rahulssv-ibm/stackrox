// Code originally generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"github.com/lib/pq"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/protocompat"
)

// ConvertPolicyFromProto converts a `*storage.Policy` to Gorm model
func ConvertPolicyFromProto(obj *storage.Policy) (*Policies, error) {
	serialized, err := obj.MarshalVT()
	if err != nil {
		return nil, err
	}
	model := &Policies{
		ID:                 obj.GetId(),
		Name:               obj.GetName(),
		Description:        obj.GetDescription(),
		Disabled:           obj.GetDisabled(),
		Categories:         pq.Array(obj.GetCategories()).(*pq.StringArray),
		LifecycleStages:    pq.Array(pgutils.ConvertEnumSliceToIntArray(obj.GetLifecycleStages())).(*pq.Int32Array),
		Severity:           obj.GetSeverity(),
		EnforcementActions: pq.Array(pgutils.ConvertEnumSliceToIntArray(obj.GetEnforcementActions())).(*pq.Int32Array),
		LastUpdated:        protocompat.NilOrTime(obj.GetLastUpdated()),
		SORTName:           obj.GetSORTName(),
		SORTLifecycleStage: obj.GetSORTLifecycleStage(),
		SORTEnforcement:    obj.GetSORTEnforcement(),
		Serialized:         serialized,
	}
	return model, nil
}

// ConvertPolicyToProto converts Gorm model `Policies` to its protobuf type object
func ConvertPolicyToProto(m *Policies) (*storage.Policy, error) {
	var msg storage.Policy
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
