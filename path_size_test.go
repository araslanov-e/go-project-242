package code

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "file")
	createFile(t, path, 49*1024)

	tests := []struct {
		name      string
		target    string
		recursive bool
		human     bool
		all       bool
		expected  string
	}{
		{name: "unhuman", target: path, expected: fmt.Sprintf("%dB", 49*1024)},
		{name: "human", target: path, human: true, expected: "49.0KB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetPathSize(tt.target, tt.recursive, tt.human, tt.all)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetPathSize_FileNotExist(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "file")

	result, err := GetPathSize(path, false, false, false)
	require.Error(t, err)
	assert.Empty(t, result)
}

func createFile(t *testing.T, path string, size int) {
	t.Helper()

	require.NoError(t, os.WriteFile(path, make([]byte, size), 0o644))
}
