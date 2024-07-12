package main

import "fmt"

func main() {
	// var sg = NewWeightedGraph[string]()
	// sg.Add(Node[string]{"book"}, map[Node[string]]int{{"lp"}: 5, {"poster"}: 0})
	// sg.Add(Node[string]{"lp"}, map[Node[string]]int{{"bass guitar"}: 15, {"drum set"}: 20})
	// sg.Add(Node[string]{"poster"}, map[Node[string]]int{{"bass guitar"}: 30, {"drum set"}: 35})
	// sg.Add(Node[string]{"bass guitar"}, map[Node[string]]int{{"piano"}: 20})
	// sg.Add(Node[string]{"drum set"}, map[Node[string]]int{{"piano"}: 10})

	// fmt.Println(sg.Search(Node[string]{"book"}, Node[string]{"piano"}))
	// fmt.Println(sg.Search(Node[string]{"poster"}, Node[string]{"piano"}))
	// fmt.Println(sg.Search(Node[string]{"piano"}, Node[string]{"book"}))

	var ig = NewWeightedGraph[int]()
	ig.Add(Node[int]{0}, map[Node[int]]int{{1}: 5, {2}: 2})
	ig.Add(Node[int]{1}, map[Node[int]]int{{3}: 4, {4}: 2})
	ig.Add(Node[int]{2}, map[Node[int]]int{{1}: 8, {4}: 7})
	ig.Add(Node[int]{3}, map[Node[int]]int{{4}: 6, {5}: 3})
	ig.Add(Node[int]{4}, map[Node[int]]int{{5}: 1})

	fmt.Println(ig.Search(Node[int]{0}, Node[int]{5}))
	fmt.Println(ig.Search(Node[int]{4}, Node[int]{5}))
	fmt.Println(ig.Search(Node[int]{2}, Node[int]{0}))
}
