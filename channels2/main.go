package main

import (
	"fmt"
	"time"
)

func getNumber(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(250*time.Millisecond)
		ch <- i

	}

}
func main() {
	ch := make(chan int)

	go getNumber(ch)

	for num := range ch {

		fmt.Println(num)

	}
	close(ch)

}
