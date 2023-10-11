package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// var stt string

	// fmt.Scanln(&stt)

	// fmt.Println(stt);
	// reder := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter a name")

	// myname, _ := reder.ReadString('\n')
	// fmt.Println(myname)

	print("Enter your number : ")

	reader:=bufio.NewReader(os.Stdin)
  temt, _ := reader.ReadString('\n')
 temt2,_ := strconv.ParseFloat(strings.TrimSpace(temt),64)


 fmt.Println(temt2 +2)
 


	

}
