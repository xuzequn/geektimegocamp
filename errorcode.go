package main

import "errors"

type ErrorType struct {
	message string
}

type ErrorTypeint interface {
	error()
}

func (err *ErrorType) error() error {
	ERRORNOTFOUD := errors.New("Not Found")
	return ERRORNOTFOUD
}
