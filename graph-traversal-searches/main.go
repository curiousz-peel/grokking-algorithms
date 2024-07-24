package main

import "fmt"

type Node[T comparable] struct {
	val   T
	depth int
}

type Queue[T comparable] struct {
	nodes []Node[T]
}

func (q *Queue[T]) Enqueue(val Node[T]) {
	q.nodes = append(q.nodes, val)
}

func (q *Queue[T]) Dequeue() (Node[T], bool) {
	var first Node[T]

	if len(q.nodes) == 0 {
		return first, false
	}

	first = q.nodes[0]
	q.nodes = q.nodes[1:]
	return first, true
}

type Stack[T comparable] struct {
	nodes []Node[T]
}

func (s *Stack[T]) Push(val Node[T]) {
	s.nodes = append(s.nodes, val)
}

func (s *Stack[T]) Pop() (Node[T], bool) {
	var last Node[T]

	if len(s.nodes) == 0 {
		return last, false
	}

	last = s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return last, false
}

type Graph[T comparable] struct {
	adjacencyList map[T][]T
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{adjacencyList: map[T][]T{}}
}

func (g *Graph[T]) Add(node T, neighbours []T) {
	g.adjacencyList[node] = neighbours
}

func (g *Graph[T]) BFS(from, to T) (int, bool) {
	visited := make(map[T]struct{}, len(g.adjacencyList))
	queue := &Queue[T]{nodes: make([]Node[T], 0, len(g.adjacencyList))}
	var depth = 0

	if from == to {
		return depth, true
	}

	visited[from] = struct{}{}
	for _, neighbour := range g.adjacencyList[from] {
		queue.Enqueue(Node[T]{val: neighbour, depth: depth + 1})
	}

	for len(queue.nodes) != 0 {
		curr, _ := queue.Dequeue()

		if curr.val == to {
			return curr.depth, true
		}

		visited[curr.val] = struct{}{}
		for _, neighbour := range g.adjacencyList[curr.val] {
			if _, ok := visited[neighbour]; !ok {
				queue.Enqueue(Node[T]{val: neighbour, depth: curr.depth + 1})
			}
		}
	}

	return 0, false
}

func (g *Graph[T]) DFS(from, to T) (int, bool) {
	visited := make(map[T]struct{}, len(g.adjacencyList))
	stack := &Stack[T]{nodes: make([]Node[T], 0, len(g.adjacencyList))}
	var depth int

	if from == to {
		return depth, true
	}

	for _, neighbour := range g.adjacencyList[from] {
		stack.Push(Node[T]{val: neighbour, depth: depth + 1})
	}

	for len(stack.nodes) != 0 {
		curr, _ := stack.Pop()

		if curr.val == to {
			return curr.depth, true
		}

		visited[curr.val] = struct{}{}
		for _, neighbour := range g.adjacencyList[curr.val] {
			if _, ok := visited[neighbour]; !ok {
				stack.Push(Node[T]{val: neighbour, depth: curr.depth + 1})
			}
		}
	}

	return 0, false
}

func main() {
	gInt := NewGraph[int]()
	gInt.Add(1, []int{2, 3})
	gInt.Add(2, []int{4, 5})
	gInt.Add(3, []int{4, 6})
	gInt.Add(6, []int{5})

	fmt.Println(gInt.BFS(1, 5))
	fmt.Println(gInt.DFS(1, 5))

	fmt.Println(gInt.BFS(2, 5))
	fmt.Println(gInt.DFS(2, 5))

	fmt.Println(gInt.BFS(2, 2))
	fmt.Println(gInt.DFS(2, 2))

	fmt.Println(gInt.BFS(2, 6))
	fmt.Println(gInt.DFS(2, 6))

	gString := NewGraph[string]()
	gString.Add("cab", []string{"cat", "car"})
	gString.Add("car", []string{"cat", "bar"})
	gString.Add("bar", []string{"bat"})
	gString.Add("cat", []string{"mat", "bat"})
	gString.Add("mat", []string{"bat"})

	fmt.Println(gString.BFS("cab", "bat"))
	fmt.Println(gString.BFS("car", "cab"))
}
