package code

import (
	"code/internal/pathsize"
	"code/internal/sizefmt"
	"fmt"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	options := pathsize.Options{Recursive: recursive, All: all}

	size, err := pathsize.Calculate(path, options)
	if err != nil {
		return "", fmt.Errorf("path size calculation %q: %w", path, err)
	}

	return sizefmt.Format(size, human), nil
}
