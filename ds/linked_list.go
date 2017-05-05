package ds

import "fmt"

type LinkedList struct {
	head    *Node
	current *Node
	size    int
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) New() {
	l.head = new(Node)
	l.current = new(Node)
}

func (l *LinkedList) Insert(value interface{}) {
	if l.head == nil {
		l.New()
		l.head = &Node{value, nil}
		l.current = l.head
	} else {
		newNode := &Node{value, l.current.next}
		l.current.next = newNode
		l.current = newNode
	}
	l.size++
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%v, ", current.data)
		current = current.next
	}
	fmt.Println()
}
