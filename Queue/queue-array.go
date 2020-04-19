package queue

import "fmt"

type ArrayQueue struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		data:     make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

func (q *ArrayQueue) Enqueue(v interface{}) bool {
	if q.tail == q.capacity {
		if q.head == 0 {
			return false
		}
		for i := q.head; i < q.tail; i++ {
			q.data[i-q.head] = q.data[i]
		}
		q.tail -= q.head
		q.head = 0
	}
	q.data[q.tail] = v
	q.tail++
	return true
}

func (q *ArrayQueue) Dequeue() interface{} {
	if q.head == q.tail {
		return nil
	}
	val := q.data[q.head]
	q.head++
	return val
}

func (q *ArrayQueue) Print() string {
	if q.head == q.tail {
		return "Queue is Empty!"
	}
	result := "head"
	for i := q.head; i <= q.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", q.data[i])
	}
	result += "<-tail"
	return result
}
