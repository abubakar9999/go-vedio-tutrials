package main

import (
	"fmt"
	"sort"
)

func main() {

	myslice := []int32{
		10, 2, 5, 6, 8, 5, 5}

	myslice = append(myslice, 100)
	fmt.Println(myslice)

	myslice = append(myslice[0:3])
	fmt.Println(myslice)

	myints := []int{1, 2, 3, 5, 8, 4, 5, 5, 58, 4, 8}

	issorted := sort.IntsAreSorted(myints)  //check this is sortede
	shortdata := sort.SearchInts(myints, 2) // indexAt

	sort.Ints(myints)
	fmt.Println(myints)

	fmt.Println(shortdata)

	fmt.Println(issorted)

}
