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
	if node == nil {
		return
	}

	InOrderRecursive(node.left)
	fmt.Printf("%v, ", node.key)
	InOrderRecursive(node.right)
}

/*
1) Create an empty stack
2) Initialize current node as root
3) Push the current node to stack and set current = current->left until current is NULL
4) If current is NULL and stack is not empty then
     a) Pop the top item from stack.
     b) Print the popped item, set current = popped_item->right
     c) Go to step 3.
5) If current is NULL and stack is empty then we are done.
 */
func InOrder(node *TreeNode) {
	stack := new(StackLL)
	currentNode := node
	done := false
	for !done {
		if currentNode != nil {
			stack.Push(currentNode)
			currentNode = currentNode.left
		} else {
			if !stack.IsEmpty() {
				currentNode = stack.Pop().(*TreeNode)
				fmt.Printf("%v, ", currentNode.key)
				currentNode = currentNode.right
			} else {
				done = true
			}
		}
	}
}

func PreOrderRecursive(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%v, ", node.key)
	PreOrderRecursive(node.left)
	PreOrderRecursive(node.right)
}

/*
1) Create an empty stack and push root node to stack.
2) Do following while stack is not empty.
….a) Pop an item from stack and print it.
….b) Push right child of popped item to stack
….c) Push left child of popped item to stack
Right child is pushed before left child to make sure that left subtree is processed first.
 */
func PreOrder(node *TreeNode) {
	stack := new(StackLL)
	stack.Push(node)
	for !stack.IsEmpty() {
		currentNode := stack.Pop().(*TreeNode)
		if currentNode != nil {
			fmt.Printf("%v, ", currentNode.key)
			stack.Push(currentNode.right)
			stack.Push(currentNode.left)
		}
	}
}

func PostOrderRecursive(node *TreeNode) {
	if node == nil {
		return
	}

	PostOrderRecursive(node.left)
	PostOrderRecursive(node.right)
	fmt.Printf("%v, ", node.key)
}

/*
1) Create an empty stack and push the root node to stack
2) While stack is not empty,
   - check if the top of the stack is a leaf node, if yes is pop it & print it
   - add the left & right nodes under current node to stack
   - after adding each left & right node NULL the left/right pointers to avoid reprocessing
 */
func PostOrder(node *TreeNode) {
	stack := new(StackLL)
	currentNode := node
	stack.Push(currentNode)
	for !stack.IsEmpty() {
		currentNode = stack.Top().(*TreeNode)
		if currentNode.left == nil && currentNode.right == nil {
			currentNode = stack.Pop().(*TreeNode)
			fmt.Printf("%v, ", currentNode.key)
		} else {
			//add the left & right nodes under each parent to stack
			//after adding each left & right node NULL the left/right pointers
			if currentNode.right != nil {
				stack.Push(currentNode.right)
				currentNode.right = nil
			}
			if currentNode.left != nil {
				stack.Push(currentNode.left)
				currentNode.left = nil
			}
		}
	}
}
