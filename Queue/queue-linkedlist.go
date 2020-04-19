package queue

import "fmt"

type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *LinkedListQueue) Enqueue(v interface{}) {
	node := &ListNode{
		val:  v,
		next: nil,
	}
	if q.tail == nil {
		q.tail = node
		q.head = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

func (q *LinkedListQueue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}
	v := q.head.val
	q.head = q.head.next
	q.length--
	return v
}

func (q *LinkedListQueue) Print() string {
	if q.head == nil {
		return "Queue is Empty!"
	}
	result := "head"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
