package mutex

type Mutex interface {
	Lock()
	Unlock()
}