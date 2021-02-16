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

// Shaper interface
type Shaper interface {
	Area() float32
}

// Square type
type Square struct {
	side float32
}

// Area for Square
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

// Rectangle type
type Rectangle struct {
	length, width float32
}

// Area for Rectangle
func (r Rectangle) Area() float32 {
	return r.length * r.width
}

// Circle type
type Circle struct {
	radius float32
}

const (
	// PI blabla
	PI = 3.14159
)

// Area ...
func (c *Circle) Area() float32 {
	return PI * c.radius * c.radius
}

func (c *Circle) String() string {
	return fmt.Sprintf("A circle has radius of %f", c.radius)
}

func goPoly() {

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	c := &Circle{4}
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q, c}
	fmt.Println("Looping through shapes for area ...")
	for n, s := range shapes {
		switch s.(type) {
		case *Square:
			fmt.Printf("Squre %T\n", s)
		case *Rectangle:
			fmt.Printf("Rectangle %T\n", s)
		case *Circle:
			fmt.Printf("Circle %T\n", s)
		default:
			fmt.Printf("No shape found.\n")
		}
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
