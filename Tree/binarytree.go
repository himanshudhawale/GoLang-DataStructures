package main

import "fmt"

type Node struct {
	Value    int
	Left  *Node
	Right *Node
}

func PreOrder(node *Node) {
	if node != nil {
		fmt.Printf("value=%d \n", node.Value)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

func InfixOrder(node *Node) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("value=%d \n", node.Value)
		InfixOrder(node.Right)
	}
}

func PostOrder(node *Node) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("value=%d \n", node.Value)
	}
}

func main() {

	root := &Node{
		No:   1,
	}

	left1 := &Node{
		No:   2,
	}

	node10 := &Node{
		No:   10,
	}

	node12 := &Node{
		No:   12,
	}
	left1.Left = node10
	left1.Right = node12

	right1 := &Node{
		No:   3,
	}

	root.Left = left1
	root.Right = right1

	right2 := &Node{
		No:   4,
	}
	right1.Right = right2

	PostOrder(root)
}
