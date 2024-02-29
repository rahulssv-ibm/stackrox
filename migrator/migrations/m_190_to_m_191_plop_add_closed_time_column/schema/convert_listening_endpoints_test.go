package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestProcessListeningOnPortStorageSerialization(t *testing.T) {
	obj := &storage.ProcessListeningOnPortStorage{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertProcessListeningOnPortStorageFromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertProcessListeningOnPortStorageToProto(m)
	assert.NoError(t, err)
	assert.Equal(t, obj, conv)
}
