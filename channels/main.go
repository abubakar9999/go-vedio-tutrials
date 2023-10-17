package main

import "fmt"

func Muyltiplay(a int, b int,ch chan int)  {
	fmt.Println("data is loaddding")
	ch <- (a * b)

}

func main() {
	ch:= make(chan int)
	go Muyltiplay(10,20,ch)
	res := <-ch
	fmt.Println("******** **********")
	// fmt.Println(Muyltiplay(10,20))
	
	fmt.Println(res)
	close(ch)
}