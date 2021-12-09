package main

import (
	"AdventOfCode2021/utils"
	"strconv"
)

type heightMap struct {
	width, height int
	cells         [][]int
}

func (m heightMap) findLowestPointFromPoint(point utils.Point) utils.Point {
	neighborPoints := []utils.Point{point}
	blackList := make(map[utils.Point][]utils.Point)
	endPoints := make(map[int]utils.Point)

	for {
		nextPoint := neighborPoints[0]
		neighborPoints = neighborPoints[1:]

		currentValue := m.cells[nextPoint.Y][nextPoint.X]
		nextPoints := m.getLowerNeighborsForPoint(nextPoint, blackList)

		if len(nextPoints) == 0 {
			endPoints[currentValue] = nextPoint

			if len(neighborPoints) == 0 {
				lowest := 10
				for i, _ := range endPoints {
					if i < lowest {
						lowest = i
					}
				}

				return endPoints[lowest]
			}
		}

		neighborPoints = append(neighborPoints, nextPoints...)
	}
}

func (m heightMap) getLowerNeighborsForPoint(point utils.Point, blackList map[utils.Point][]utils.Point) (points []utils.Point) {
	neighbors := []utils.Point{
		{X: 0, Y: -1},
		{X: -1, Y: 0}, {X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	currentValue := m.cells[point.Y][point.X]

	for _, neighbor := range neighbors {
		newPoint := neighbor.Add(point)

		if newPoint.X < 0 || newPoint.Y < 0 || newPoint.X >= m.width || newPoint.Y >= m.height {
			continue
		}

		isRestricted := false
		for _, u := range blackList[point] {
			if u.X == newPoint.X || u.Y == newPoint.Y {
				isRestricted = true
				break
			}
		}

		if isRestricted {
			continue
		}

		neighborValue := m.cells[newPoint.Y][newPoint.X]

		if currentValue > neighborValue {
			points = append(points, newPoint)
		}

		if currentValue == neighborValue {
			points = append(points, newPoint)
			blackList[point] = append(blackList[point], newPoint)
		}
	}

	return
}

func (m heightMap) getHigherNeighborsForPoint(point utils.Point) (points []utils.Point) {
	neighbors := []utils.Point{
		{X: 0, Y: -1},
		{X: -1, Y: 0}, {X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	currentValue := m.cells[point.Y][point.X]

	for _, neighbor := range neighbors {
		newPoint := neighbor.Add(point)

		if newPoint.X < 0 || newPoint.Y < 0 || newPoint.X >= m.width || newPoint.Y >= m.height {
			continue
		}

		neighborValue := m.cells[newPoint.Y][newPoint.X]

		if neighborValue == 9 {
			continue
		}

		if currentValue < neighborValue {
			points = append(points, newPoint)
		}
	}

	return
}

func (m heightMap) calculateBasinSize(original utils.Point) int {
	basinPoints := make(map[utils.Point]struct{})

	basinPoints[original] = struct{}{}
	neighbors := []utils.Point{original}

	for {
		point := neighbors[0]
		neighbors = neighbors[1:]

		newPoints := m.getHigherNeighborsForPoint(point)

		if len(neighbors) == 0 && len(newPoints) == 0 {
			return len(basinPoints)
		}

		for _, newPoint := range newPoints {
			if _, ok := basinPoints[newPoint]; !ok {
				basinPoints[newPoint] = struct{}{}
			}
		}

		neighbors = append(neighbors, newPoints...)
	}
}

func newHeightMap(lines []string) heightMap {
	height := len(lines)
	width := len(lines[0])

	cells := make([][]int, height)

	for k, _ := range cells {
		cells[k] = make([]int, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			char := string(lines[y][x])
			number, _ := strconv.Atoi(char)

			cells[y][x] = number
		}
	}

	return heightMap{height: height, width: width, cells: cells}
}
