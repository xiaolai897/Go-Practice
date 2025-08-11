package main

import (
	"fmt"
	"sync"
)

func sendOnly(max int, ch chan<- int) {
	for i := 1; i <= max; i++ {
		fmt.Printf("send value: %d\n", i)
		ch <- i
	}
	close(ch)
}

func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("receive value: %d\n", v)
	}
}

func main() {
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	max := 100
	wg.Add(2)
	go func() {
		defer wg.Done()
		sendOnly(max, ch)
	}()

	go func() {
		defer wg.Done()
		receiveOnly(ch)
	}()
	wg.Wait()
	fmt.Println("程序结束")
}
