package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go test -v -run=TestHelloWorldAssert
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Michi")
	assert.Equal(t, "Hello Michi", result, "Result must be 'Hello Michi")
	fmt.Println("Done")
}

// go test -v -run=TestHelloWorldRequire
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Michi")
	require.Equal(t, "Hello Michi", result, "Result must be 'Hello Michi")
	fmt.Println("Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on Linux")
	}
	result := HelloWorld("Michi")
	require.Equal(t, "Hello Michi", result, "Result must be 'Hello Michi'")
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Michi",
			request:  "Michi",
			expected: "Hello Michi",
		},
		{
			name:     "Nico",
			request:  "Nico",
			expected: "Hello Nico",
		},
		{
			name:     "Uwe",
			request:  "Uwe",
			expected: "Hello Uwe",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
