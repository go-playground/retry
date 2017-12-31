package retry

import (
	"errors"
	"fmt"
	"testing"
)

func TestRetry(t *testing.T) {
	err := Run(5,
		func() (bool, error) {
			return false, errors.New("ERR")
		},
		func(attempt uint16, err error) {
			fmt.Printf("Attempt: %d Error: %s\n", attempt, err)
		},
	)
	if err == nil {
		t.Fatalf("Expected %v Got %s", nil, err)
	}
}

func TestRetryBail(t *testing.T) {
	var count uint
	err := Run(5,
		func() (bool, error) {
			count++
			return true, errors.New("ERR")
		},
		func(attempt uint16, err error) {
			fmt.Printf("Attempt: %d Error: %s\n", attempt, err)
		},
	)
	if err == nil {
		t.Fatalf("Expected %v Got %s", nil, err)
	}
	if count != 1 {
		t.Fatalf("Expected %v Got %d", 1, count)
	}
}
