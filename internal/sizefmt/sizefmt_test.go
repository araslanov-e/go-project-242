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
		{"kilobytes", 49 * KB, "49.0KB"},
		{"megabytes", 2 * MB, "2.0MB"},
		{"gigabytes", 3 * GB, "3.0GB"},
		{"terabytes", 4 * TB, "4.0TB"},
		{"petabytes", 5 * PB, "5.0PB"},
		{"exabytes", 6 * EB, "6.0EB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Format(tt.size, true))
		})
	}
}

func TestFormat_Unhuman(t *testing.T) {
	var kilobytes int64 = 49 * KB

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
