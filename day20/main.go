package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

type Grid struct {
	cells    map[utils.Point]struct{}
	tracking string
}

func newGrid() Grid {
	return Grid{cells: make(map[utils.Point]struct{})}
}

func (g Grid) Draw() {
	minX := math.MaxInt
	minY := math.MaxInt

	maxX := math.MinInt
	maxY := math.MinInt

	for point := range g.cells {
		if point.X < minX {
			minX = point.X
		}

		if point.X > maxX {
			maxX = point.X
		}

		if point.Y < minY {
			minY = point.Y
		}

		if point.Y > maxY {
			maxY = point.Y
		}
	}

	onCharacter := g.tracking
	var offCharacter string

	if onCharacter == "#" {
		offCharacter = "."
	} else {
		offCharacter = "#"
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if g.IsOn(utils.Point{X: x, Y: y}) {
				fmt.Printf("%v", onCharacter)
			} else {
				fmt.Printf("%v", offCharacter)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) IsOn(point utils.Point) bool {
	_, isOn := g.cells[point]
	return isOn
}

func (g Grid) Enhance(encoding string) Grid {
	newGrid := newGrid()
	checked := make(map[utils.Point]struct{})

	if g.tracking == "#" && string(encoding[0]) == "#" {
		newGrid.tracking = "."
	} else if g.tracking == "." && string(encoding[len(encoding)-1]) == "." {
		newGrid.tracking = "#"
	} else {
		newGrid.tracking = g.tracking
	}

	for point := range g.cells {
		for y := point.Y - 3; y <= point.Y; y++ {
			for x := point.X - 3; x <= point.X; x++ {
				changedPoint := utils.Point{X: x + 1, Y: y + 1}
				if _, isChecked := checked[changedPoint]; isChecked {
					continue
				}

				binString := ""

				for ny := 0; ny < 3; ny++ {
					for nx := 0; nx < 3; nx++ {
						if g.IsOn(utils.Point{X: x + nx, Y: y + ny}) && g.tracking == "#" {
							binString += "1"
						} else if !g.IsOn(utils.Point{X: x + nx, Y: y + ny}) && g.tracking == "." {
							binString += "1"
						} else {
							binString += "0"
						}
					}
				}

				value, _ := strconv.ParseInt(binString, 2, 64)

				newValue := string(encoding[value])
				checked[changedPoint] = struct{}{}

				if newValue == newGrid.tracking {
					newGrid.cells[changedPoint] = struct{}{}
				}
			}
		}
	}

	return newGrid
}

func main() {
	var lines []string
	var encodingString string
	if *useTestInputs {
		encodingString = "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"
		lines = []string{
			"#..#.",
			"#....",
			"##..#",
			"..#..",
			"..###",
		}
	} else {
		lines, encodingString, _ = newGridReader("./inputs/day20.txt")
	}

	grid := newGrid()
	grid.tracking = "#"

	for y, line := range lines {
		for x, c := range line {
			if string(c) == grid.tracking {
				grid.cells[utils.Point{X: x, Y: y}] = struct{}{}
			}
		}
	}

	utils.PrintDayResultsWithDuration(20, part1(grid, encodingString), part2(grid, encodingString))
}

func part1(grid Grid, encodingString string) utils.ResultWithTime {
	t1 := time.Now()
	newGrid := grid

	for i := 0; i < 2; i++ {
		newGrid = newGrid.Enhance(encodingString)
	}

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    len(newGrid.cells),
		Duration: diff,
	}
}

func part2(grid Grid, encodingString string) utils.ResultWithTime {
	t1 := time.Now()
	newGrid := grid

	for i := 0; i < 50; i++ {
		newGrid = newGrid.Enhance(encodingString)
	}

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    len(newGrid.cells),
		Duration: diff,
	}
}
