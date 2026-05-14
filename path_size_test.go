package code

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		target    string
		recursive bool
		human     bool
		all       bool
		expected  string
	}{
		{"testdata/file49kb", false, true, false, "49.0KB"},
		{"testdata/file49kb", false, false, false, fmt.Sprintf("%dB", 49*1024)},
		{"testdata", false, true, false, "100.0KB"},
		{"testdata", false, false, false, fmt.Sprintf("%dB", 100*1024)},
		{"testdata", false, false, true, fmt.Sprintf("%dB", 125*1024)},
		{"testdata", true, false, false, fmt.Sprintf("%dB", 200*1024)},
		{"testdata", true, false, true, fmt.Sprintf("%dB", 225*1024)},
	}

	for _, test := range tests {
		result, err := GetPathSize(test.target, test.recursive, test.human, test.all)
		a.Nil(err)
		a.Equal(test.expected, result)
	}

	result, err := GetPathSize("", false, false, false)
	a.EqualError(err, "path is empty")
	a.Empty(result)
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{"bytes", 512, "512B"},
		{"kilobytes", 49 * 1024, "49.0KB"},
		{"megabytes", 2 * 1024 * 1024, "2.0MB"},
		{"gigabytes", 3 * 1024 * 1024 * 1024, "3.0GB"},
		{"terabytes", 4 * 1024 * 1024 * 1024 * 1024, "4.0TB"},
		{"petabytes", 5 * 1024 * 1024 * 1024 * 1024 * 1024, "5.0PB"},
		{"exabytes", 6 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, "6.0EB"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, formatSize(test.size, true))
		})
	}

	var kilobytes int64 = 49 * 1024
	expected := "50176B"
	t.Run("unhuman", func(t *testing.T) {
		assert.Equal(t, expected, formatSize(kilobytes, false))
	})
}
