package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Gibs")
	if result != "Hello Gibs" {
		// error
		panic("Result is not Hello Gibs")
	}
}

func TestHelloWorldPanic(t *testing.T) {
	result := HelloWorld("Gibran")
	if result != "Hello Gibran" {
		// error
		panic("Result is not Hello Gibs")
	}
}

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("Gibs")
	if result != "Hello Gibran" {
		// error
		t.Fail()
	}
	fmt.Println("ini unit test Fail")
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("Dika")
	if result != "Hello Gibran" {
		// error
		t.FailNow()
	}
	fmt.Println("ini unit test FailNow")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Error")
	if result != "Hello Gibran" {
		// error
		t.Error("Testing Error, error")
	}
	fmt.Println("ini unit test Error")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Fatal")
	if result != "Hello Gibran" {
		// error
		t.Fatal("Testing Fatal, error")
	}
	fmt.Println("ini unit test Fatal")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fatal")
	assert.Equal(t, "Hello Gibran", result, "Result must be Hello Gibran")
	fmt.Println("TestHelloWorldAssert")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Gibran", result, "result must be Hello Gibran")
	fmt.Println("TestHelloWorldRequire")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		result := HelloWorld("Test1")
		require.Equal(t, "Hello Gibran", result, "TestSubTest1")
	})

	t.Run("Test2", func(t *testing.T) {
		result := HelloWorld("Test2")
		require.Equal(t, "Hello Test2", result, "TestSubTest2")
	})
}

func TableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Gibs",
			request:  "Gibs",
			expected: "Hello Gibs",
		},
		{
			name:     "Gibran",
			request:  "Gibran",
			expected: "Hello Gibran",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
