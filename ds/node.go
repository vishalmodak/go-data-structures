package ds

type Node struct {
	data interface{}
	next *Node
}

func (n *Node) GetData() interface{} {
	return n.data
}

func (n *Node) Next() *Node {
	return n.next
}
