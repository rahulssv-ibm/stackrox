// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/protocompat"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestPolicySerialization(t *testing.T) {
	obj := &storage.Policy{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertPolicyFromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertPolicyToProto(m)
	assert.NoError(t, err)
	assert.True(t, protocompat.Equal(obj, conv))
}
