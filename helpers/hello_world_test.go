package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
	result := HelloWorld("Bro")
	require.Equal(t, "Hello Viqi", result, "Error: result must be 'Hello Viqi'")
	fmt.Println("Unit test is done.")
}