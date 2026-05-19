package code

import (
	"code/internal/sizefmt"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	if path == "" {
		return "", errors.New("path is empty")
	}

	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return sizefmt.Format(info.Size(), human), nil
	}

	size, err := getDirSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return sizefmt.Format(size, human), nil
}

func getDirSize(path string, recursive, all bool) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var size int64

	for _, e := range entries {
		if strings.HasPrefix(e.Name(), ".") && !all {
			continue
		}

		if e.IsDir() {
			if !recursive {
				continue
			}

			dirPath := filepath.Join(path, e.Name())
			dirSize, err := getDirSize(dirPath, recursive, all)
			if err != nil {
				return 0, err
			}

			size += dirSize
			continue
		}

		info, err := e.Info()
		if err != nil {
			return 0, err
		}

		size += info.Size()
	}

	return size, nil
}
