package stack

import "fmt"

type node struct {
	next *node
	val  interface{}
}

type LinkedListStack struct {
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		topNode: nil,
	}
}

func (s *LinkedListStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.topNode.val
}

func (s *LinkedListStack) Push(v interface{}) {
	s.topNode = &node{
		next: s.topNode,
		val:  v,
	}
}

func (s *LinkedListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	val := s.topNode.val
	s.topNode = s.topNode.next
	return val
}

//Check if empty

func (s *LinkedListStack) IsEmpty() bool {
	return s.topNode == nil
}

func (s *LinkedListStack) Print() {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty!")
	} else {
		cur := s.topNode
		for nil != cur {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
