package ds

import "fmt"

type StackLL struct {
	top  *Node
	size int
}

func (s *StackLL) Print() {
	fmt.Printf("Stack size: %d\n", s.size)
	fmt.Printf("Stack Elements: ")
	var ptr *Node = s.top
	for ptr != nil {
		fmt.Printf("%v, ", ptr.data)
		ptr = ptr.next
	}
	fmt.Println()
}

func (s *StackLL) Top() interface{} {
	return s.top.data
}

func (s *StackLL) IsEmpty() bool {
	return s.top == nil
}

func (s *StackLL) Clear() {
	fmt.Println("Stack is cleared.")
	s.top = nil
	s.size = 0
}

func (s *StackLL) Push(data interface{}) {
	s.top = &Node{data, s.top}
	s.size++
}

func (s *StackLL) Pop() interface{} {
	if s.IsEmpty() {
		return ""
	}
	var data = s.top.data
	s.top = s.top.next
	s.size--
	return data
}

func (s *StackLL) Contains(value interface{}) bool {
	current := s.top
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}
