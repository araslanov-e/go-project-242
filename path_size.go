package code

import (
	"code/internal/sizefmt"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var ErrEmptyPath = errors.New("path is empty")

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	if path == "" {
		return "", ErrEmptyPath
	}

	info, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("lstat path %q: %w", path, err)
	}

	if !info.IsDir() {
		return sizefmt.Format(info.Size(), human), nil
	}

	size, err := getDirSize(path, recursive, all)
	if err != nil {
		return "", fmt.Errorf("get directory size %q: %w", path, err)
	}

	return sizefmt.Format(size, human), nil
}

func getDirSize(path string, recursive, all bool) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("read directory %q: %w", path, err)
	}

	var size int64

	for _, e := range entries {
		if !all && isHiddenEntry(e.Name()) {
			continue
		}

		if e.IsDir() {
			if !recursive {
				continue
			}

			dirPath := filepath.Join(path, e.Name())

			dirSize, err := getDirSize(dirPath, recursive, all)
			if err != nil {
				return 0, fmt.Errorf("get subdirectory size %q: %w", dirPath, err)
			}

			size += dirSize

			continue
		}

		info, err := e.Info()
		if err != nil {
			return 0, fmt.Errorf("get info for entry %q: %w", e.Name(), err)
		}

		size += info.Size()
	}

	return size, nil
}

func isHiddenEntry(name string) bool {
	return strings.HasPrefix(name, ".")
}
