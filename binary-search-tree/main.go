package main

import (
	"fmt"
	"sync"
)

type Item interface{}

type Node struct {
	key   int
	value Item
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
	lock *sync.RWMutex
}

func (bt *BinaryTree) Insert(key int, value Item) {
	bt.lock.Lock()
	defer bt.lock.Unlock()
	if bt.root == nil {
		bt.root = &Node{
			key:   key,
			value: value,
		}
	} else {
		newNode := &Node{
			key:   key,
			value: value,
		}
		insertNode(bt.root, newNode)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (bt *BinaryTree) InOrderTraverse(f func(Item)) {
	bt.lock.Lock()
	defer bt.lock.Unlock()
	inOrderTraverse(bt.root, f)
}

func inOrderTraverse(node *Node, f func(Item)) {
	if node != nil {
		inOrderTraverse(node.left, f)
		f(node.value)
		inOrderTraverse(node.right, f)
	}
}

func (bt *BinaryTree) PreOrderTraverse(f func(Item)) {
	bt.lock.Lock()
	defer bt.lock.Unlock()
	preOrderTraverse(bt.root, f)
}

func preOrderTraverse(node *Node, f func(Item)) {
	if node != nil {
		f(node.value)
		preOrderTraverse(node.left, f)
		preOrderTraverse(node.right, f)
	}
}

func (bt *BinaryTree) PostOrderTraverse(f func(Item)) {
	//bt.lock.Lock()
	//defer bt.lock.Unlock()
	postOrderTraverse(bt.root, f)
}

func postOrderTraverse(node *Node, f func(Item)) {
	if node != nil {
		postOrderTraverse(node.left, f)
		postOrderTraverse(node.right, f)
		f(node.value)
	}
}

func (bt *BinaryTree) RemoveNode(key int) {
	bt.root = removeNodeHelper(bt.root, key)
}

func removeNodeHelper(node *Node, key int) *Node {
	if node == nil {
		return node
	} else if key < node.key {
		node.left = removeNodeHelper(node.left, key)
	} else if key > node.key {
		node.right = removeNodeHelper(node.right, key)
	} else {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.right != nil {
			node.key = sucessor(node)
			node.right = removeNodeHelper(node.right, node.key)
		} else if node.left != nil {
			node.key = precessor(node)
			node.left = removeNodeHelper(node.left, node.key)
		}
	}
	return node
}

func sucessor(node *Node) int {
	for node.left != nil {
		node = node.left
	}
	return node.key
}

func precessor(node *Node) int {
	for node.right != nil {
		node = node.right
	}
	return node.key
}

// String prints a visual representation of the tree
func (bt *BinaryTree) String() {
	bt.lock.Lock()
	defer bt.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bt.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}

func main() {
	lock := sync.RWMutex{}
	bt := BinaryTree{
		lock: &lock,
	}
	fillTree(&bt)

	bt.RemoveNode(8)

	bt.InOrderTraverse(func(i Item) {
		fmt.Printf("Valor do nó: %s \n", i)
	})

	/* bt.PostOrderTraverse(func(i Item) {
		fmt.Printf("Valor do nó: %s \n", i)
	}) */

	bt.String()
}

func fillTree(bt *BinaryTree) {
	bt.Insert(8, "8")
	bt.Insert(4, "4")
	bt.Insert(10, "10")
	bt.Insert(2, "2")
	bt.Insert(6, "6")
	bt.Insert(1, "1")
	bt.Insert(3, "3")
	bt.Insert(5, "5")
	bt.Insert(7, "7")
	bt.Insert(9, "9")
}
