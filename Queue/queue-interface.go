package queue

type Stack interface {
	Enqueue(v interface{})
	Dequeue() interface{}
}
