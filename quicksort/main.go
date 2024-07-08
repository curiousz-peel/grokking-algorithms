package main

import (
	"cmp"
	"fmt"
	"math/rand"
)

func quicksort[T cmp.Ordered](slice []T) []T {
	if len(slice) < 2 {
		return slice
	}

	pivotIdx := rand.Intn(len(slice))
	var less, greater []T

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
	intSlice := make([]int, 10000)

	for i := range intSlice {
		intSlice[i] = rand.Intn(10000)
	}

	// fmt.Println(quicksort(intSlice))

	stringSlice := []string{"orange", "kiwi", "grape", "mango", "banana",
		"raspberry", "cherry", "fig", "elderberry", "lemon",
		"honeydew", "tangerine", "papaya", "strawberry", "iceberg",
		"jackfruit", "apple", "quince", "nectarine", "date"}

	fmt.Println(quicksort(stringSlice))
}
