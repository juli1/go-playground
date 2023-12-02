package main

import (
	"errors"
	"fmt"
)

// One error
var err1 = errors.New("foo")

// Another one
var err2 = errors.New("bar")

// Another one different from the first one
var err3 = errors.New("foo")

func main() {
	println("errors")
	// equality because same value
	println("err1==err1", err1 == err1)
	// the rest is not true
	println("err1==err3", err1 == err3)
	println("err1 is err3", errors.Is(err1, err3))

	// wrapping the errors
	err4 := fmt.Errorf("another error: %w", err1)
	// false of course because value is different
	println("err1==err3", err1 == err4)

	// true - check the wrapping chain
	println("err4 is err1", errors.Is(err4, err1))
}
