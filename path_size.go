package code

import (
	"errors"
	"fmt"
	"os"
)

func GetPathSize(path string) (string, error) {
	if path == "" {
		return "", errors.New("path is empty")
	}

	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if info.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return "", err
		}

		var size int64
		for _, e := range entries {
			if e.IsDir() {
				continue
			}

			info, err := e.Info()
			if err != nil {
				return "", err
			}

			size += info.Size()
		}

		return formatSize(size), nil
	}

	return formatSize(info.Size()), nil
}

func formatSize(size int64) string {
	const (
		kb = 1024
		mb = kb * 1024
		gb = mb * 1024
	)

	switch {
	case size >= gb:
		return fmt.Sprintf("%.1fGB", float64(size)/gb)
	case size >= mb:
		return fmt.Sprintf("%.1fMB", float64(size)/mb)
	case size >= kb:
		return fmt.Sprintf("%.1fKB", float64(size)/kb)
	default:
		return fmt.Sprintf("%dB", size)
	}
}
