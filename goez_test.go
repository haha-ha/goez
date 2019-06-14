package goez

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	err := NewError("Test Error")
	str := ToString(err)
	fmt.Println(str)
	assert.True(t, strings.Index(str, "Test Error\ngoez.TestToString") == 0, "Exception [Test Error\ngoez.TestToString]", "got [", str, "]")
}
