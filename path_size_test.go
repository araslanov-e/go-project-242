package code

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDirPath = "testdata"

func TestGetPathSize_File(t *testing.T) {
	tests := []struct {
		name      string
		target    string
		recursive bool
		human     bool
		all       bool
		expected  string
	}{
		{name: "unhuman", target: testDirPath + "/file49kb", expected: fmt.Sprintf("%dB", 49*1024)},
		{name: "human", target: testDirPath + "/file49kb", human: true, expected: "49.0KB"},
		{name: "hidden file", target: testDirPath + "/.file25kb", expected: fmt.Sprintf("%dB", 25*1024)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetPathSize(tt.target, tt.recursive, tt.human, tt.all)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetPathSize_Directory(t *testing.T) {
	tests := []struct {
		name      string
		target    string
		recursive bool
		human     bool
		all       bool
		expected  string
	}{
		{name: "unhuman", target: testDirPath, expected: fmt.Sprintf("%dB", 100*1024)},
		{name: "human", target: testDirPath, human: true, expected: "100.0KB"},
		{name: "with hidden files", target: testDirPath, all: true, expected: fmt.Sprintf("%dB", 125*1024)},
		{name: "recursive", target: testDirPath, recursive: true, expected: fmt.Sprintf("%dB", 200*1024)},
		{
			name:      "with recursive and hidden files",
			target:    testDirPath,
			recursive: true,
			all:       true,
			expected:  fmt.Sprintf("%dB", 225*1024),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetPathSize(tt.target, tt.recursive, tt.human, tt.all)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetPathSize_FileNotExist(t *testing.T) {
	a := assert.New(t)

	result, err := GetPathSize(testDirPath+"/nonexistent", false, false, false)
	require.Error(t, err)
	a.Empty(result)
}

func TestGetPathSize_EmptyPath(t *testing.T) {
	a := assert.New(t)

	result, err := GetPathSize("", false, false, false)
	require.ErrorIs(t, err, ErrEmptyPath)
	a.Empty(result)
}
