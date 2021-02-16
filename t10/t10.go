package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

type st1 struct {
	i1  int
	f1  float32
	str string
}

// type Person struct {
// 	firstName	string
// 	lastName	string
// }

// Address the address class
type Address struct {
	// Address
	addr string
}

// VCard the name card
type VCard struct {
	// "Name card"
	firstName string
	lastName  string
	address   *Address
	birthday  string
	photo     string
}

type anonySt struct {
	int
	string
}

type employee struct {
	name   string
	salary int
}

func main() {
	func1()
	vcard()
	func2()
	ee := employee{"Molly", 55}
	fmt.Printf("Current salary is %d\n", ee.giveRaise((88)))
	point := Point{3, 4}
	fmt.Printf("Abs is %d, Scaled is %v\n", point.Abs(), point.Scale(5))
	e2 := &Employee{100, Person{"Leo", "Cao", Base{1}}}
	fmt.Printf("%d\n", e2.Id())
	fmt.Printf("%v\n", e2)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}

func func1() {
	ms := new(st1)
	ms.i1 = 10
	ms.f1 = 3.14
	ms.str = "haha"

	fmt.Printf("%d\t%f\t%s\n", ms.i1, ms.f1, ms.str)

	ms1 := &st1{5, 6.18, "hehe"}
	fmt.Printf("%d\t%f\t%s\n", ms1.i1, ms1.f1, ms1.str)
}

func vcard() {
	var myCard VCard
	myCard = VCard{"Felix", "Cao", &Address{"4554 Deerfield"}, "1/1/2", "photo"}
	fmt.Printf("%v\n", myCard)

	yourCard := NewVCard("You", "Are", &Address{"Nowhere"})
	fmt.Printf("%v\n", *yourCard)
}

// NewVCard Create new VCard struct and return pointer
func NewVCard(fn string, ln string, addr *Address) *VCard {
	return &VCard{fn, ln, addr, "1/1/3", "photo2"}
}

func func2() {
	as := anonySt{5, "my name"}
	fmt.Printf("%d %s\n", as.int, as.string)
}

func (e *employee) giveRaise(inc int) int {
	e.salary = e.salary + inc
	return e.salary
}

// Engine the engine clas
type Engine interface {
	Start()
	Stpo()
}

// Car the car class
type Car struct {
	Engine
}

// Point some test class
type Point struct {
	x, y int
}

// Abs return abs
func (p *Point) Abs() int {
	return int(math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2)))
}

// Scale return scaled point
func (p *Point) Scale(factor int) *Point {
	return &Point{p.x * factor, p.y * factor}
}

// Base xxx
type Base struct {
	id int
}

// Id xxx
func (b *Base) Id() int {
	return b.id
}

// Person person
type Person struct {
	FirstName, LastName string
	Base
}

// Employee employee
type Employee struct {
	salary int
	Person
}

func (e *Employee) String() string {
	return "The employee with id " + strconv.Itoa(e.Id())
}
