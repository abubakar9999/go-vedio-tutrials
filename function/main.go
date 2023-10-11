package main

import "fmt"

func main() {
	coustomfunction(10, "aub baldr")
	a, b := myFunction(5, "Hello")
	fmt.Println(a, b)


}
func coustomfunction(x int, y string) {
	var n = x
	var z = y

	fmt.Println(n)
	fmt.Println(z)

}

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
  }
