// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/protocompat"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestTestChild1P4Serialization(t *testing.T) {
	obj := &storage.TestChild1P4{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertTestChild1P4FromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertTestChild1P4ToProto(m)
	assert.NoError(t, err)
	assert.True(t, protocompat.Equal(obj, conv))
}
