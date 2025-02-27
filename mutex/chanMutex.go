package mutex

type chanMutex struct {
	ch chan struct{}
}

func NewChanMutex() *chanMutex {
	return &chanMutex{ch : make(chan struct{}, 1)}
}

func (m *chanMutex) Lock() {
	m.ch <- struct{}{}
}

func (m *chanMutex) Unlock() {
	select {
	case <-m.ch:
	default:	
	}
}
