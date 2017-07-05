package ds

import (
	"fmt"
)

type GraphNode struct {
	data      interface{}
	neighbors []*GraphNode
	visited   bool
}

type Graph struct {
	nodes []*GraphNode
}

func (gn *GraphNode) hasNeighbors() bool {
	return len(gn.neighbors) > 0
}

func (g *Graph) isEmpty() bool {
	return len(g.nodes) == 0
}

func (g *Graph) Size() int {
	return len(g.nodes)
}

func (g *Graph) AddNode(value interface{}) {
	if g.isEmpty() {
		g.nodes = make([]*GraphNode, len(g.nodes))
	}
	g.nodes = append(g.nodes, &GraphNode{value, nil, false})
}

func (g *Graph) GetNode(value interface{}) (*GraphNode, bool) {
	for i := 0; i < len(g.nodes); i++ {
		if g.nodes[i].data == value {
			return g.nodes[i], true
		}
	}
	return nil, false
}

func (g *Graph) AddEdge(srcValue interface{}, dstValue interface{}) bool {
	srcNode, found := g.GetNode(srcValue)
	if !found {
		return false
	}
	dstNode, found := g.GetNode(dstValue)
	if !found {
		return false
	}

	if !srcNode.hasNeighbors() {
		srcNode.neighbors = make([]*GraphNode, len(srcNode.neighbors))
	}
	srcNode.neighbors = append(srcNode.neighbors, dstNode)
	return true
}

func DFS(startNode *GraphNode) {
	stack := new(StackLL)
	stack.Push(startNode)
	for !stack.IsEmpty() {
		//fmt.Println("===============")
		currentNode := stack.Pop().(*GraphNode)
		//fmt.Printf("Current node: %v\n", currentNode.data)
		if !currentNode.visited {
			for i := 0; currentNode.hasNeighbors() && i < len(currentNode.neighbors); i++ {
				neighbor := currentNode.neighbors[i]
				//fmt.Printf("Adding Neighbor: %v\n", neighbor)
				stack.Push(neighbor)
			}
			currentNode.visited = true
			//fmt.Printf("Marking node: %v\n", currentNode)
		}
	}
}

func BFS(startNode *GraphNode) {
	queue := new(QueueLL)
	startNode.visited = true
	queue.Enqueue(startNode)
	for !queue.isEmpty() {
		fmt.Printf("Queue size: %v\n", queue.size)
		currentNode := queue.Deque().(*GraphNode)
		fmt.Printf("Current node: %v\n", currentNode)
		for i := 0; currentNode.hasNeighbors() && i < len(currentNode.neighbors); i++ {
			neighbor := currentNode.neighbors[i]
			if !neighbor.visited {
				neighbor.visited = true
				//fmt.Printf("Marking neighbor: %v\n", neighbor)
			}
			//fmt.Printf("Adding Neighbor: %v\n", neighbor)
			queue.Enqueue(neighbor)
		}
	}
}

func HasCycle(startNode *GraphNode) bool {
	visitedList := new(LinkedList)
	stack := new(StackLL)
	stack.Push(startNode)
	visitedList.Insert(startNode)
	for !stack.IsEmpty() {
		currentNode := stack.Pop().(*GraphNode)
		fmt.Printf("Current: %s\n", currentNode.data)
		fmt.Printf("VisitedList: %s\n", visitedList.String())
		if !currentNode.visited {
			for i := 0; currentNode.hasNeighbors() && i < len(currentNode.neighbors); i++ {
				neighbor := currentNode.neighbors[i]
				fmt.Printf("Neighbor: %v\n", neighbor)
				if visitedList.Contains(neighbor) && neighbor.visited {
					return true
				}
				stack.Push(neighbor)
			}
			currentNode.visited = true
			visitedList.Insert(currentNode)
			//fmt.Printf("Marking node: %v\n", currentNode)
		}
	}
	return false
}
