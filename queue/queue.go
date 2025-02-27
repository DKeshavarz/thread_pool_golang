package queue

import (
	"errors"
	"thread_pool/mutex"
)

type Queue[T any] struct {
	item []T
	mtx  mutex.Mutex
}

func New[T any]() *Queue[T] {
	obj := &Queue[T]{
		item: make([]T, 0),
	}

	// obj.mtx = mutex.NewAtomicMutex()
	obj.mtx = mutex.NewChanMutex()

	return obj
}

func (q *Queue[T]) IsEmpty() bool {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	return len(q.item) <= 0
}

func (q *Queue[T]) Pop() (T, error) {
	var ret T

	q.mtx.Lock()
	defer q.mtx.Unlock()

	if len(q.item) == 0 {
		return ret, errors.New("empty queue")
	}

	ret = q.item[0]
	q.item = q.item[1:]
	return ret, nil
}

func (q *Queue[T]) Top() (T, error) {
	if q.IsEmpty() {
		var ret T
		return ret, errors.New("empty queue")
	}

	q.mtx.Lock()
	defer q.mtx.Unlock()

	return q.item[0], nil
}

func (q *Queue[T]) Push(element T) {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	q.item = append(q.item, element)
}
