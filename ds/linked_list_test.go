package ds

import (
	"testing"
)

func TestListContainsPositive(t *testing.T) {
	list := new(LinkedList)
	list.Insert("A")
	list.Insert("B")
	list.Insert("C")

	result := list.Contains("C")
	if !result {
		t.Errorf("Contains() result incorrect, got %s, expected %s", result, true)
	}
}

func TestListContainsNeative(t *testing.T) {
	list := new(LinkedList)
	list.Insert("A")
	list.Insert("B")

	result := list.Contains("C")
	if result {
		t.Errorf("Contains() result incorrect, got %s, expected %s", result, false)
	}
}
