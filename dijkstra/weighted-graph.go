package main

import (
	"math"
)

type Node[T comparable] struct {
	val T
}

type WeightedGraph[T comparable] struct {
	adjacencyList map[Node[T]]map[Node[T]]int
	costMap       map[Node[T]]int
	parentMap     map[Node[T]]Node[T]
	processed     map[Node[T]]struct{}
}

func NewWeightedGraph[T comparable]() *WeightedGraph[T] {
	return &WeightedGraph[T]{
		adjacencyList: map[Node[T]]map[Node[T]]int{},
		costMap:       map[Node[T]]int{},
		parentMap:     map[Node[T]]Node[T]{},
		processed:     map[Node[T]]struct{}{},
	}
}

func (w *WeightedGraph[T]) Add(node Node[T], neighbours map[Node[T]]int) {
	w.adjacencyList[node] = neighbours
}

func (w *WeightedGraph[T]) initSource(source Node[T]) *WeightedGraph[T] {
	al := w.adjacencyList
	w = NewWeightedGraph[T]()
	w.adjacencyList = al

	for node, cost := range w.adjacencyList[source] {
		w.costMap[node] = cost
		w.parentMap[node] = source
	}
	w.processed[source] = struct{}{}

	return w
}

func (w *WeightedGraph[T]) findLowestCostNode() Node[T] {
	lc := math.MaxInt
	var lcn Node[T]

	for node, cost := range w.costMap {
		if _, ok := w.processed[node]; !ok {
			if cost < lc {
				lc = cost
				lcn = node
			}
		}
	}
	return lcn
}

func (w *WeightedGraph[T]) buildPath(source, destination Node[T]) ([]Node[T], int, bool) {
	var none []Node[T]
	var s Stack[Node[T]]

	if _, ok := w.parentMap[destination]; !ok {
		return none, 0, false
	}

	s.Push(destination)
	for curr := w.parentMap[destination]; curr != source; curr = w.parentMap[curr] {
		s.Push(curr)
	}
	s.Push(source)

	var path []Node[T]
	for node, ok := s.Pop(); ok; node, ok = s.Pop() {
		path = append(path, node)
	}
	return path, w.costMap[destination], true
}

func (w *WeightedGraph[T]) Search(source, destination Node[T]) ([]Node[T], int, bool) {
	w = w.initSource(source)

	var none Node[T]
	lcn := w.findLowestCostNode()

	for lcn != none {
		cost := w.costMap[lcn]
		neighbours := w.adjacencyList[lcn]

		for neighbour, neighbourCost := range neighbours {
			newCost := cost + neighbourCost
			oldCost, ok := w.costMap[neighbour]
			if !ok || newCost < oldCost {
				w.costMap[neighbour] = newCost
				w.parentMap[neighbour] = lcn
			}
		}
		w.processed[lcn] = struct{}{}
		lcn = w.findLowestCostNode()
	}
	return w.buildPath(source, destination)
}
