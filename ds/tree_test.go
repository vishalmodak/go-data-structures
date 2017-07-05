package ds

import (
	"testing"
)

func TestInOrderRecursive(t *testing.T) {
	tree := new(BST)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(11)
	tree.Insert(17)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)

	expected := "1, 2, 5, 10, 11, 15, 17,"
	InOrderRecursive(tree.GetRoot())
	if result != expected {
		t.Errorf("In-Order Traversal (Recursive) incorrect, got %s, expected %s", *result, expected)
	}
}
