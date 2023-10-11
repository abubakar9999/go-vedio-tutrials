package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "time"
)

func main() {

	var name string
	var rating string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter you name")
	name, _ = reader.ReadString('\n')

	// reader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter our rateing")

	rating, _ = reader.ReadString('\n')
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	fmt.Printf(name,mynumrating)

}
