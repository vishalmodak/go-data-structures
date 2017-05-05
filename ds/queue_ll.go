package ds

import (
	"fmt"
)

type QueueLL struct {
	front *Node
	back  *Node
	size  int
}

func (q *QueueLL) Clear() {
	q.front = nil
	q.back = nil
	q.size = 0
}

func (q *QueueLL) isEmpty() bool {
	return q.front == nil && q.back == nil
}

func (q *QueueLL) Print() {
	fmt.Printf("Queue size: %d\n", q.size)
	fmt.Printf("Queue Elements: ")
	var ptr *Node = q.front
	for ptr != nil {
		fmt.Printf("%v, ", ptr.data)
		ptr = ptr.next
	}
	fmt.Println()
}

func (q *QueueLL) GetFront() *Node {
	return q.front
}

func (q *QueueLL) GetBack() *Node {
	return q.back
}

func (q *QueueLL) Enqueue(data interface{}) {
	if q.isEmpty() {
		q.front = &Node{data, nil}
		q.back = q.front
	} else {
		q.back.next = &Node{data, nil}
		q.back = q.back.next
	}
	q.size++
}

func (q *QueueLL) Deque() interface{} {
	if q.isEmpty() {
		return ""
	}
	var data = q.front.data
	q.front = q.front.next
	q.size--
	return data
}
