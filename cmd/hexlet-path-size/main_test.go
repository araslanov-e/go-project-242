package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCLIFlags(t *testing.T) {
	binaryPath := buildBinary(t)

	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "file size in bytes",
			args:     []string{"testdata/file49kb"},
			expected: "50176B\ttestdata/file49kb\n",
		},
		{
			name:     "file size in human-readable format",
			args:     []string{"-H", "testdata/file49kb"},
			expected: "49.0KB\ttestdata/file49kb\n",
		},
		{
			name:     "directory size",
			args:     []string{"testdata"},
			expected: "102400B\ttestdata\n",
		},
		{
			name:     "recursive directory size",
			args:     []string{"-r", "testdata"},
			expected: "204800B\ttestdata\n",
		},
		{
			name:     "directory size with hidden files",
			args:     []string{"-a", "testdata"},
			expected: "128000B\ttestdata\n",
		},
		{
			name:     "recursive directory size with hidden files",
			args:     []string{"-r", "-a", "testdata"},
			expected: "230400B\ttestdata\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, err := runBinary(t, binaryPath, tt.args...)

			require.NoError(t, err)
			require.Empty(t, stderr)
			require.Equal(t, tt.expected, stdout)
		})
	}
}

func TestCLIWithoutPathUsesCurrentDirectory(t *testing.T) {
	binaryPath := buildBinary(t)

	stdout, stderr, err := runBinary(t, binaryPath)

	require.NoError(t, err)
	require.Empty(t, stderr)
	require.True(t, strings.HasSuffix(stdout, "\t.\n"))
}

func TestCLIWithTooManyArguments(t *testing.T) {
	binaryPath := buildBinary(t)

	stdout, stderr, err := runBinary(t, binaryPath, "testdata", "extra")

	require.Error(t, err)
	require.Empty(t, stdout)
	require.Contains(t, stderr, "too many arguments: expected 0 or 1")
}

func buildBinary(t *testing.T) string {
	t.Helper()

	rootDir := projectRoot(t)
	binaryPath := filepath.Join(t.TempDir(), "hexlet-path-size")
	cmd := exec.Command("go", "build", "-o", binaryPath, "./cmd/hexlet-path-size")
	cmd.Dir = rootDir

	output, err := cmd.CombinedOutput()
	require.NoError(t, err, string(output))

	return binaryPath
}

func runBinary(t *testing.T, binaryPath string, args ...string) (string, string, error) {
	t.Helper()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(binaryPath, args...)
	cmd.Dir = projectRoot(t)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func projectRoot(t *testing.T) string {
	t.Helper()

	rootDir, err := filepath.Abs("../..")
	require.NoError(t, err)

	return rootDir
}
