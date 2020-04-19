package linkedlist

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length int
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (n *ListNode) GetNext() *ListNode {
	return n.next
}

func (n *ListNode) GetValue() interface{} {
	return n.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   NewListNode(0),
		length: 0,
	}
}

func (l *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}

	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	l.length++
	return true
}

func (l *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil || p == l.head {
		return false
	}

	cur := l.head.next
	pre := l.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	l.length++
	return true
}

func (l *LinkedList) InsertToTail(v interface{}) bool {
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}

	return l.InsertAfter(cur, v)
}


func (l *LinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}


func (l *LinkedList) FindByIndex(index int) *ListNode {
	if index >= l.length {
		return nil
	}

	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := l.head.next
	pre := l.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	l.length--
	return true
}

// prints
func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

// O(n) : Time complexity
func (l *LinkedList) Reverse() {
	if l.head == nil || l.head.next == nil || l.head.next.next == nil {
		return
	}

	var pre *ListNode
	cur := l.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	l.head.next = pre
}
