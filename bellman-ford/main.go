package main

import "fmt"

func main() {
	// var allPositive = NewWeightedGraph[int]()
	// allPositive.Add(0, map[int]int{1: 5, 2: 2})
	// allPositive.Add(1, map[int]int{3: 4, 4: 2})
	// allPositive.Add(2, map[int]int{1: 8, 4: 7})
	// allPositive.Add(3, map[int]int{4: 6, 5: 3})
	// allPositive.Add(4, map[int]int{5: 1})

	// fmt.Println(allPositive.BellmanFord(0, 5))
	// fmt.Println(allPositive.BellmanFord(4, 5))
	// fmt.Println(allPositive.BellmanFord(2, 0))

	// var hasNegative = NewWeightedGraph[int]()
	// hasNegative.Add(0, map[int]int{1: 2, 2: 2})
	// hasNegative.Add(1, map[int]int{3: 2, 4: 2})
	// hasNegative.Add(2, map[int]int{1: 2})
	// hasNegative.Add(4, map[int]int{2: -4, 3: 2})

	// fmt.Println(hasNegative.BellmanFord(0, 3))

	var hasNegativeCycle = NewWeightedGraph[string]()
	hasNegativeCycle.Add("A", map[string]int{"C": 4, "E": 5})
	hasNegativeCycle.Add("D", map[string]int{"A": 4, "C": 7, "E": 3})
	hasNegativeCycle.Add("C", map[string]int{"A": -9})
	hasNegativeCycle.Add("E", map[string]int{"B": 2, "C": 3})
	hasNegativeCycle.Add("B", map[string]int{"C": -4})

	fmt.Println(hasNegativeCycle.BellmanFord("A", "B"))
}
