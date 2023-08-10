package main

import (
	"fmt"
	"time"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key, Left: nil, Right: nil}
	}

	if key < root.Key {
		root.Left = insert(root.Left, key)
	} else if key > root.Key {
		root.Right = insert(root.Right, key)
	}

	return root
}

func search(root *Node, key int) bool {
	if root == nil {
		return false
	}

	if key == root.Key {
		return true
	} else if key < root.Key {
		return search(root.Left, key)
	} else {
		return search(root.Right, key)
	}
}

func main() {
	var root *Node

	// Create a 16-level Binary Search Tree with random elements
	numElements := 1 << 16 // 2^16 nodes
	for i := 0; i < numElements; i++ {
		root = insert(root, i)
	}

	// Element to search for
	target := numElements / 2 // Searching for the middle element

	// Start the timer
	start := time.Now()

	// Search for the element in the Binary Search Tree
	found := search(root, target)

	// Stop the timer
	elapsed := time.Since(start)

	if found {
		fmt.Printf("Element %d found in the tree.\n", target)
	} else {
		fmt.Printf("Element %d not found in the tree.\n", target)
	}

	fmt.Printf("Time taken: %s\n", elapsed)
}
