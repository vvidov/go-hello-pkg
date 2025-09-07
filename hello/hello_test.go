package hello

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
    // Arrange
    want := "hello wolrd"
    assert := require.New(t)

    // Act
    got := Hello()

    // Assert
    assert.Equal(want, got, "Hello() should return 'hello wolrd'")
}
