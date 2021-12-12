package main

import (
	"AdventOfCode2021/utils"
	"fmt"
)

type Grid struct {
	cells         [][]int
	width, height int
}

func (g Grid) Draw() {
	for x := 0; x < g.width+2; x++ {
		fmt.Printf("=")
	}

	fmt.Println()
	for y := 0; y < g.height; y++ {
		fmt.Printf("=")
		for x := 0; x < g.width; x++ {
			fmt.Print(g.cells[y][x])
		}
		fmt.Printf("=")
		fmt.Println()
	}

	for x := 0; x < g.width+2; x++ {
		fmt.Printf("=")
	}

	fmt.Println()

}

func (g Grid) Tick() (Grid, int) {
	directions := []utils.Point{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	newCells := make([][]int, g.height)

	for k := range newCells {
		newCells[k] = make([]int, g.width)
	}

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			newCells[y][x] = g.cells[y][x]
		}
	}

	cellsToFlash := make(map[utils.Point]struct{})
	cellsThatFlashed := make(map[utils.Point]struct{})

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if g.cells[y][x] == 9 {
				cellsToFlash[utils.Point{X: x, Y: y}] = struct{}{}
				newCells[y][x] = 9
			} else {
				newCells[y][x] = g.cells[y][x] + 1
			}
		}
	}

	for {
		if len(cellsToFlash) == 0 {
			break
		}

		for point := range cellsToFlash {
			for _, direction := range directions {
				newPoint := point.Add(direction)

				if newPoint.X < 0 || newPoint.Y < 0 || newPoint.X >= g.width || newPoint.Y >= g.height {
					continue
				}
				if newCells[newPoint.Y][newPoint.X] == 9 {
					if _, ok := cellsThatFlashed[newPoint]; !ok {
						cellsToFlash[newPoint] = struct{}{}
					}
				} else {
					newCells[newPoint.Y][newPoint.X] = newCells[newPoint.Y][newPoint.X] + 1
				}
			}
			cellsThatFlashed[point] = struct{}{}
			delete(cellsToFlash, point)
		}
	}

	for point := range cellsThatFlashed {
		newCells[point.Y][point.X] = 0
	}

	return Grid{cells: newCells, width: g.width, height: g.height}, len(cellsThatFlashed)
}
