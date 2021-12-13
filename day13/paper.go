package main

import (
	"AdventOfCode2021/utils"
	"fmt"
	"math"
)

type Paper struct {
	dots          map[utils.Point]struct{}
	width, height int
}

func newPaper(points []utils.Point) Paper {
	width, height := 0, 0
	dots := make(map[utils.Point]struct{})

	for _, point := range points {
		dots[point] = struct{}{}
		if point.X > width {
			width = point.X
		}

		if point.Y > height {
			height = point.Y
		}
	}

	return Paper{
		width:  width + 1,
		height: height + 1,
		dots:   dots,
	}
}

func (paper Paper) Draw() {
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt

	for point := range paper.dots {
		if point.X < minX {
			minX = point.X
		}

		if point.Y < minY {
			minY = point.Y
		}

		if point.X > maxX {
			maxX = point.X
		}

		if point.Y > maxY {
			maxY = point.Y
		}
	}

	fmt.Println()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, ok := paper.dots[utils.Point{X: x, Y: y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (paper Paper) FoldHorizontally(line int) Paper {
	dots := make(map[utils.Point]struct{})

	for point := range paper.dots {
		if point.Y < line {
			dots[point] = struct{}{}
		} else {
			y := point.Y
			newPoint := utils.Point{X: point.X, Y: y - ((y - line) * 2)}

			dots[newPoint] = struct{}{}
		}
	}

	return Paper{
		dots:   dots,
		width:  paper.width,
		height: paper.height - line - 1,
	}
}

func (paper Paper) FoldVertically(line int) Paper {
	dots := make(map[utils.Point]struct{})

	for point := range paper.dots {
		if point.X < line {
			dots[point] = struct{}{}
		} else {
			x := point.X
			newPoint := utils.Point{X: x - ((x - line) * 2), Y: point.Y}

			dots[newPoint] = struct{}{}
		}
	}

	return Paper{
		dots:   dots,
		width:  paper.width - line - 1,
		height: paper.height,
	}
}

func (paper Paper) RunCommand(command command) Paper {
	if command.direction == Horizontal {
		return paper.FoldHorizontally(command.point)
	} else {
		return paper.FoldVertically(command.point)
	}
}
