package main

import (
	"fmt"
)

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Queue size:", queue.Size()) // Output: Queue size: 3

	item := queue.Dequeue()
	fmt.Println("Dequeued item:", item) // Output: Dequeued item: 1

	fmt.Println("Queue size:", queue.Size())        // Output: Queue size: 2
	fmt.Println("Is queue empty:", queue.IsEmpty()) // Output: Is queue empty: false

	queue.Dequeue()
	queue.Dequeue()

	fmt.Println("Is queue empty:", queue.IsEmpty()) // Output: Is queue empty: true
}
