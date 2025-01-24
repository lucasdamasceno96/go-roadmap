package main

import (
	"fmt"
)

type Node struct {
	value string
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) insert(value string) {
	node := &Node{value: value}

	if l.head == nil {
		l.head = node
	} else {
		last := l.head
		for last.next != nil {
			last = last.next
		}
		last.next = node
	}
}

func (l *LinkedList) display() {
	node := l.head
	for node != nil {
		fmt.Printf("%s -> ", node.value)
		node = node.next
	}
	fmt.Println("nil")

}
func (l *LinkedList) contains(value string) bool {
	current := l.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	list := LinkedList{}

	list.insert("Sneaker: Jordan 1")
	list.insert("Sneaker: Jordan 2")
	list.insert("Sneaker: Jordan 3")
	list.insert("Sneaker: Jordan 4")
	list.display()
	list.contains("Sneaker: Jordan 3")
}
