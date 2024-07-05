package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func binarySearch(list []int, elem int) (int, bool) {
	low := 0
	high := len(list) - 1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		curr := list[mid]
		if curr == elem {
			return mid, true
		} else if curr < elem {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0, false
}

func main() {
	randList := make([]int, 100)
	for i := range randList {
		randList[i] = rand.Intn(1000)
	}
	sort.Ints(randList)

	pos, found := binarySearch(randList, 103)
	fmt.Println(pos, found)

	pos, found = binarySearch(randList, 678)
	fmt.Println(pos, found)

	pos, found = binarySearch(randList, 293)
	fmt.Println(pos, found)

	pos, found = binarySearch(randList, 113)
	fmt.Println(pos, found)

	pos, found = binarySearch(randList, 42)
	fmt.Println(pos, found)

	pos, found = binarySearch(randList, 602)
	fmt.Println(pos, found)

	pos, found = binarySearch(randList, 513)
	fmt.Println(pos, found)

	pos, found = binarySearch(randList, 37)
	fmt.Println(pos, found)

	fmt.Println(randList)
}
