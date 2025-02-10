package queue

import (
	"errors"
	"sync"
)

type Queue[T any] struct {
	item []T
	mutex sync.Mutex
}

func New[T any]()(*Queue[T]){
	obj := &Queue[T]{
		item: make([]T,0),
	}
	return obj
}

func (q *Queue[T])IsEmpty()bool{
	defer q.mutex.Unlock()
	q.mutex.Lock()
	return len(q.item) <= 0
}

func (q *Queue[T])Pop()(T,error){
	ret,err := q.Top()

	defer q.mutex.Unlock()
	q.mutex.Lock()

	if(err != nil){
		return ret, err
	}

	q.item = q.item[1:]
	return ret, nil
}

func (q *Queue[T])Top()(T,error){	
	if(q.IsEmpty()){
		var ret T
		return ret,errors.New("Empty Queue")
	}
	defer q.mutex.Unlock()
	q.mutex.Lock()
	return q.item[0], nil
}

func (q *Queue[T])Push(element T){
	defer q.mutex.Unlock()
	q.mutex.Lock()
	
	q.item = append(q.item, element)
}