package my_shapes

import (
	"fmt"
)

type shape interface {
	printArea()
}

type triangle struct {
	base   int
	height int
}

func (t triangle) printArea() {
	fmt.Printf("area of %T is %.2f\n", t, (0.5 * float64(t.base) * float64(t.height)))
}

type square struct {
	side int
}

func (s square) printArea() {
	fmt.Printf("area of %T is %.2f\n", s, float64(s.side*s.side))
}

func printArea(sh shape) {
	sh.printArea()
}

func RunIt() {
	s := square{side: 10}
	t := triangle{base: 10, height: 5}

	printArea(s)
	printArea(t)
}
