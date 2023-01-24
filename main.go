package main

import (
	"fmt"

	"github.com/amirfakhrullah/go-dsAlgo/hashTable"
)

func main() {
	testStr := []string{"amir", "rima", "omar", "roma", "amrf", "fakhrullah"}

	table := hashtable.InitHashTable()

	for _, str := range testStr {
		if err := table.Insert(str); err != nil {
			panic(err.Error())
		}
	}
	// table.Remove(str2)

	for i, v := range table.Array {
		fmt.Println(i)
		currNode := v.Head
		for currNode != nil {
			fmt.Println(currNode.GetCurrentKey())
			currNode = currNode.GetNextNode()
		}
	}
}
