package hashtable

import (
	"errors"

	linkedlist "github.com/amirfakhrullah/go-dsAlgo/linkedList"
)

type HashTable struct {
	Array [20]*linkedlist.LinkedList
}

func InitHashTable() *HashTable {
	var array [20]*linkedlist.LinkedList
	for i := range array {
		array[i] = linkedlist.InitLinkedList()
	}
	return &HashTable{
		Array: array,
	}
}

func getHashCode(k string) int {
	sum := 0
	for _, v := range k {
		sum += int(v)
	}
	return sum % len(k)
}

func (h *HashTable) Insert(k string) error {
	hashCode := getHashCode(k)
	lists := h.Array[hashCode]
	isExists := lists.Find(k)
	if isExists {
		return errors.New(k + " already stored")
	}
	lists.Push(k)
	return nil
}

func (h *HashTable) Remove(k string) {
	hashCode := getHashCode(k)
	lists := h.Array[hashCode]
	lists.Remove(k)
}
