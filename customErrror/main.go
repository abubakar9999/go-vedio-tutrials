package main

import (
	// "errors"
	"fmt"
)

type Customerror struct {
	Msg string
	val int
}

// Error implements error.
func (Customerror) Error() string {
	panic("unimplemented")
}

func (e Customerror) myerror() string {
	return fmt.Sprint(e.Msg)
}

func Mydevited(num1 int, num2 int) (int, error) {
	if num2 == 0 {

		return 0, Customerror{Msg: "error is value zerro", val: 50}
	} else if num1%num2 != 0 {
		return 0, Customerror{Msg: "this is odd number", val: 50}

	}
	return num1 / num2, nil
}

func main() {
	data, err := Mydevited(4, 2)

	if err != nil {

		fmt.Println(data)

	}

}
