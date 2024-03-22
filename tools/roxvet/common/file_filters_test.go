package common

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsGeneratedFile(t *testing.T) {
	tests := []struct {
		name string
		file string
		want bool
	}{
		{
			name: "top generated comment",
			file: `// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
package file1pkg

func funcInFile1() int {
	return 42
}
`,
			want: true,
		},
		{
			name: "generated comment before package",
			file: `// Lorem ipsum
// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
package file1pkg

func funcInFile1() int {
	return 42
}
`,
			want: true,
		},
		{
			name: "generated comment before package after blank (not ideal implementation)",
			file: `
// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
package file1pkg

func funcInFile1() int {
	return 42
}
`,
			want: true,
		},
		{
			name: "generated comment after package",
			file: `
package file1pkg
// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
func funcInFile1() int {
	return 42
}
`,
			want: false,
		},
		{
			name: "no generated comment",
			file: `
package file1pkg
func funcInFile1() int {
	return 42
}
`,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, "file.go", tt.file, parser.ParseComments)
			require.NoError(t, err)
			assert.Equalf(t, tt.want, IsGeneratedFile(file), "IsGeneratedFile(%s)", tt.file)
		})
	}
}
