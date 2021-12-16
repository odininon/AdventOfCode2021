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

type Graph struct {
	cells         map[utils.Point]int
	width, height int
}

func main() {
	cells := make(map[utils.Point]int)
	width, height := 0, 0

	var lines []string
	if *useTestInputs {
		lines = []string{
			"1163751742",
			"1381373672",
			"2136511328",
			"3694931569",
			"7463417111",
			"1319128137",
			"1359912421",
			"3125421639",
			"1293138521",
			"2311944581",
		}
	} else {
		lines, _ = utils.ReadInputFile(15)
	}

	for y, line := range lines {
		for x, char := range line {
			number, _ := strconv.Atoi(string(char))
			cells[utils.Point{X: x, Y: y}] = number
		}
	}
	height = len(lines)
	width = len(lines[0])

	graph := Graph{cells: cells, width: width, height: height}

	utils.PrintDayResultsWithDuration(15, part1(graph), part2(graph))
}

func part1(graph Graph) utils.ResultWithTime {
	t1 := time.Now()

	start := utils.Point{X: 0, Y: 0}
	end := utils.Point{X: graph.width - 1, Y: graph.height - 1}

	riskTotal := safestPath(graph, start, end)

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    riskTotal,
		Duration: diff,
	}
}

func part2(graph Graph) utils.ResultWithTime {
	t1 := time.Now()

	newGraph := grow(graph, 5)

	start := utils.Point{X: 0, Y: 0}
	end := utils.Point{X: newGraph.width - 1, Y: newGraph.height - 1}

	riskTotal := safestPath(newGraph, start, end)

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    riskTotal,
		Duration: diff,
	}
}

func grow(graph Graph, size int) Graph {
	newCells := make(map[utils.Point]int)

	for y := 0; y < graph.height; y++ {
		for x := 0; x < graph.width; x++ {
			point := utils.Point{X: x, Y: y}
			originalValue := graph.cells[point]
			newCells[point] = originalValue
		}
	}

	for shiftedY := 0; shiftedY < size; shiftedY++ {
		for shiftedX := 0; shiftedX < size; shiftedX++ {
			for y := 0; y < graph.height; y++ {
				for x := 0; x < graph.width; x++ {
					point := utils.Point{X: x + shiftedX*graph.width, Y: y + shiftedY*graph.height}
					if point.X < graph.width && point.Y < graph.height {
						continue
					}

					shiftPoint := utils.Point{X: graph.width, Y: 0}

					if shiftedX == 0 {
						shiftPoint.X = 0
						shiftPoint.Y = graph.height
					}

					previousPoint := point.Subtract(shiftPoint)

					originalValue := newCells[previousPoint]

					newValue := originalValue + 1
					if newValue > 9 {
						newValue = 1
					}
					newCells[point] = newValue
				}
			}
		}
	}

	return Graph{
		cells:  newCells,
		width:  graph.width * size,
		height: graph.height * size,
	}
}

func draw(graph Graph) {
	for y := 0; y < graph.height; y++ {
		for x := 0; x < graph.width; x++ {
			point := utils.Point{X: x, Y: y}
			fmt.Printf("%v", graph.cells[point])
		}
		fmt.Println()
	}
}

func safestPath(graph Graph, start utils.Point, end utils.Point) int {
	queue := make(map[utils.Point]struct{})
	visited := make(map[utils.Point]struct{})

	distances := make(map[utils.Point]int)
	previous := make(map[utils.Point]*utils.Point)

	previous[start] = nil
	queue[start] = struct{}{}
	distances[start] = 0

	directions := []utils.Point{
		{0, -1},
		{-1, 0}, {1, 0},
		{0, 1},
	}

	for {
		if len(queue) == 0 {
			break
		}

		var node utils.Point
		dst := math.MaxInt

		for point := range queue {
			if currentDst := distances[point]; currentDst <= dst {
				node = point
				dst = currentDst
			}
		}

		delete(queue, node)

		visited[node] = struct{}{}

		if node == end {
			break
		}

		for _, direction := range directions {
			neighbor := node.Add(direction)

			if neighbor.X < 0 || neighbor.X >= graph.width || neighbor.Y < 0 || neighbor.Y >= graph.height {
				continue
			}

			if _, isInQueue := queue[neighbor]; isInQueue {
				continue
			}

			if _, hasVisited := visited[neighbor]; hasVisited {
				continue
			}

			queue[neighbor] = struct{}{}

			alt := distances[node] + graph.cells[neighbor]
			distances[neighbor] = alt
			previous[neighbor] = &node
		}
	}

	cost := 0
	node := &end

	for {
		if node == nil {
			break
		}

		cost += graph.cells[*node]
		node = previous[*node]
	}

	cost -= graph.cells[start]

	return cost
}
