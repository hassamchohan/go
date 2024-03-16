package main

import "fmt"

type Queue struct {
	List []int
}

// function to add element in queue
func (q *Queue) Enqueue(element int) {
	q.List = append(q.List, element)
}

// function to delete element in the queue
func (q *Queue) Dequeue() int {
	if q.isEmpty() {
		fmt.Println("Queue is empty.")
		return 0
	}
	element := q.List[0]
	q.List = q.List[1:]

	return element
}

// function check that queue is empty or not
func (q *Queue) isEmpty() bool {
	return len(q.List) == 0
}

// BFS() is a function with matrix and int value as parameter
func BFS(graph [][]int, node int) {
	// initializing the map that will keep
	// the track is the node is visited or not
	isvisited := make(map[int]bool)

	// creating a Queue variable
	// in which we will add an element at the same level
	// of that node
	var bfsQueue Queue

	// marking current node as visited
	isvisited[node] = true

	// adding a current node in the queue
	bfsQueue.Enqueue(node)

	// running a for loop until the queue becomes empty
	for !bfsQueue.isEmpty() {
		currNode := bfsQueue.List[0]
		fmt.Print(currNode, " ")
		// adding all the connected node in queue if not visted
		for nodes := 0; nodes < len(graph[currNode]); nodes++ {
			if graph[currNode][nodes] == 1 && !isvisited[nodes] {
				bfsQueue.Enqueue(nodes)
				isvisited[nodes] = true
			}
		}
		// removing the current node from queue
		// after visiting
		bfsQueue.Dequeue()
	}
}

func main() {
	// matrix representation of the undirected connected graph
	// where if the value is 1 the node i is connected
	// with node j
	graph := [][]int{{0, 1, 0, 1}, {1, 0, 1, 0}, {0, 1, 0, 1}, {1, 0, 1, 0}}

	fmt.Println("Golang program to do Breath first search of an undirected connected graph represented in the form of a matrix.")

	// calling BFS() function for the breadth-first search
	// traversal of a graph
	BFS(graph, 0)

	fmt.Println()
}
