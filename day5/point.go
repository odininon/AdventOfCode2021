package main

import (
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func newPoint(s string) point {
	parts := strings.Split(s, ",")

	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return point{
		x: x,
		y: y,
	}
}

func (p1 point) Subtract(p2 point) point {
	return p1.Add(p2.Multiply(-1))
}

func (p1 point) Multiply(i int) point {
	return point{
		x: i * p1.x,
		y: i * p1.y,
	}
}

func (p1 point) Add(p2 point) point {
	return point{
		x: p1.x + p2.x,
		y: p1.y + p2.y,
	}
}

func (p1 point) Sign() point {
	return point{
		x: Sign(p1.x),
		y: Sign(p1.y),
	}
}
