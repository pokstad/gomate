package gomate

import (
	"errors"
	"fmt"
	"runtime"
)

type e struct {
	cause error    // originating error
	calls []string // call stack
}

// E returns a new gomate error with the annotation
func E(annotation string) error {
	ee := e{
		cause: errors.New(annotation),
	}

	recordCall(&ee.calls, annotation)

	return ee

}

// PushE will take an existing error and push the call stack
func PushE(cause error, annotation string) error {
	ee, ok := cause.(e)
	if !ok {
		ee = e{
			cause: cause,
		}
	}

	recordCall(&ee.calls, annotation)
	return ee
}

// Cause returns the original error
func Cause(err error) error {
	ee, ok := err.(e)
	if ok {
		return ee.cause
	}
	return err
}

// TODO: move recordCall to tagged debug and release source files where release
// is a no op
func recordCall(calls *[]string, msg string) {
	if msg == "" {
		msg = "DOH"
	}

	_, file, line, ok := runtime.Caller(2)
	if ok {
		*calls = append(
			*calls,
			fmt.Sprintf(
				"%s: file: %s on line %d",
				msg, file, line,
			),
		)
	}
}

// Error string for cause
func (ee e) Error() string {
	if ee.cause != nil {
		return ee.cause.Error()
	}
	return ""
}
