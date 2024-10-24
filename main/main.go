package main

import (
	"fmt"
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
	stdout    *os.File
	calulator *calc.Addition
}

func NewHandler(stdout *os.File, calculator *calc.Addition) *Handler {
	return &Handler{
		stdout:    stdout,
		calulator: calculator,
	}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: calc [a] [b]")
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	calculator := &calc.Addition{}
	result := calculator.Calculate(a, b)

	_, err = fmt.Println(result)
	if err != nil {
		return err
	}

	return nil
}
