package mutex

import (
    "fmt"
    "sync"
)

func TestMutex() {
    mtx := NewChanMutex()
    // mtx := NewAtomicMutex()

	count := 0
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
			defer wg.Done()
			defer mtx.Unlock() 
            mtx.Lock() 
            count++
            fmt.Printf("Goroutine %d after op : count = %d\n", id, count)
        }(i)
    }

    wg.Wait()
}
