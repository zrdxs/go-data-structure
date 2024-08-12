package main

import "fmt"

type Node struct {
	next *Node
	data int
}

type List struct {
	head *Node
	len  int
}

func (l *List) Insert(data int) *List {

	n := &Node{
		data: data,
	}

	if l.head == nil {
		l.head = n
		l.len = 1
	} else {
		currNode := l.head

		n.next = currNode
		l.head = n
		l.len++
	}

	return l
}

func (l *List) InsertAfter(node *Node, data int) {

	if node == nil {
		fmt.Println("Null node informed!!")
		return
	}

	n := &Node{
		data: data,
		next: node.next,
	}

	node.next = n
	l.len++

	return
}

func (l *List) Append(data int) {

	n := &Node{
		data: data,
	}

	prt := l.head

	if l.head == nil {
		l.head = n
		l.len++
	} else {
		for prt.next != nil {
			prt = prt.next
		}

		prt.next = n
		l.len++
	}

}

func (l *List) DeleteNode(data int) {

	prt := l.head

	for prt.next != nil {
		if prt.next.data == data {
			fmt.Printf("\n Deleting node: %d \n", data)
			prev := prt
			prev.next = prt.next.next
			l.len--

			return
		}
		prt = prt.next
	}

	fmt.Println("Data not found!!")

}

func (l *List) PrintAllNodes() {
	prt := l.head

	for i := 0; i < l.len; i++ {
		fmt.Printf("Node: %d => ", prt.data)
		prt = prt.next
	}
}

func main() {
	l := &List{}

	l.Insert(5)
	l.Insert(4)
	l.Insert(2)
	l.InsertAfter(l.head, 3)
	l.Insert(1)
	l.Append(6)

	fmt.Printf("List Length: %d \n", l.len)
	l.PrintAllNodes()

	l.DeleteNode(6)
	l.PrintAllNodes()
}
