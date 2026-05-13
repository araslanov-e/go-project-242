package code

import (
	"errors"
	"fmt"
	"os"
)

func GetPathSize(path string, human bool) (string, error) {
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

		return formatSize(size, human), nil
	}

	return formatSize(info.Size(), human), nil
}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	const (
		kb = 1024
		mb = kb * 1024
		gb = mb * 1024
		tb = gb * 1024
		pb = tb * 1024
		eb = pb * 1024
	)

	switch {
	case size >= eb:
		return fmt.Sprintf("%.1fEB", float64(size)/eb)
	case size >= pb:
		return fmt.Sprintf("%.1fPB", float64(size)/pb)
	case size >= tb:
		return fmt.Sprintf("%.1fTB", float64(size)/tb)
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
