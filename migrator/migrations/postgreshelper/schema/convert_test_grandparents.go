// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"github.com/stackrox/rox/generated/storage"
)

// ConvertTestGrandparentFromProto converts a `*storage.TestGrandparent` to Gorm model
func ConvertTestGrandparentFromProto(obj *storage.TestGrandparent) (*TestGrandparents, error) {
	serialized, err := obj.MarshalVT()
	if err != nil {
		return nil, err
	}
	model := &TestGrandparents{
		ID:         obj.GetId(),
		Val:        obj.GetVal(),
		Priority:   obj.GetPriority(),
		RiskScore:  obj.GetRiskScore(),
		Serialized: serialized,
	}
	return model, nil
}

// ConvertTestGrandparent_EmbeddedFromProto converts a `*storage.TestGrandparent_Embedded` to Gorm model
func ConvertTestGrandparent_EmbeddedFromProto(obj *storage.TestGrandparent_Embedded, idx int, testGrandparentID string) (*TestGrandparentsEmbeddeds, error) {
	model := &TestGrandparentsEmbeddeds{
		TestGrandparentsID: testGrandparentID,
		Idx:                idx,
		Val:                obj.GetVal(),
	}
	return model, nil
}

// ConvertTestGrandparent_Embedded_Embedded2FromProto converts a `*storage.TestGrandparent_Embedded_Embedded2` to Gorm model
func ConvertTestGrandparent_Embedded_Embedded2FromProto(obj *storage.TestGrandparent_Embedded_Embedded2, idx int, testGrandparentID string, testGrandparentEmbeddedIdx int) (*TestGrandparentsEmbeddedsEmbedded2, error) {
	model := &TestGrandparentsEmbeddedsEmbedded2{
		TestGrandparentsID:           testGrandparentID,
		TestGrandparentsEmbeddedsIdx: testGrandparentEmbeddedIdx,
		Idx:                          idx,
		Val:                          obj.GetVal(),
	}
	return model, nil
}

// ConvertTestGrandparentToProto converts Gorm model `TestGrandparents` to its protobuf type object
func ConvertTestGrandparentToProto(m *TestGrandparents) (*storage.TestGrandparent, error) {
	var msg storage.TestGrandparent
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}
