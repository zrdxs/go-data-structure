package main

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{left: nil, right: nil, data: data}
	} else {
		t.root.insert(data)
	}

	return t
}

func (n *BinaryNode) insert(data int64) {
	if data < n.data {
		if n.left == nil {
			n.left = &BinaryNode{left: nil, right: nil, data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{left: nil, right: nil, data: data}
		} else {
			n.right.insert(data)
		}
	}
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func printNodeByLayer(root *BinaryNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var nodes []*BinaryNode
	var node *BinaryNode
	nodes = append(nodes, root)
	idx := 0

	for len(nodes) != 0 {

		qtdNodes := len(nodes)
		res = append(res, []int{})

		for i := 0; i < qtdNodes; i++ {
			node, nodes = nodes[0], nodes[1:]
			res[idx] = append(res[idx], int(node.data))
			if node.left != nil {
				nodes = append(nodes, node.left)
			}
			if node.right != nil {
				nodes = append(nodes, node.right)
			}
		}
		idx++
	}

	return res
}

func main() {
	tree := &BinaryTree{}
	tree.insert(20).insert(5).insert(-10)

	valuesByNode := printNodeByLayer(tree.root)
	fmt.Println(valuesByNode)
	//print(os.Stdout, tree.root, 0, 'M')
}
