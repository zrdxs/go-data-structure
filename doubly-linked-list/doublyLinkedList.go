package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	data int
}

type List struct {
	head *Node
}

func (l *List) Insert(data int) *List {
	n := &Node{
		data: data,
	}

	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head.prev = n

		l.head = n
	}

	return l
}

func (l *List) InsertAfter(node *Node, data int) {

	n := &Node{
		data: data,
	}

	if node.next != nil {
		nextNode := node.next
		nextNode.prev = n

		n.next = nextNode
		n.prev = node
		node.next = n
	} else {
		node.next = n
		n.prev = node
	}

}

func (list *List) InsertBefore(nextNode *Node, data int) {

	n := &Node{
		data: data,
	}

	prevNode := nextNode.prev

	prevNode.next = n
	n.prev = prevNode

	n.next = nextNode
	nextNode.prev = n

}

func (l *List) Append(data int) {

	prt := l.head

	n := &Node{
		data: data,
	}

	for prt.next != nil {
		if prt.next.next == nil {
			prt.next.next = n
			n.prev = prt.next
			return
		}

		prt = prt.next
	}

}

func (l *List) PrintNodes() {

	prt := l.head

	for prt.next != nil {
		if prt.prev != nil {
			fmt.Printf("Node: %d - Prev: %d \n", prt.data, prt.prev.data)
		}
		fmt.Printf("Node: %d \n", prt.data)
		prt = prt.next
	}
}

func main() {

	l := &List{}

	l.Insert(10)
	l.Insert(15)
	l.Insert(20)
	l.Insert(25)
	l.Append(5)
	l.Append(4)

	l.PrintNodes()
}
