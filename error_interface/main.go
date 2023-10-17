package main

import (
	"errors"
	"fmt"
)

func Devided(a, b int) (int, error) {
	if a == 0 {
		return 0, errors.New("this is error")

	}
	return a / b, errors.New("no error")
}

func main() {


	res, err := Devided(10, 0)

	if err != nil {

		fmt.Println("result is", res)
	}

}

