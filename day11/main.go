package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"log"
	"strconv"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	var cells [][]int
	if *useTestInputs {
		cells = [][]int{
			{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
			{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
			{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
			{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
			{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
			{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
			{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
			{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
			{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
			{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
		}

	} else {
		lines, err := utils.ReadInputFile(11)

		if err != nil {
			log.Fatal(err)
		}

		height := len(lines)
		width := len(lines[0])

		cells = make([][]int, height)

		for k := range cells {
			cells[k] = make([]int, width)
		}

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				char := string(lines[y][x])
				number, _ := strconv.Atoi(char)

				cells[y][x] = number
			}
		}
	}

	grid := Grid{
		cells:  cells,
		width:  len(cells[0]),
		height: len(cells),
	}

	utils.PrintDayResultsWithDuration(11, part1(grid), part2(grid))
}

func part1(grid Grid) utils.ResultWithTime {
	flashes := 0
	newGrid := grid

	for i := 0; i < 100; i++ {
		newFlashes := 0
		newGrid, newFlashes = newGrid.Tick()
		flashes += newFlashes
	}

	return utils.ResultWithTime{
		Value: flashes,
	}
}

func part2(grid Grid) utils.ResultWithTime {
	newGrid := grid
	flashes := 0
	steps := 0

	for {
		steps += 1
		newGrid, flashes = newGrid.Tick()

		if flashes == newGrid.width*newGrid.height {
			break
		}
	}

	return utils.ResultWithTime{
		Value: steps,
	}
}
