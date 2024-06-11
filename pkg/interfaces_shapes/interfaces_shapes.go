package pkg_shapes

import (
	"fmt"
)

type shape interface {
	printArea()
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
func (t triangle) printArea() {
	fmt.Printf("area of %T is %.2f\n", t, (0.5 * float64(t.base) * float64(t.height)))
}

func (s square) printArea() {
	fmt.Printf("area of %T is %.2f\n", s, float64(s.side*s.side))
}

// function that take a shape interface (ie: can accept a promoted square or a promoted triangle)
func printArea(sh shape) {
	sh.printArea()
}

func RunIt() {
	s := square{side: 10}
	t := triangle{base: 10, height: 5}

	printArea(s)
	printArea(t)
}
