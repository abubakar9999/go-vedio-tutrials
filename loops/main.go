package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5}

	fmt.Print(array)

	for i := 0; i < len(array); i++ {

		fmt.Printf("%v\n", array[i])

	}
	println("**********")

	for _, data := range array {
		fmt.Println(data)
	}
}
