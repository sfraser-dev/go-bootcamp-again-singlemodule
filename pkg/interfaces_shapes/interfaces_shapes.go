package pkg_interfaces_shapes

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

// triangle and square structs
type triangle struct {
	base   int
	height int
}

type square struct {
	side int
}

// attach printArea() methods to both triangle and square promoting them both to a shape interface
func (tr triangle) getArea() float64 {
	return 0.5 * float64(tr.base) * float64(tr.height)
}

func (sq square) getArea() float64 {
	return float64(sq.side * sq.side)
}

// function that take a shape interface (ie: can accept a promoted square or a promoted triangle)
func area(sh shape) {
	fmt.Printf("area of %T is %.2f\n", sh, sh.getArea()) // %T is format specifier to variable type
}

func RunIt() {
	// create a square and a triangle
	sq := square{side: 10}
	tr := triangle{base: 10, height: 5}

	// print area of square and triangle (using function that takes a shape interface argument)
	area(sq)
	area(tr)
}
