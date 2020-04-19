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



func MergeTwoSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}

	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}

	newList := &LinkedList{
		head: &ListNode{
			next:  nil,
			value: nil,
		},
		length: 0,
	}

	cur := newList.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	for cur1 != nil && cur2 != nil {
		if cur1.value.(int) <= cur2.value.(int) {
			cur.next = cur1
			cur1 = cur1.next
		} else {
			cur.next = cur2
			cur2 = cur2.next
		}
		cur = cur.next
	}

	if cur1 != nil {
		cur.next = cur1
	} else if cur2 != nil {
		cur.next = cur2
	}

	return newList

}
