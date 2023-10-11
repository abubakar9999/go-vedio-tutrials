package main

import "fmt"

func main() {
	mymap := make(map[string]int)

	mymap["m"] = 10
	mymap["n"] = 120
	mymap["h"] = 102
	mymap["u"] = 1074
	mymap["d"] = 1058
	fmt.Println(mymap)

	delete(mymap, "u")
	fmt.Println(mymap)

	for k, v := range mymap {
		fmt.Printf("scross of %v is %v \n", k, v)

	}
}
