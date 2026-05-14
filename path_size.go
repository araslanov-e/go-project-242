package code

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, resursive, human, all bool) (string, error) {
	if path == "" {
		return "", errors.New("path is empty")
	}

	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return formatSize(info.Size(), human), nil
	}

	size, err := getDirSize(path, resursive, all)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
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

func getDirSize(path string, recursive, all bool) (int64, error) {
	var size int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return size, err
	}

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
				return size, err
			}

			size += dirSize
		} else {
			info, err := e.Info()
			if err != nil {
				return size, err
			}

			size += info.Size()
		}
	}

	return size, nil
}
