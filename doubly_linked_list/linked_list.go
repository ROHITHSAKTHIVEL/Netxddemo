package main

import "fmt"

type node struct {
	prev *node
	data int
	next *node
}

func createNode(a int) *node {
	curr := &node{prev: nil, data: a, next: nil}
	return curr
}

func forwardtraversal(head *node) {
	ptr := head
	for ptr != nil {
		fmt.Print(ptr.data)
		ptr = ptr.next
	}
	fmt.Println()
}
func backwardtraversal(curr *node) {
	ptr := curr
	for ptr != nil {
		fmt.Print(ptr.data)
		ptr = ptr.prev
	}
	fmt.Println()
}

func insertbeg(a int, ptr *node) *node {
	ptr1 := ptr
	temp := createNode(a)
	ptr1.prev = temp
	temp.next = ptr1
	return temp
}

func insertend(a int, ptr *node) *node {
	ptr1 := ptr
	temp := createNode(a)
	ptr1.next = temp
	temp.prev = ptr1
	return temp
}

func count_no_of_nodes(ptr *node) {
	ptr1 := ptr
	count := 1
	for ptr1.next != nil {
		count++
		ptr1 = ptr1.next
	}
	fmt.Println("count=", count)
}

func insertatpos(a, pos int, ptr *node) *node {
	ptr1 := ptr
	head := ptr
	for i := 1; i < pos; i++ {

		ptr1 = ptr1.next
	}
	temp := createNode(a)

	temp.prev = ptr1

	temp.next = ptr1.next
	if ptr1.next != nil {
		ptr1.next.prev = temp
	}
	ptr1.next = temp

	return head
}

func deletepos(pos int, ptr *node) *node {
	ptr1 := ptr
	temp := ptr1
	for i := 1; i < pos; i++ {
		ptr1 = ptr1.next
	}

	ptr1.prev.next = ptr1.next
	ptr1.next.prev = ptr1.prev

	return temp
}

func main() {

	head := createNode(10)

	curr := createNode(20)
	head.next = curr
	curr.prev = head

	curr = createNode(30)

	head.next.next = curr
	curr.prev = head.next
	temp := insertbeg(100, head)
	head = temp
	temp2 := insertend(200, curr)
	curr = temp2
	temp3 := insertatpos(500, 4, head)
	head = temp3
	temp4 := deletepos(4, head)
	head = temp4
	forwardtraversal(head)
	backwardtraversal(curr)
	count_no_of_nodes(head)

}

