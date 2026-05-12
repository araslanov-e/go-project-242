package code

import (
	"errors"
	"os"
)

func GetPathSize(path string) (int64, error) {
	if path == "" {
		return 0, errors.New("path is empty")
	}

	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if info.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}

		var size int64
		for _, e := range entries {
			if e.IsDir() {
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

	return info.Size(), nil
}
