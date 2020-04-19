package stack

import "fmt"

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}


func (s *ArrayStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

func (s *ArrayStack) Push(v interface{}) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top++
	}

	if s.top > len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	topElement := s.data[s.top]
	s.top--
	return topElement
}

func (s *ArrayStack) IsEmpty() bool {
	return s.top < 0
}


func (s *ArrayStack) Print() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty!")
	} else {
		for i := s.top; i >= 0; i-- {
			fmt.Println(s.data[i])
		}
	}
}
