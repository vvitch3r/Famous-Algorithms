package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Quick sort algorithm----------------------------------------------------------------

var quickSortMovement int   // Variable to count the number of movements (swaps) in QuickSort
var quickSortComparison int // Variable to count the number of comparisons in QuickSort

// quickSort is the main function that implements the QuickSort algorithm
func quickSort(arr []int, start, end int) {
	if end <= start {
		return // Base case: If the end index is less than or equal to the start index, return
	}
	pivot := partition(arr, start, end) // Partition the array and get the pivot index
	quickSort(arr, start, pivot-1)      // Recursively sort the elements before the pivot
	quickSort(arr, pivot+1, end)        // Recursively sort the elements after the pivot
}

// partition rearranges the elements in the array so that all elements less than the pivot are to the left
// of the pivot and all elements greater than the pivot are to the right of the pivot
func partition(arr []int, start, end int) int {
	pivot := arr[end] // Choose the last element as the pivot
	i := start - 1    // Index of the smaller element

	for j := start; j <= end-1; j++ {
		if arr[j] < pivot {
			i++ // Increment index of the smaller element
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
			quickSortMovement++ // Increment movement counter
		}
		quickSortComparison++ // Increment comparison counter
	}
	i++
	// Swap the pivot element with the element at index i
	arr[i], arr[end] = arr[end], arr[i]
	quickSortMovement++ // Increment movement counter
	return i            // Return the index of the pivot
}

// Merge sort algorithm--------------------------------------------------------------

var mergeSortMovement int   // Variable to count the number of movements (copies) in MergeSort
var mergeSortComparison int // Variable to count the number of comparisons in MergeSort

// mergeSort is the main function that implements the MergeSort algorithm
func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return // Base case: the array is already sorted or empty
	}

	mid := len(arr) / 2                // Find the middle point to divide the array into two halves
	left := make([]int, mid)           // Create a temporary array for the left half
	right := make([]int, len(arr)-mid) // Create a temporary array for the right half

	copy(left, arr[:mid])  // Copy the left half of arr to left
	copy(right, arr[mid:]) // Copy the right half of arr to right

	mergeSort(left)  // Recursively sort the left half
	mergeSort(right) // Recursively sort the right half

	merge(arr, left, right) // Merge the sorted halves back into arr
}

// merge combines two sorted arrays (left and right) into a single sorted array (arr)
func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0 // Initialize indices for left, right, and merged arrays

	// Merge the two arrays while there are elements in both left and right
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++                 // Move to the next element in left
			mergeSortMovement++ // Increment movement counter
		} else {
			arr[k] = right[j]
			j++                 // Move to the next element in right
			mergeSortMovement++ // Increment movement counter
		}
		k++                   // Move to the next position in merged array
		mergeSortComparison++ // Increment comparison counter
	}

	// Copy any remaining elements from left
	for i < len(left) {
		arr[k] = left[i]
		i++                 // Move to the next element in left
		k++                 // Move to the next position in merged array
		mergeSortMovement++ // Increment movement counter
	}

	// Copy any remaining elements from right
	for j < len(right) {
		arr[k] = right[j]
		j++                 // Move to the next element in right
		k++                 // Move to the next position in merged array
		mergeSortMovement++ // Increment movement counter
	}
}

// Floyd algorithm-----------------------------------------------------------

// Define a large number to represent infinity
const INF = math.MaxInt32

// FloydWarshall function to implement the algorithm
func FloydWarshall(graph [][]int) [][]int {
	// Get the number of vertices in the graph
	V := len(graph)

	// Initialize the distance matrix
	dist := make([][]int, V)
	for i := range dist {
		dist[i] = make([]int, V)
		copy(dist[i], graph[i])
	}

	// Implement the Floyd-Warshall algorithm
	for k := 0; k < V; k++ {
		for i := 0; i < V; i++ {
			for j := 0; j < V; j++ {
				// If vertex k is on the shortest path from i to j,
				// then update the value of dist[i][j]
				if dist[i][k] != INF && dist[k][j] != INF && dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

// Dijkstra algorithm----------------------------------------------------------

// Edge represents an edge in the graph
type Edge struct {
	to, weight int
}

// Graph represents a weighted directed graph using an adjacency list
type Graph struct {
	vertices int
	adjList  [][]Edge
}

// NewGraph initializes a new graph with a given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make([][]Edge, vertices),
	}
}

// AddEdge adds a directed edge to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.adjList[from] = append(g.adjList[from], Edge{to, weight})
}

// Item represents an item in the priority queue
type Item struct {
	vertex, distance int
}

// PriorityQueue implements a min-heap for Items
type PriorityQueue []*Item

// Len returns the length of the priority queue
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares two items in the priority queue
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

// Swap swaps two items in the priority queue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds an item to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

// Pop removes and returns the smallest item from the priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Dijkstra finds the shortest path from a source vertex to all other vertices
func (g *Graph) Dijkstra(src int) []int {
	// Initialize distances with infinity
	distances := make([]int, g.vertices)
	for i := range distances {
		distances[i] = math.MaxInt32
	}
	distances[src] = 0

	// Priority queue to select the vertex with the smallest distance
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{vertex: src, distance: 0})

	for pq.Len() > 0 {
		// Get the vertex with the smallest distance
		item := heap.Pop(pq).(*Item)
		vertex := item.vertex
		distance := item.distance

		// Iterate over adjacent vertices
		for _, edge := range g.adjList[vertex] {
			newDist := distance + edge.weight
			if newDist < distances[edge.to] {
				distances[edge.to] = newDist
				heap.Push(pq, &Item{vertex: edge.to, distance: newDist})
			}
		}
	}

	return distances
}

