package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestResourceCollectionSerialization(t *testing.T) {
	obj := &storage.ResourceCollection{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertResourceCollectionFromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertResourceCollectionToProto(m)
	assert.NoError(t, err)
	assert.Equal(t, obj, conv)
}
