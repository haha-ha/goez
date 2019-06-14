package logez

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// Error log if the err is not nill
func Error(err error, args ...interface{}) {
	if err != nil {
		logger.Error("Error: ", err, ", detail: ", fmt.Sprint(args...))
	}
}

// Panic log and panic if the err is not nill
func Panic(err error, args ...interface{}) {
	if err != nil {
		logger.Panic("Error: ", err, ", detail: ", fmt.Sprint(args...))
	}
}
