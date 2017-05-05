package ds

import (
	"testing"
)

func TestStackLL_Pop(t *testing.T) {
	stack := new(StackLL)
	stack.Push("a")
	stack.Push("b")

	result := stack.Pop()
	if result != "b" {
		t.Errorf("Popped data incorrect, got %s, expected %s", result, "b")
	}
	if stack.size != 1 {
		t.Errorf("Stack size incorrect, got %v, expected %v", stack.size, 1)
	}
}

func TestStackLL_Pop2(t *testing.T) {
	stack := new(StackLL)
	stack.Pop()
}

func TestStackLL_Push(t *testing.T) {
	stack := new(StackLL)
	stack.Push("a")
	stack.Push("b")

	if stack.size != 2 {
		t.Errorf("Stack size incorrect, got %v, expected %v", stack.size, 2)
	}
}

func TestStackLL_IsEmpty(t *testing.T) {
	stack := new(StackLL)
	stack.Push("a")
	stack.Push("b")
	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Errorf("Stack should be empty, got %v, expected %v", stack.IsEmpty(), true)
	}
}
