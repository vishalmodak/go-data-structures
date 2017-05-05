package main

import (
	"./ds"
	"container/list"
	"fmt"
)

func main() {
	//stackLLRun()
	//queueLLRun()
	//treeRun()
	//graphRun()
	linkedListRun()
}

func listsIntersect() {
	list1 := new(ds.LinkedList)
	list1.Insert(3)
	list1.Insert(6)
	list1.Print()
	list1.Insert(9)
	list1.Print()
	list1.Insert(11)
	list1.Insert(12)
	list1.Insert(15)
	list1.Print()

	list2 := new(ds.LinkedList)
	list2.Insert(6)
	list2.Insert(5)
	list2.Insert(7)
	list2.Insert(11)
	list2.Insert(12)
	list2.Insert(15)
	list2.Print()

	list1Idx := list1.Head()
	list2Idx := list2.Head()
	for i := 0; i < list.Size(); i++ {

	}

}

func graphRun() {
	graph := new(ds.Graph)
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddEdge("A", "B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddEdge("C", "D")
	//graph.AddEdge("C", "B")
	graph.AddEdge("A", "D")
	graph.AddNode("E")
	graph.AddEdge("B", "E")
	graph.AddEdge("D", "E")
	graph.AddNode("F")
	graph.AddEdge("D", "F")
	fmt.Printf("Graph size: %d\n", graph.Size())
	startNode, _ := graph.GetNode("A")
	//ds.DFS(startNode)
	ds.BFS(startNode)
}

func treeRun() {
	tree := new(ds.BST)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	fmt.Print("In-Order Traversal: ")
	ds.InOrderRecursive(tree.GetRoot())
	fmt.Print("\nPre-Order Traversal: ")
	ds.PreOrderRecursive(tree.GetRoot())
	fmt.Print("\nPost-Order Traversal: ")
	ds.PostOrderRecursive(tree.GetRoot())
}

func stackLLRun() {
	stack := new(ds.StackLL)
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
	fmt.Printf("pop() data: %s\n", stack.Pop())
}

func queueLLRun() {
	queue := new(ds.QueueLL)
	queue.Enqueue("x")
	queue.Enqueue("y")
	queue.Enqueue("z")
	queue.Print()

	fmt.Printf("Front of the queue is: %v\n", queue.GetFront().GetData())
	fmt.Printf("Back of the queue is: %v\n", queue.GetBack().GetData())

	fmt.Printf("deque() data: %s\n", queue.Deque())
	fmt.Printf("deque() data: %s\n", queue.Deque())
	queue.Print()

	queue.Clear()
	queue.Print()
	queue.Deque()

}
