package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type Stack struct {
	top  *Node
	size int
}

func main() {
	stack := new(Stack)
	if stack.IsEmpty() {
		fmt.Println("Stack is empty.")
	}
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")
	stack.Print()

	fmt.Printf("pop() data: %s\n", stack.Pop())
	fmt.Printf("pop() data: %s\n", stack.Pop())
	stack.Print()

	stack.Clear()
	stack.Print()
	stack.Pop()
}

func (s *Stack) Print() {
	fmt.Printf("Stack size: %d\n", s.size)
	var elements string
	var ptr *Node = s.top
	for ptr != nil {
		elements += ptr.data + ","
		ptr = ptr.next
	}
	fmt.Printf("Stack Elements: [%s]\n", elements)
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Clear() {
	fmt.Println("Stack is cleared.")
	s.top = nil
	s.size = 0
}

func (s *Stack) Push(data string) {
	s.top = &Node{data, s.top}
	s.size++
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	var data = s.top.data
	s.top = s.top.next
	s.size--
	return data
}
