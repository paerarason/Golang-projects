package main

import "fmt"

//NODE
type Node struct {
	data int
	next *Node
}

// HEAD node of the LINKEDLIST
type List struct {
	head *Node
}

func add(l *List, value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

/*
func remove_first(l *List) {
	l.head = l.head.next

}


*/
func remove_last(l *List) {

	curr := l.head
	for curr.next.next != nil {
		curr = curr.next
	}
	curr.next = nil
}

func Printlink(l *List) {
	temp := l.head
	for temp.next != nil {
		fmt.Println(temp.data, " ")
		temp = temp.next
	}
}
