package helpers

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run in Mac OS")
	}

	result := HelloWorld("Bro")
	assert.Equal(t, "Hello Bro", result, "Error: result must be 'Hello Viqi'")
	fmt.Println("Unit test is done.")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Viqi")

	if result != "Hello Viqi" {
		// error
		t.Error("Error: result must be 'Hello Viqi'")
	}

	fmt.Println("Hello World Test is done.")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Bro")
	assert.Equal(t, "Hello Bro", result, "Error: result must be 'Hello Viqi'")
	fmt.Println("Unit test is done.")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Viqi")
	require.Equal(t, "Hello Viqi", result, "Error: result must be 'Hello Viqi'")
	fmt.Println("Unit test is done.")
}

func TestSubTest(t *testing.T) {
	t.Run("Viqi", func(t *testing.T) {
		result := HelloWorld("Viqi")
		assert.Equal(t, "Hello Viqi", result, "Error: result must be 'Hello Viqi'")
	})
	t.Run("Nurhaqiqi", func(t *testing.T) {
		result := HelloWorld("Nurhaqiqi")
		assert.Equal(t, "Hello Nurhaqiqi", result, "Error: result must be 'Hello Nurhaqiqi'")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Viqi",
			request: "Viqi",
			expected: "Hello Viqi",
		},
		{
			name: "Nurhaqiqi",
			request: "Nurhaqiqi",
			expected: "Hello Nurhaqiqi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
