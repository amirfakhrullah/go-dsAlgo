package node

type Node struct {
	Key  string
	Next *Node
}

func (n *Node) GetCurrentKey() string {
	return n.Key
}

func (n *Node) GetNextNode() *Node {
	return n.Next
}

func (n *Node) SetNextNode(nextNode *Node) {
	n.Next = nextNode
}
