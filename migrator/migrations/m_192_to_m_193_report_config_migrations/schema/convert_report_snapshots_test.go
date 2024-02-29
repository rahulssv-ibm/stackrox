package schema

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestReportSnapshotSerialization(t *testing.T) {
	obj := &storage.ReportSnapshot{}
	assert.NoError(t, testutils.FullInit(obj, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
	m, err := ConvertReportSnapshotFromProto(obj)
	assert.NoError(t, err)
	conv, err := ConvertReportSnapshotToProto(m)
	assert.NoError(t, err)
	assert.Equal(t, obj, conv)
}
