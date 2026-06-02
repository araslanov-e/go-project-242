package main

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCLIFlags(t *testing.T) {
	dir := createTestDir(t)
	file := filepath.Join(dir, "file49kb")
	binaryPath := buildBinary(t)

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
	dir := t.TempDir()

	stdout, stderr, err := runBinary(t, binaryPath, dir, "extra")

	require.Error(t, err)
	require.Empty(t, stdout)
	require.Contains(t, stderr, "too many arguments: expected 0 or 1")
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

func buildBinary(t *testing.T) string {
	t.Helper()

	rootDir := projectRoot(t)
	binaryPath := filepath.Join(t.TempDir(), "hexlet-path-size")
	cmd := exec.CommandContext(context.Background(), "go", "build", "-o", binaryPath, "./cmd/hexlet-path-size")
	cmd.Dir = rootDir

	output, err := cmd.CombinedOutput()
	require.NoError(t, err, string(output))

	return binaryPath
}

func runBinary(t *testing.T, binaryPath string, args ...string) (string, string, error) {
	t.Helper()

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd := exec.CommandContext(context.Background(), binaryPath, args...)
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
