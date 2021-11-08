package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList_Insert(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.InsertEnd(1)
	linkedList.InsertEnd(2)
	linkedList.InsertEnd(3)
	linkedList.Travese(func(node *Node) {
		fmt.Printf("Current Node Value Is %v \n", node.value)
	})
}
