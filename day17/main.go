package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"math"
	"time"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

type Area struct {
	topLeft     utils.Point
	bottomRight utils.Point
}

func main() {
	startingPoint := utils.Point{X: 0, Y: 0}
	var targetArea Area
	if *useTestInputs {
		targetArea = Area{topLeft: utils.Point{X: 20, Y: -5}, bottomRight: utils.Point{X: 30, Y: -10}}
	} else {
		targetArea = Area{topLeft: utils.Point{X: 135, Y: -78}, bottomRight: utils.Point{X: 155, Y: -102}}
	}

	t1 := time.Now()

	maxHeights := make(map[utils.Point]int)

	for x := 0; x <= targetArea.bottomRight.X; x++ {
		for y := targetArea.bottomRight.Y; y <= 500; y++ {
			velocity := utils.Point{X: x, Y: y}
			inTrench, maxHeight := tryFiring(startingPoint, targetArea, velocity)

			if inTrench {
				maxHeights[velocity] = maxHeight
			}
		}
	}

	utils.PrintDayResultsWithDuration(17, part1(maxHeights, t1), part2(maxHeights, t1))
}

func part1(maxHeights map[utils.Point]int, t1 time.Time) utils.ResultWithTime {
	maxHeightReached := math.MinInt

	for _, height := range maxHeights {
		if height > maxHeightReached {
			maxHeightReached = height
		}
	}

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    maxHeightReached,
		Duration: diff,
	}
}

func part2(maxHeights map[utils.Point]int, t1 time.Time) utils.ResultWithTime {
	value := len(maxHeights)

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    value,
		Duration: diff,
	}
}

func withinArea(point utils.Point, target Area) bool {
	return point.X >= target.topLeft.X && point.X <= target.bottomRight.X && point.Y <= target.topLeft.Y && point.Y >= target.bottomRight.Y
}

func tryFiring(startingPoint utils.Point, target Area, velocity utils.Point) (bool, int) {
	position := startingPoint

	var trajectory []utils.Point
	inTrench := false

	for {
		if position.Y < target.bottomRight.Y || position.X > target.bottomRight.X {
			break
		}

		if withinArea(position, target) {
			inTrench = true
			break
		}

		nextPosition, nextVelocity := step(position, velocity)
		trajectory = append(trajectory, nextPosition)

		position = nextPosition
		velocity = nextVelocity
	}

	return inTrench, maxHeight(trajectory)
}

func maxHeight(trajectory []utils.Point) int {
	topHeight := math.MinInt

	for _, point := range trajectory {
		if point.Y > topHeight {
			topHeight = point.Y
		}
	}

	return topHeight
}

func step(position, velocity utils.Point) (newPosition utils.Point, newVelocity utils.Point) {
	newPosition = position.Add(velocity)

	newVelocity = velocity

	if newVelocity.X > 0 {
		newVelocity.X -= 1
	} else if newVelocity.X < 0 {
		newVelocity.X += 1
	}

	newVelocity.Y -= 1

	return
}
