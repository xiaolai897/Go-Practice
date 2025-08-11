package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func paramRun(s Shape) (float64, float64) {
	return s.Area(), s.Perimeter()
}

func (e *Employee) PrintInfo() {
	fmt.Printf("employee id: %d, name: %s, age: %d\n", e.EmployeeID, e.Person.Name, e.Person.Age)
}

func main() {
	rectangle := Rectangle{Width: 2.4, Length: 3.3}
	circle := Circle{Radius: 2.5}
	rAre, rPer := paramRun(&rectangle)
	fmt.Printf("rectangle are: %.2f, rectangle perimeter: %.2f\n", rAre, rPer)
	cAre, cPer := paramRun(&circle)
	fmt.Printf("circle are: %.2f, circle perimeter: %.2f\n", cAre, cPer)

	e := Employee{EmployeeID: 10, Person: Person{Name: "Tom", Age: 23}}
	e.PrintInfo()
}
