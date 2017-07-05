package main

import (
	"./ds"
	"fmt"
)

func main() {
	//stackLLRun()
	//queueLLRun()
	//treeRun()
	//graphRun()
	graphCycleRun()
}

/*
		 C
		 |
	A -> D -> E
	 \	 |   /
	  \	 F  /
	   \   /
	    \ /
	     B
 */
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

func graphCycleRun() {
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
	//graph.AddEdge("D", "A")
	fmt.Printf("Graph size: %d\n", graph.Size())
	startNode, _ := graph.GetNode("A")
	fmt.Printf("Graph has Cycle is %t", ds.HasCycle(startNode))
	//ds.DFS(startNode)
}

/*
               10
       2               15
   1       5       11      17
*/
func treeRun() {
	tree := new(ds.BST)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(11)
	tree.Insert(17)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	fmt.Print("In-Order Traversal (Recursive): ")
	ds.InOrderRecursive(tree.GetRoot())
	fmt.Print("\nIn-Order Traversal: ")
	ds.InOrder(tree.GetRoot())
	fmt.Print("\nPre-Order Traversal(Recursive): ")
	ds.PreOrderRecursive(tree.GetRoot())
	fmt.Print("\nPre-Order Traversal: ")
	ds.PreOrder(tree.GetRoot())
	fmt.Print("\nPost-Order Traversal(Recursive): ")
	ds.PostOrderRecursive(tree.GetRoot())
	fmt.Print("\nPost-Order Traversal: ")
	ds.PostOrder(tree.GetRoot())
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
