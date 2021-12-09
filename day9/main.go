package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"sort"
	"time"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	var lines []string
	if *useTestInputs {
		lines = []string{
			"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678",
		}
	} else {
		lines, _ = utils.ReadInputFile(9)
	}

	heightMap := newHeightMap(lines)

	utils.PrintDayResultsWithDuration(9, part1(heightMap), part2(heightMap))
}

func part1(heightMap heightMap) utils.ResultWithTime {
	t1 := time.Now()
	lowPoints := findLowPoints(heightMap)

	riskLevel := 0
	for _, point := range lowPoints {
		value := heightMap.cells[point.Y][point.X]

		riskLevel += value + 1
	}

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    riskLevel,
		Duration: diff,
	}
}

func part2(heightMap heightMap) utils.ResultWithTime {
	t1 := time.Now()
	lowPoints := findLowPoints(heightMap)

	var sizes []int
	for _, lowPoint := range lowPoints {
		basinSize := heightMap.calculateBasinSize(lowPoint)

		sizes = append(sizes, basinSize)
	}

	sort.Ints(sizes)
	topTree := sizes[len(sizes)-3:]

	mul := 1

	for _, i := range topTree {
		mul *= i
	}

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    mul,
		Duration: diff,
	}
}

func findLowPoints(heightMap heightMap) (points []utils.Point) {
	lowPoints := make(map[utils.Point]struct{})

	for y := 0; y < heightMap.height; y++ {
		for x := 0; x < heightMap.width; x++ {
			point := heightMap.findLowestPointFromPoint(utils.Point{X: x, Y: y})
			if _, ok := lowPoints[point]; !ok {
				lowPoints[point] = struct{}{}
			}
		}
	}

	for point, _ := range lowPoints {
		points = append(points, point)
	}

	return
}
