// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/protoassert"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestTestChild1Serialization(t *testing.T) {
	obj := &storage.TestChild1{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertTestChild1FromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertTestChild1ToProto(m)
	assert.NoError(t, err)
	protoassert.Equal(t, obj, conv)
}
