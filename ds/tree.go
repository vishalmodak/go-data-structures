package ds

import (
	"fmt"
)

type TreeNode struct {
	key   int
	left  *TreeNode
	right *TreeNode
}

type BST struct {
	root *TreeNode
}

func (t *BST) GetRoot() *TreeNode {
	return t.root
}

func (t *TreeNode) isLeaf() bool {
	return t.left == nil && t.right == nil
}

func (t *BST) Insert(newKey int) {
	newNode := &TreeNode{newKey, nil, nil}

	if t.root == nil {
		t.root = newNode
		return
	}

	var root *TreeNode = t.root
	var parent *TreeNode
	for root != nil {
		parent = root
		if root.key > newKey {
			root = root.left
		} else if root.key < newKey {
			root = root.right
		} else {
			return
		}
	}
	if parent.key > newKey {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

}

func InOrderRecursive(node *TreeNode) {
	if node.left != nil {
		InOrderRecursive(node.left)
	}
	fmt.Printf("%v, ", node.key)
	if node.right != nil {
		InOrderRecursive(node.right)
	}
}

func PreOrderRecursive(node *TreeNode) {
	fmt.Printf("%v, ", node.key)
	if node.left != nil {
		PreOrderRecursive(node.left)
	}
	if node.right != nil {
		PreOrderRecursive(node.right)
	}
}

func PostOrderRecursive(node *TreeNode) {
	if node.left != nil {
		PostOrderRecursive(node.left)
	}
	if node.right != nil {
		PostOrderRecursive(node.right)
	}
	fmt.Printf("%v, ", node.key)
}
