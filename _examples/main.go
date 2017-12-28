package main

import (
	"errors"
	"fmt"

	"github.com/go-playground/retry"
)

func main() {
	err := retry.Run(5, func() (bool, error) {
		return false, errors.New("ERR")
	},
		func(attempt uint16, err error) {
			fmt.Printf("Attempt: %d Error: %s\n", attempt, err)
		},
	)
	if err != nil {
		panic(err)
	}
}
