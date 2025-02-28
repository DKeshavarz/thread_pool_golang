package mutex

import (
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
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
