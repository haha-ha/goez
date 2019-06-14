package goez

import (
	"fmt"
	"runtime/debug"
	"strings"
)

//ErrorEz an error with stack info
type ErrorEz struct {
	stack string
	msg   string
}

//Error implements error interface
func (e *ErrorEz) Error() string {
	return fmt.Sprintf("%v\n%v", e.msg, e.stack)
}

//NewError creates a new ErrorEz
func NewError(msg string) *ErrorEz {
	stack := string(debug.Stack())
	lines := strings.Split(stack, "\n")
	if len(lines) > 5 {
		lines = lines[5:]
	}
	stack = strings.Join(lines, "\n")
	return &ErrorEz{
		msg:   msg,
		stack: stack,
	}
}

//ToString convert any value to string
func ToString(a interface{}) string {
	return fmt.Sprintf("%+v", a)
}
