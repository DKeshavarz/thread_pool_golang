package mutex

import "sync/atomic"

type atomicMutex struct {
	s int32
}

func NewAtomicMutex() *atomicMutex {
	return &atomicMutex{s: 0}
}

func (m *atomicMutex) Lock() {
	for !atomic.CompareAndSwapInt32(&m.s, 0, 1) {
		// Loop until the lock is acquired
	}
}

func (m *atomicMutex) Unlock() {
	atomic.CompareAndSwapInt32(&m.s, 1, 0)
}
