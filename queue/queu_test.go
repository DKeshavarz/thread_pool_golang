package queue

import (
	"testing"
)

func TestQueue_IsEmpty(t *testing.T) {
	q := New[int]()

	if !q.IsEmpty() {
		t.Error("Expected queue to be empty initially")
	}

	q.Push(1)
	if q.IsEmpty() {
		t.Error("Expected queue not to be empty after pushing an element")
	}
}

func TestQueue_Push(t *testing.T) {
	q := New[int]()

	q.Push(1)
	if q.IsEmpty() {
		t.Error("Expected queue not to be empty after pushing an element")
	}

	q.Push(2)
	q.Push(3)
	if len(q.item) != 3 {
		t.Errorf("Expected queue length to be 3, got %d", len(q.item))
	}
}

func TestQueue_Top(t *testing.T) {
	q := New[int]()

	_, err := q.Top()
	if err == nil {
		t.Error("Expected error when calling Top on an empty queue")
	}

	q.Push(1)
	q.Push(2)

	top, err := q.Top()
	if err != nil {
		t.Errorf("Unexpected error when calling Top on a non-empty queue: %v", err)
	}
	if top != 1 {
		t.Errorf("Expected top element to be 1, got %d", top)
	}
}

func TestQueue_Pop(t *testing.T) {
	q := New[int]()

	_, err := q.Pop()
	if err == nil {
		t.Error("Expected error when calling Pop on an empty queue")
	}

	q.Push(1)
	q.Push(2)

	popped, err := q.Pop()
	if err != nil {
		t.Errorf("Unexpected error when calling Pop on a non-empty queue: %v", err)
	}
	if popped != 1 {
		t.Errorf("Expected popped element to be 1, got %d", popped)
	}
	if len(q.item) != 1 {
		t.Errorf("Expected queue length to be 1 after pop, got %d", len(q.item))
	}

	popped, err = q.Pop()
	if err != nil {
		t.Errorf("Unexpected error when calling Pop on a non-empty queue: %v", err)
	}
	if popped != 2 {
		t.Errorf("Expected popped element to be 2, got %d", popped)
	}
	if !q.IsEmpty() {
		t.Error("Expected queue to be empty after popping all elements")
	}
}

func TestQueue_Concurrency(t *testing.T) {
	q := New[int]()

	done := make(chan bool)
	go func() {
		q.Push(1)
		done <- true
	}()
	go func() {
		q.Push(2)
		done <- true
	}()

	<-done
	<-done

	if len(q.item) != 2 {
		t.Errorf("Expected queue length to be 2 after concurrent pushes, got %d", len(q.item))
	}

	go func() {
		_, _ = q.Pop()
		done <- true
	}()
	go func() {
		_, _ = q.Pop()
		done <- true
	}()

	<-done
	<-done

	if !q.IsEmpty() {
		t.Error("Expected queue to be empty after concurrent pops")
	}
}

func TestQueue_Concurrency_all(t *testing.T) {
	q := New[int]()

	expectedSize := 100000
	done := make(chan struct{}, expectedSize)
	for i := 0; i < expectedSize; i++ {
		go func(v int) {
			q.Push(v)
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < expectedSize; i++ {
		<-done
	}

	if len(q.item) != expectedSize {
		t.Errorf("Expected queue length to be %d after concurrent pushes, got %d", expectedSize, len(q.item))
	}

	for i := 0; i < expectedSize; i++ {
		go func() {
			_, _ = q.Pop()
			done <- struct{}{}
		}()
		go func() {
			_, _ = q.Top()
		}()
	}
	for i := 0; i < expectedSize; i++ {
		<-done
	}

	if !q.IsEmpty() {
		t.Errorf("Expected queue length to be %d after concurrent pushes, got %d", 0, len(q.item))
	}
}

func TestQueue_Concurrency_all_include_IsEmpty(t *testing.T) {
	q := New[int]()

	expectedSize := 1000000
	done := make(chan struct{}, expectedSize)
	for i := 0; i < expectedSize; i++ {
		go func() {
			q.Push(i)
			done <- struct{}{}
		}()
	}

	for i := 0; i < expectedSize; i++ {
		<-done
	}

	if len(q.item) != expectedSize {
		t.Errorf("Expected queue length to be %d after concurrent pushes, got %d", expectedSize, len(q.item))
	}

	for i := 0; i < expectedSize; i++ {
		go func() {
			_, _ = q.Pop()
			done <- struct{}{}
		}()
		go func() {
			_, _ = q.Top()
		}()
	}
	for i := 0; i < expectedSize; i++ {
		<-done
	}

	for i := 0; i < expectedSize; i++ {
		go func() {
			q.IsEmpty()
			done <- struct{}{}
		}()
	}
	for i := 0; i < expectedSize; i++ {
		<-done
	}
	if !q.IsEmpty() {
		t.Errorf("Expected queue length to be %d after concurrent pushes, got %d", 0, len(q.item))
	}
}
