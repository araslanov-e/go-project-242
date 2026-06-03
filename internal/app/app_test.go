package app

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunFlags(t *testing.T) {
	dir := createTestDir(t)
	file := filepath.Join(dir, "file49kb")

	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "file size in bytes",
			args:     []string{file},
			expected: "50176B\t" + file + "\n",
		},
		{
			name:     "file size in human-readable format",
			args:     []string{"-H", file},
			expected: "49.0KB\t" + file + "\n",
		},
		{
			name:     "directory size",
			args:     []string{dir},
			expected: "102400B\t" + dir + "\n",
		},
		{
			name:     "recursive directory size",
			args:     []string{"-r", dir},
			expected: "204800B\t" + dir + "\n",
		},
		{
			name:     "directory size with hidden files",
			args:     []string{"-a", dir},
			expected: "128000B\t" + dir + "\n",
		},
		{
			name:     "recursive directory size with hidden files",
			args:     []string{"-r", "-a", dir},
			expected: "230400B\t" + dir + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := runApp(t, tt.args...)

			require.NoError(t, err)
			require.Equal(t, tt.expected, output)
		})
	}
}

func TestRunWithoutPathUsesCurrentDirectory(t *testing.T) {
	output, err := runApp(t)

	require.NoError(t, err)
	require.True(t, strings.HasSuffix(output, "\t.\n"))
}

func TestRunWithTooManyArguments(t *testing.T) {
	dir := t.TempDir()

	output, err := runApp(t, dir, "extra")

	require.Error(t, err)
	require.Empty(t, output)
	require.Contains(t, err.Error(), "too many arguments: expected 0 or 1")
}

func runApp(t *testing.T, args ...string) (string, error) {
	t.Helper()

	var output bytes.Buffer

	runArgs := append([]string{"hexlet-path-size"}, args...)
	err := Run(context.Background(), runArgs, &output)

	return output.String(), err
}

func createTestDir(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	createFile(t, filepath.Join(dir, "file49kb"), 49*1024)
	createFile(t, filepath.Join(dir, "file51kb"), 51*1024)
	createFile(t, filepath.Join(dir, ".file25kb"), 25*1024)

	subdir := filepath.Join(dir, "subdir")
	require.NoError(t, os.Mkdir(subdir, 0o755))

	createFile(t, filepath.Join(subdir, "file49kb"), 49*1024)
	createFile(t, filepath.Join(subdir, "file51kb"), 51*1024)

	return dir
}

func createFile(t *testing.T, path string, size int) {
	t.Helper()

	require.NoError(t, os.WriteFile(path, make([]byte, size), 0o644))
}
