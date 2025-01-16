package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func main() {
	var myList linkedList
	myList.addNodeToEnd(1)
	myList.addNodeToEnd(2)
	myList.addNodeToEnd(3)
	myList.addNodeToEnd(4)
	myList.printList()

}

func (list *linkedList) addNodeToEnd(data int) {
	nn := &node{data, nil}
	if list.head == nil {
		list.head = nn
	} else {
		cn := list.head
		for cn.next != nil {
			cn = cn.next
		}
		cn.next = nn
	}
}

func (list *linkedList) printList() {
	cn := list.head
	for cn != nil {
		fmt.Printf("current node address: %p, current node data: %d, next node address: %p\n", cn, cn.data, cn.next)
		cn = cn.next
	}
}
