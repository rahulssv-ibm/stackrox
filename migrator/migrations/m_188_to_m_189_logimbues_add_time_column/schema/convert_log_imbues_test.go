package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestLogImbueSerialization(t *testing.T) {
	obj := &storage.LogImbue{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertLogImbueFromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertLogImbueToProto(m)
	assert.NoError(t, err)
	assert.Equal(t, obj, conv)
}
