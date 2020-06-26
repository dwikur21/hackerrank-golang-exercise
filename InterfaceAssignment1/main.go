package main

import "fmt"

// Shape interface
type Shape interface {
	printArea()
}

func (t Triangle) printArea() {
	fmt.Printf("Triangle's area is : %v\n", t.getArea())
}

func (sq Square) printArea() {
	fmt.Printf("Square's area is : %v\n", sq.getArea())
}

// Triangle struct
type Triangle struct {
	height float64
	base   float64
}

// Square struct
type Square struct {
	sideLength float64
}

func (t Triangle) getArea() float64 {
	return float64(0.5 * t.base * t.height)
}

func (sq Square) getArea() float64 {
	return float64(sq.sideLength * sq.sideLength)
}

func main() {
	t := Triangle{height: 10, base: 6}
	sq := Square{10}

	s := Shape(t)
	s.printArea()

	s = Shape(sq)
	s.printArea()
}
