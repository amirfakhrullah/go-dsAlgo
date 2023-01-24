package linkedlist

import "github.com/amirfakhrullah/go-dsAlgo/node"

type LinkedList struct {
	Head *node.Node
}

func InitLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

func (l *LinkedList) Push(k string) {
	newNode := &node.Node{
		Key:  k,
		Next: l.Head,
	}
	l.Head = newNode
}

func (l *LinkedList) Pop() {
	if l.Head == nil {
		return
	}
	if l.Head.GetNextNode() == nil {
		l.Head = nil
		return
	}
	prevNode := l.Head
	for prevNode.GetNextNode() != nil {
		if prevNode.GetNextNode().GetNextNode() == nil {
			prevNode.SetNextNode(nil)
		}
		prevNode = prevNode.GetNextNode()
	}
}

func (l *LinkedList) Remove(k string) {
	if l.Head == nil {
		return
	}
	if l.Head.GetCurrentKey() == k {
		l.Head = l.Head.GetNextNode()
		return
	}
	prevNode := l.Head
	for prevNode.GetNextNode() != nil {
		if prevNode.GetNextNode().GetCurrentKey() == k {
			prevNode.SetNextNode(prevNode.GetNextNode().GetNextNode())
			return
		}
		prevNode = prevNode.GetNextNode()
	}
}

func (l *LinkedList) Find(k string) bool {
	if l.Head == nil {
		return false
	}
	currNode := l.Head
	for currNode != nil {
		if currNode.GetCurrentKey() == k {
			return true
		}
		currNode = currNode.GetNextNode()
	}
	return false
}