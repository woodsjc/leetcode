package main

import (
	"fmt"
)

type Node int

type AdjacencyList map[Node][]Node

type TestRoute struct {
	List   AdjacencyList
	Start  Node
	End    Node
	Exists bool
}

var visited map[Node]bool

func main() {
	// Takes adjacency list and returns route if exists
	graph := AdjacencyList{
		0: {1},
		1: {2},
		2: {0, 3},
		3: {2},
		4: {6},
		5: {4},
		6: {5},
	}

	tests := []TestRoute{
		{
			List:   graph,
			Start:  0,
			End:    6,
			Exists: false,
		},
		{
			List:   graph,
			Start:  1,
			End:    3,
			Exists: true,
		},
		{
			List:   graph,
			Start:  4,
			End:    6,
			Exists: true,
		},
	}

	for _, t := range tests {
		visited = make(map[Node]bool)
		result := routeExists(t.Start, t.End, t.List)

		fmt.Printf("Test: %v\n", t)
		fmt.Printf("Result ")
		if result == t.Exists {
			fmt.Println("success")
		} else {
			fmt.Println("fail")
		}
	}
}

func routeExists(start Node, end Node, routes AdjacencyList) bool {
	// try depth first search
	current := start
	visited[current] = true

	if current == end {
		return true
	}

	for _, next := range routes[current] {
		if alreadyVisited, exists := visited[next]; exists && alreadyVisited {
			continue
		}

		if routeExists(next, end, routes) {
			return true
		}
	}

	return false
}
