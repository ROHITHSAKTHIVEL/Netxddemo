package main

import (
	"fmt"
)

type Node struct {
	data     int
	previous *Node
	next     *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) Append(data int) {
	newNode := &Node{data: data, previous: nil, next: nil}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.previous = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

func (dll *DoublyLinkedList) PrintForward() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (dll *DoublyLinkedList) PrintBackward() {
	current := dll.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.previous
	}
	fmt.Println("nil")
}

func main() {
	dll := DoublyLinkedList{}

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	fmt.Println("Forward:")
	dll.PrintForward()

	fmt.Println("Backward:")
	dll.PrintBackward()
}
