package sizefmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat_Human(t *testing.T) {
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
			assert.Equal(t, test.expected, Format(test.size, true))
		})
	}
}

func TestFormat_Unhuman(t *testing.T) {
	var kilobytes int64 = 49 * 1024
	expected := "50176B"
	t.Run("unhuman", func(t *testing.T) {
		assert.Equal(t, expected, Format(kilobytes, false))
	})
}

func TestFormat_Negative(t *testing.T) {
	expected := "0B"
	t.Run("negative", func(t *testing.T) {
		assert.Equal(t, expected, Format(-1, false))
	})
}
