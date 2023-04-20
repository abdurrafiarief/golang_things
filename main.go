package main

import (
	"fmt"
)

// A generic linked list node
type Node[T comparable] struct {
	data T
	next *Node[T]
}

// A generic linked list
type List[T comparable] struct {
	head *Node[T]
	size int
}

// Create a new node with the given data
func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{data: data, next: nil}
}

// Create a new empty list
func NewList[T comparable]() *List[T] {
	return &List[T]{head: nil, size: 0}
}

func InitializeNewFilledList[T comparable](data T) *List[T] {
	list := NewList[T]()
	list.Append(data)
	return list
}

func (l *List[T]) Append(data T) {
	var (
		newNode  *Node[T]
		currNode *Node[T]
	)
	// There is no head
	if l.head == nil {
		newNode = NewNode(data)
		l.head = newNode
	} else {
		// There is a head
		currNode = l.head
		for {
			if currNode.next == nil {
				newNode = NewNode(data)
				currNode.next = newNode
				break
			} else {
				currNode = currNode.next
			}
		}
	}
	l.size++
}

func (l *List[T]) Print() {
	current := l.head
	for current != nil {
		if current == l.head {
			fmt.Printf("%v", current.data)
		} else {
			fmt.Printf(" <- %v", current.data)
		}
		current = current.next
	}
	fmt.Println("")
}

func (l *List[T]) DeleteData(data T) {
	// Data is in head
	if l.head.data == data {
		l.head = l.head.next
		l.size--
		return
	}

	// Data is not in head
	current := l.head.next
	previous := l.head
	for current != nil {
		if current.data == data {
			previous.next = current.next
			return
		}
		previous = current
		current = current.next
	}
}

func main() {
	fmt.Println("Hello world")
	testList := InitializeNewFilledList[string]("Head")
	testList.Append("body2")
	testList.Append("body3")
	testList.Append("a")
	testList.Append("tail")
	testList.Print()
	testList.DeleteData("a")
	testList.Print()
	testList.DeleteData("Head")
	testList.Print()
}
