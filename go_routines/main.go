package main

import (
	"fmt"
	"time"
)

func numberPrint() {

	for i := 0; i <= 5; i++ {
		time.Sleep(250 * time.Microsecond)
		fmt.Println(i)

	}

}

func latterPrint() {

	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Microsecond)
		fmt.Printf("%c\n", i)

	}

}

func main() {
	go numberPrint()
	go latterPrint()

	time.Sleep(2000 * time.Microsecond)

	fmt.Println("finish code")
}
