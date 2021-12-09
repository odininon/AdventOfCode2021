package utils

import (
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func NewPoint(s string) Point {
	parts := strings.Split(s, ",")

	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return Point{
		X: x,
		Y: y,
	}
}

func (p1 Point) Subtract(p2 Point) Point {
	return p1.Add(p2.Multiply(-1))
}

func (p1 Point) Multiply(i int) Point {
	return Point{
		X: i * p1.X,
		Y: i * p1.Y,
	}
}

func (p1 Point) Add(p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func (p1 Point) Sign() Point {
	return Point{
		X: Sign(p1.X),
		Y: Sign(p1.Y),
	}
}
