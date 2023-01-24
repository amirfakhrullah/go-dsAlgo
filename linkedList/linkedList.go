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
	currNode := l.Head
	for currNode != nil {
		if currNode.GetNextNode() == nil {
			currNode.SetNextNode(nil)
		}
		currNode = currNode.GetNextNode()
	}
}

func (l *LinkedList) Remove(k string) {
	if l.Head == nil {
		return
	}
	prevNode := l.Head
	if prevNode.GetCurrentKey() == k {
		l.Head = prevNode.GetNextNode()
		return
	}
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