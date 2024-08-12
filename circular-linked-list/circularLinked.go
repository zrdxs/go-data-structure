package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Insert(data int) *List {

	n := &Node{
		data: data,
		next: l.head,
	}

	if l.head == nil {
		n.next = n
		l.head = n
	} else {
		temp := l.head

		n.next = temp
		l.head = n
		temp.next = l.head
	}

	return l
}

func (l *List) PrintList() {

	list := l.head

	for i := 0; i < 3; i++ {
		fmt.Printf("Node: %d ", list.data)
		list = list.next
	}

}

func main() {

	list := &List{}

	list.Insert(20)
	list.Insert(15)
	list.Insert(10)

	list.PrintList()
}
