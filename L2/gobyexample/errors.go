package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

var ErrorOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrorOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	} else {
		return nil
	}
}

func main() {

	for _, v := range []int{7, 42} {
		if i, err := f(v); err != nil {
			fmt.Println("f failed:", err)
		} else {
			fmt.Println("f worked:", i)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrorOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}

//f worked: 10
//f failed: can't work with 42
//Tea is ready!
//Tea is ready!
//We should buy new tea!
//Tea is ready!
//Now it is dark.
