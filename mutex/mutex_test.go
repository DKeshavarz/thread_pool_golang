package mutex

import (
	"sync"
	"testing"
)

func TestAtomicMutex(t *testing.T) {
    mtx := NewAtomicMutex()

	count, finalCount:= 0, 10000
    var wg sync.WaitGroup
    for i := 0; i < finalCount; i++ {

        wg.Add(1)
        go func(id int) {
			defer wg.Done()
			defer mtx.Unlock() 

            mtx.Lock() 
            count++
        }(i)
    }
    wg.Wait()
    if finalCount != count {
        t.Fatalf("expected %d, but got %d", finalCount, count)
    }
}

func TestChanMutex(t *testing.T) {
    mtx := NewChanMutex()

	count, finalCount:= 0, 10000
    var wg sync.WaitGroup
    for i := 0; i < finalCount; i++ {

        wg.Add(1)
        go func(id int) {
			defer wg.Done()
			defer mtx.Unlock() 

            mtx.Lock() 
            count++
        }(i)
    }
    wg.Wait()
    if finalCount != count {
        t.Fatalf("expected %d, but got %d", finalCount, count)
    }
}
