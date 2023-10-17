package main

import "fmt"

func server1(ch1 chan string) {
	ch1 <- "Server 1"

}

func sever2(ch2 chan string) {

	ch2 <- "Server 2"

}

func main() {
	output1 := make(chan string)

	output2 := make(chan string)

	go server1(output1)

	go sever2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <- output2:
		fmt.Println(s2)
	}

}