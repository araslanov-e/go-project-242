package pathsize

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculate_File(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "file")
	createFile(t, path, 49*1024)

	size, err := Calculate(path, Options{})

	require.NoError(t, err)
	assert.Equal(t, int64(49*1024), size)
}

func TestCalculate_HiddenFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, ".file")
	createFile(t, path, 49*1024)

	size, err := Calculate(path, Options{})

	require.NoError(t, err)
	assert.Equal(t, int64(49*1024), size)
}

func TestCalculate_EmptyPath(t *testing.T) {
	path := ""

	result, err := Calculate(path, Options{})
	require.ErrorIs(t, err, ErrEmptyPath)
	assert.Empty(t, result)
}

func TestCalculate_Directory(t *testing.T) {
	dir := createTestDir(t)

	tests := []struct {
		name     string
		options  Options
		expected int64
	}{
		{name: "without recursive and hidden files", expected: 100 * 1024},
		{name: "with hidden files", options: Options{All: true}, expected: 125 * 1024},
		{name: "recursive", options: Options{Recursive: true}, expected: 200 * 1024},
		{
			name:     "recursive with hidden files",
			options:  Options{Recursive: true, All: true},
			expected: 225 * 1024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size, err := Calculate(dir, tt.options)

			require.NoError(t, err)
			assert.Equal(t, tt.expected, size)
		})
	}
}

func TestCalculate_SymlinkPath(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "file")
	link := filepath.Join(dir, "link")

	createFile(t, target, 49*1024)
	createSymlink(t, target, link)

	size, err := Calculate(link, Options{})

	require.NoError(t, err)
	assert.Empty(t, size)
}

func TestCalculate_DirectoryWithSymlinkFile(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "file")
	link := filepath.Join(dir, "link")

	createFile(t, target, 49*1024)
	createSymlink(t, target, link)

	size, err := Calculate(dir, Options{})

	require.NoError(t, err)
	assert.Equal(t, int64(49*1024), size)
}

func TestCalculate_DirectoryWithSymlinkDirectory(t *testing.T) {
	dir := t.TempDir()
	subdir := filepath.Join(dir, "subdir")
	link := filepath.Join(dir, "link")

	require.NoError(t, os.Mkdir(subdir, 0o755))
	createFile(t, filepath.Join(subdir, "file51kb"), 51*1024)
	createSymlink(t, subdir, link)

	size, err := Calculate(dir, Options{Recursive: true})

	require.NoError(t, err)
	assert.Equal(t, int64(51*1024), size)
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

func createSymlink(t *testing.T, oldname, newname string) {
	t.Helper()

	if err := os.Symlink(oldname, newname); err != nil {
		t.Skipf("symlink is not supported: %v", err)
	}
}
