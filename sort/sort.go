package main

import (
	"fmt"
	"sort"
)

func main() {
	// list is a slice of ints. It is unsorted as you can see
	list := []int{1, 23, 65, 11, 0, 3, 233, 88, 99}
	fmt.Println("The list is: ", list)

	// let's use Ints function that comes in sort
	// Ints([]int) sorts its parameter in ibcreasing order. Go read its doc.
	sort.Ints(list)
	fmt.Println("The sorted list is: ", list)
}
