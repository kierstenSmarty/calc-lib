package main

import (
	"errors"
	"testing"
)

func assertError(t *testing.T, actual, target error) {
	t.Helper()
	if !errors.Is(actual, target) {
		t.Errorf("expected: %v, got: %v", target, actual)
	}
}

func TestHandler_WrongNumberOfArguments(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	assertError(t, err, errWrongNumberOfArgs)
}
func TestHandler_InvalidFirstArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"INVALID", "1"})
	assertError(t, err, errInvalidArgument)
}
func TestHandler_InvalidSecondArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"1", "INVALID"})
	assertError(t, err, errInvalidArgument)
}
