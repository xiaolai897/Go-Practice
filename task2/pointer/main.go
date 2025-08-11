package main

import "fmt"

func add(a *int) {
	*a += 10
}

func sliceMultiply(s *[]int) {
	for k := range *s {
		(*s)[k] *= 2
	}
}

func main() {
	fmt.Println("pointer!")
	a := 20
	fmt.Printf("add before: %d\n", a)
	add(&a)
	fmt.Printf("add after: %d\n", a)

	slice := []int{3, 5, 7, 9, 11, 13, 15}
	fmt.Printf("sliceMultiply before: %#v\n", slice)
	sliceMultiply(&slice)

	fmt.Printf("sliceMultiply after: %#v\n", slice)
}
