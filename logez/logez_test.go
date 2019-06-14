package logez

import (
	"fmt"
	"goez"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	logger.Println("Hello")
}

func TestLogError(t *testing.T) {
	var err = goez.NewError("Test Error")
	Error(err, "Has Log")

	var writer strings.Builder
	defer logger.SetOutput(logger.Out)

	logger.SetOutput(&writer)
	Error(nil, "No Log")
	assert.Equal(t, "", writer.String(), "Expect No log")
	Error(err, "Has Log")
	assert.True(t, strings.Index(writer.String(), "Error: Test Error") >= 0, "Expect log")
}

func TestLogPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("OK: we got the panic:", err)
		} else {
			assert.Fail(t, "Expecting panic")
		}
	}()

	var err = goez.NewError("Test Error")

	var writer strings.Builder
	defer logger.SetOutput(logger.Out)

	logger.SetOutput(&writer)
	Panic(nil, "No Log")
	assert.Equal(t, "", writer.String(), "Expect No log")
	Panic(err, "Has Log")
	assert.Fail(t, "Expecting panic. This line should not be excuted!")
}
