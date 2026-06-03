package pathsize

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	Recursive bool
	All       bool
}

var ErrEmptyPath = errors.New("path is empty")

func Calculate(path string, options Options) (int64, error) {
	if isEmpty(path) {
		return 0, ErrEmptyPath
	}

	info, err := getPathInfo(path)
	if err != nil {
		return 0, fmt.Errorf("path info %q: %w", path, err)
	}

	if isSymlink(info.Mode()) {
		return 0, nil
	}

	if !info.IsDir() {
		return calculateFileSize(info), nil
	}

	size, err := calculateDirSize(path, options)
	if err != nil {
		return 0, fmt.Errorf("calculate directory size %q: %w", path, err)
	}

	return size, nil
}

func getPathInfo(path string) (os.FileInfo, error) {
	return os.Lstat(path)
}

func calculateFileSize(info os.FileInfo) int64 {
	return info.Size()
}

func calculateDirSize(path string, options Options) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("read directory %q: %w", path, err)
	}

	var size int64

	for _, entry := range entries {
		entrySize, err := calculateEntrySize(path, entry, options)
		if err != nil {
			return 0, fmt.Errorf("calculate entries size %q: %w", path, err)
		}

		size += entrySize
	}

	return size, nil
}

func calculateEntrySize(dir string, entry os.DirEntry, options Options) (int64, error) {
	if shouldSkipEntry(entry, options) {
		return 0, nil
	}

	if entry.IsDir() {
		return calculateSubdirSize(dir, entry, options)
	}

	info, err := entry.Info()
	if err != nil {
		return 0, fmt.Errorf("get info for entry %q: %w", entry.Name(), err)
	}

	return calculateFileSize(info), nil
}

func calculateSubdirSize(dir string, entry os.DirEntry, options Options) (int64, error) {
	if !options.Recursive {
		return 0, nil
	}

	path := filepath.Join(dir, entry.Name())

	size, err := calculateDirSize(path, options)
	if err != nil {
		return 0, fmt.Errorf("calculate subdirectory size %q: %w", path, err)
	}

	return size, nil
}

func shouldSkipEntry(entry os.DirEntry, options Options) bool {
	return isSymlink(entry.Type()) || (!options.All && isHidden(entry.Name()))
}

func isEmpty(path string) bool {
	return path == ""
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

func isSymlink(mode os.FileMode) bool {
	return mode&os.ModeSymlink != 0
}