func main() {
	Flag := false
	for !Flag {
		fmt.Println("----------------------------------")
		fmt.Println("1) Quick sort")
		fmt.Println("2) Merge sort")
		fmt.Println("3) Floyd algorithm")
		fmt.Println("4) Dijkstra algorithm")
		fmt.Println("5) Exit")
		fmt.Println("----------------------------------")
		fmt.Print("Enter your chosen option number =====> ")
		inputOption := 0
		fmt.Scanln(&inputOption)
		switch inputOption {
		case 1:
			fmt.Print("Creating a random list")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
			fmt.Println()
			//creating random list:
			var qlist = [1000]int{}
			for i := range qlist {
				qlist[i] = rand.Intn(1000)
			}
			for j := range qlist {
				p := qlist[j]
				fmt.Print(p, " - ")
			}
			fmt.Println()
			//arr := []int{10, 7, 8, 9, 1, 5}
			n := len(qlist)
			quickSort(qlist[:], 0, n-1)
			fmt.Println("Sorted array:", qlist)
			fmt.Println()
			fmt.Println("Number of comparisons: ", quickSortComparison)
			fmt.Println("Number of movements: ", quickSortMovement)
			fmt.Println()
			fmt.Scanln()
		case 2:
			fmt.Print("Creating a random list")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
			fmt.Println()
			//creating random list:
			qlist := [1000]int{}
			for i := range qlist {
				qlist[i] = rand.Intn(1000)
			}
			for j := range qlist {
				p := qlist[j]
				fmt.Print(p, " - ")
			}
			fmt.Println()
			// sorting
			mergeSort(qlist[:])
			//print the sorted list
			fmt.Print("Sorted list: ")
			for j := range qlist {
				p := qlist[j]
				fmt.Print(p, " - ")
			}
			fmt.Println()
			fmt.Println()
			fmt.Println("Number of comparisons: ", mergeSortComparison)
			fmt.Println("Number of movements: ", mergeSortMovement)
			fmt.Println()
			fmt.Scanln()
			//tests:
			//arr := []int{9, 7, 2, 11, 1, 39, 54}
			// fmt.Println("Original array:", qlist)
			// mergSort(arr)
			// fmt.Println("Sorted array:", qlist)
		case 3:
			// Define the graph as an adjacency matrix
			graph := [][]int{
				{0, 3, 6, INF},
				{INF, 0, 2, 1},
				{INF, INF, 0, 1},
				{INF, INF, INF, 0},
			}

			// Run the Floyd-Warshall algorithm
			dist := FloydWarshall(graph)

			// Print the distance matrix
			fmt.Println("The following matrix shows the shortest distances between every pair of vertices:")
			for i := 0; i < len(dist); i++ {
				for j := 0; j < len(dist[i]); j++ {
					if dist[i][j] == INF {
						fmt.Print("INF ")
					} else {
						fmt.Printf("%3d ", dist[i][j])
					}
				}
				fmt.Println()
			}
			fmt.Scanln()
		case 4:
			// Create a new graph with 5 vertices
			graph := NewGraph(5)

			// Add edges to the graph
			graph.AddEdge(0, 1, 10)
			graph.AddEdge(0, 4, 5)
			graph.AddEdge(1, 2, 1)
			graph.AddEdge(1, 4, 2)
			graph.AddEdge(2, 3, 4)
			graph.AddEdge(3, 0, 7)
			graph.AddEdge(3, 2, 6)
			graph.AddEdge(4, 1, 3)
			graph.AddEdge(4, 2, 9)
			graph.AddEdge(4, 3, 2)

			// Find the shortest path from vertex 0
			distances := graph.Dijkstra(0)

			// Print the shortest distances to all vertices
			fmt.Println("Shortest distances from vertex 0:")
			for i, d := range distances {
				fmt.Printf("Vertex %d: %d\n", i, d)
			}
			fmt.Scanln()
		case 5:
			fmt.Print("logging out")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
			fmt.Println()
			Flag = true
		default:
			fmt.Print("Error: ")
			fmt.Println("Entered value is not valid.")

		}
	}
}
