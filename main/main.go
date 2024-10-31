package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/kierstenSmarty/calc-lib"
)

func main() {
	handler := NewHandler(os.Stdout, &calc.Addition{})

	args := os.Args[1:]
	err := handler.Handle(args)

	if err != nil {
		log.Fatal(err)
	}
}

type Handler struct {
	stdout    io.Writer
	calulator *calc.Addition
}

func NewHandler(stdout io.Writer, calculator *calc.Addition) *Handler {
	return &Handler{
		stdout:    stdout,
		calulator: calculator,
	}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongNumberOfArgs
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %w", errInvalidArgument, err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: %w", errInvalidArgument, err)
	}

	calculator := &calc.Addition{}
	result := calculator.Calculate(a, b)

	_, err = fmt.Println(result)
	if err != nil {
		return err
	}

	return nil
}

var errWrongNumberOfArgs = errors.New("usage: calc [a] [b]")
var errInvalidArgument = errors.New("invalid argument")
var errWriterFailure = errors.New("writer failure")
