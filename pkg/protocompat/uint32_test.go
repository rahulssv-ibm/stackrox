//go:build test_all

package protocompat

import (
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stackrox/rox/pkg/protoassert"
)

func TestProtoUInt32Value(t *testing.T) {
	input1 := uint32(0)
	expectedVal1 := &types.UInt32Value{
		Value: input1,
	}

	val1 := ProtoUInt32Value(input1)
	protoassert.Equal(t, expectedVal1, val1)

	input2 := uint32(1234567890)
	expectedVal2 := &types.UInt32Value{
		Value: input2,
	}

	val2 := ProtoUInt32Value(input2)
	protoassert.Equal(t, expectedVal2, val2)
}
