package hello

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	assert := require.New(t)
	testCases := []struct {
		name string
		want string
	}{
		{"World", "Hello, World"},
		{"Alice", "Hello, Alice"},
		{"Bob", "Hello, Bob"},
		{"", "Hello, "},
		{"Oliver", "Hello, Oli"},
	}
	for _, tc := range testCases {
		// Arrange
		name := tc.name
		want := tc.want

		// Act
		got := Hello(name)

		// Assert
		assert.Equal(want, got, "Hello(%q) should return %q", name, want)
	}
}
