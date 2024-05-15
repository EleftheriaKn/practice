package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(pointer *point, x int, y int) {
	pointer.x = x
	pointer.y = y
}

func main() {
	points := point{}
	points.x = 42
	points.y = 21
	setPoint(&points, points.x, points.y)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune('4')
	z01.PrintRune('2')
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune('2')
	z01.PrintRune('1')
	z01.PrintRune('\n')
}
