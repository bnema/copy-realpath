package path

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolver_Resolve(t *testing.T) {
	cwd, err := os.Getwd()
	require.NoError(t, err)

	resolver := NewResolver()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input returns current directory",
			input:    "",
			expected: cwd,
		},
		{
			name:     "dot returns current directory",
			input:    ".",
			expected: cwd,
		},
		{
			name:     "relative path is resolved against cwd",
			input:    "foo.txt",
			expected: filepath.Join(cwd, "foo.txt"),
		},
		{
			name:     "relative path with subdirectory",
			input:    "sub/foo.txt",
			expected: filepath.Join(cwd, "sub", "foo.txt"),
		},
		{
			name:     "parent directory reference is resolved",
			input:    "../foo.txt",
			expected: filepath.Clean(filepath.Join(cwd, "..", "foo.txt")),
		},
		{
			name:     "absolute path is returned as-is",
			input:    "/tmp/test.txt",
			expected: "/tmp/test.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := resolver.Resolve(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
