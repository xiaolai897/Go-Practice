package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func add(count *int) {
	for i := 0; i < 1000; i++ {
		*count++
	}
}

func atomicAdd(count *int64) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(count, 1)
	}
}

func main() {
	count := 0
	var countAtomic int64
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			add(&count)
			defer mu.Unlock()
			defer wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			atomicAdd(&countAtomic)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("count: %d\n", count)
	fmt.Printf("atomic count: %d\n", countAtomic)
}
