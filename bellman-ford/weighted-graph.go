package main

import (
	"errors"
)

type WeightedGraph[T comparable] struct {
	adjacencyList map[T]map[T]int
	costMap       map[T]int
	parentMap     map[T]T
	processed     map[T]struct{}
}

func NewWeightedGraph[T comparable]() *WeightedGraph[T] {
	return &WeightedGraph[T]{
		adjacencyList: map[T]map[T]int{},
		costMap:       map[T]int{},
		parentMap:     map[T]T{},
		processed:     map[T]struct{}{},
	}
}

func (g *WeightedGraph[T]) Add(node T, neighbours map[T]int) {
	g.adjacencyList[node] = neighbours
}

func (g *WeightedGraph[T]) initSource(node T) *WeightedGraph[T] {
	al := g.adjacencyList
	g = NewWeightedGraph[T]()
	g.adjacencyList = al

	var zero T
	g.costMap[node] = 0
	g.parentMap[node] = zero

	return g
}

func (g *WeightedGraph[T]) buildPath(source, dest T) ([]T, int, error) {
	var path []T
	var s Stack[T]

	if _, ok := g.parentMap[dest]; !ok {
		return path, 0, errors.New("can't reach destination node from source node")
	}

	s.Push(dest)
	for node := g.parentMap[dest]; node != source; node = g.parentMap[node] {
		s.Push(node)
	}
	s.Push(source)

	for node, ok := s.Pop(); ok; node, ok = s.Pop() {
		path = append(path, node)
	}
	return path, g.costMap[dest], nil
}

func (g *WeightedGraph[T]) relaxEdges() {
	for node, neighbours := range g.adjacencyList {
		for neighbour, cost := range neighbours {
			costTo, ok := g.costMap[node]
			if !ok {
				continue
			}
			newCost := costTo + cost
			oldCost, ok := g.costMap[neighbour]
			if !ok || (newCost < oldCost) {
				g.costMap[neighbour] = newCost
				g.parentMap[neighbour] = node
			}
		}
	}
}

func (g *WeightedGraph[T]) BellmanFord(source, dest T) ([]T, int, error) {
	g = g.initSource(source)
	var zero []T

	for range len(g.adjacencyList) - 1 {
		g.relaxEdges()
	}

	//check all edges one more time to detect negative cycles
	//computing the shortest path in a graph with negative cycles makes no sense ->
	//because you can always find a shorter distance by looping one more time

	for node, neighbours := range g.adjacencyList {
		for neighbour, cost := range neighbours {
			costTo := g.costMap[node]
			newCost := costTo + cost
			oldCost := g.costMap[neighbour]
			if newCost < oldCost {
				return zero, 0, errors.New("graph has a negative cycle")
			}
		}
	}
	return g.buildPath(source, dest)
}
