package main

import (
	"fmt"
)

func main() {
	simpleInterface()
	goPoly()
}

// Simpler interface
type Simpler interface {
	Get()
	Set(i int)
}

// Simple struct
type Simple struct {
	num int
}

// Get func for Simple struct
func (s *Simple) Get() int {
	return s.num
}

// Set func for Simple struct
func (s *Simple) Set(i int) int {
	s.num = i
	return i
}

func simpleInterface() {
	s := Simple{}
	s.Set(100)
	fmt.Printf("Current value %d\n", s.Get())
}

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

const (
	/* PI blabla */
	PI = 3.14159
)

func (c *Circle) Area() float32 {
	return PI * c.radius * c.radius
}


func goPoly() {

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	c := &Circle{4}
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q, c}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}