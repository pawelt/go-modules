package utils

import (
	"fmt"
	"github.com/pkg/errors"
)

func Triple(x int) int {
	return x * 3
}

func Fail1(msg string) error {
	return errors.Wrap(fmt.Errorf("Fail1(...)"), msg)
}
