package main

import (
	"fmt"
	"math/rand"
)

func quicksort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	pivotIdx := rand.Intn(len(slice))
	var less, greater []int

	for i, el := range slice {
		if i == pivotIdx {
			continue
		}
		if el <= slice[pivotIdx] {
			less = append(less, el)
		} else {
			greater = append(greater, el)
		}
	}

	return append(append(quicksort(less), slice[pivotIdx]), quicksort(greater)...)
}

func main() {
	slice := make([]int, 10000)

	for i := range slice {
		slice[i] = rand.Intn(10000)
	}

	fmt.Println(quicksort(slice))
}
