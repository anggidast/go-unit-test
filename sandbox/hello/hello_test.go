package sandbox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// do something before the test run
	fmt.Println("before unit test")

	m.Run() // this line will execute all unit tests in the package

	// do something after the test run
	fmt.Println("after unit test")
}

func TestSubTest(t *testing.T) {
	t.Run("Anggi", func(t *testing.T) {
		result := Hello("Anggi")
		require.Equal(t, "Hello Anggi", result, "result must be 'Hello Anggi'")
	})

	t.Run("Dastariana", func(t *testing.T) {
		result := Hello("Dastariana")
		require.Equal(t, "Hello Dastariana", result, "result must be 'Hello Dastariana'")
	})
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Anggi",
			request:  "Anggi",
			expected: "Hello Anggi",
		},
		{
			name:     "Dastariana",
			request:  "Dastariana",
			expected: "Hello Dastariana",
		},
		{
			name:     "Kuang",
			request:  "Kuang",
			expected: "Hello Kuang",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Hello(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestHelloAnggi(t *testing.T) {
	result := Hello("Anggi")
	// if result != "Hello Anggi" {
	// t.Fail()
	// t.Error("the result is not 'Hello Anggi'")
	// }

	assert.Equal(t, "Hello Angg", result, "result must be 'Hello Angg'")
}

func TestHelloDastariana(t *testing.T) {
	result := Hello("Dastariana")
	if result != "Hello Dastariana" {
		// t.FailNow()
		t.Fatal("the result is not 'Hello Dastariana'")
	}
}
