package main

import (
	"errors"
	"fmt"
)

type argError struct {
	age int
	msg string
}

func (a *argError) Error() string {
	return fmt.Sprintf("%d - %s", a.age, a.msg)
}
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var as *argError
	if errors.As(err, &as) {
		fmt.Println(as.age)
		fmt.Println(as.msg)
	} else {
		fmt.Println("err doesn't match argError")
	}
}

//42
//can't work with it
