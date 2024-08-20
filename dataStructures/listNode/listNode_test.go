package listNode

import (
	"fmt"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	singlyLinkedList := NewSinglyLinkedList[int]()
	singlyLinkedList.AddFirst(1)
	singlyLinkedList.Print()
	fmt.Println("*********")

	singlyLinkedList.AddLast(3)
	singlyLinkedList.Print()
	fmt.Println("*********")

	singlyLinkedList.Add(1, 2)
	singlyLinkedList.Print()
	fmt.Println("*********")

	singlyLinkedList.Get(1)
	singlyLinkedList.Print()
	fmt.Println("*********")

	singlyLinkedList.Remove(1)
	singlyLinkedList.Print()
	fmt.Println("*********")

	singlyLinkedList.Get(1)
	singlyLinkedList.Print()
	fmt.Println("*********")
}
