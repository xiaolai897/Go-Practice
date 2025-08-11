package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task func()

func oddNumber(r int) {
	for i := 1; i <= r; i++ {
		if i%2 != 0 {
			fmt.Printf("odd number: %d\n", i)
		}
	}
}

func evenNumber(r int) {
	for i := 1; i <= r; i++ {
		if i%2 == 0 {
			fmt.Printf("even number: %d\n", i)
		}
	}
}

func runTasks(tasks []Task) {
	for i, task := range tasks {
		go func(idex int, t Task) {
			start := time.Now()
			t()
			end := time.Since(start)
			fmt.Printf("task %d end, elapsed time: %v \n", idex, end)
		}(i, task)
	}
}

func main() {
	go oddNumber(10)
	go evenNumber(10)
	tasks := []Task{
		func() {
			time.Sleep(time.Duration(rand.Intn(2000)+500) * time.Millisecond)
		},
		func() {
			time.Sleep(time.Duration(rand.Intn(2000)+500) * time.Millisecond)
		},
		func() {
			time.Sleep(time.Duration(rand.Intn(2000)+500) * time.Millisecond)
		},
	}
	runTasks(tasks)
	time.Sleep(3 * time.Second)
}
