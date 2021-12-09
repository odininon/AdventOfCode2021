package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"regexp"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

type lineSegment struct {
	start, end utils.Point
}

func (l lineSegment) isVertical() bool {
	return l.start.X == l.end.X
}

func (l lineSegment) isHorizontal() bool {
	return l.start.Y == l.end.Y
}

func main() {
	var inputs []string

	if *useTestInputs {
		inputs = []string{
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
			"2,2 -> 2,1",
			"7,0 -> 7,4",
			"6,4 -> 2,0",
			"0,9 -> 2,9",
			"3,4 -> 1,4",
			"0,0 -> 8,8",
			"5,5 -> 8,2",
		}
	} else {
		inputs, _ = utils.ReadInputFile(5)
	}

	lineSegments := convertToLineSegments(inputs)

	utils.PrintDayResults(5, part1(lineSegments), part2(lineSegments))
}

func part1(segments []lineSegment) int {
	return getOverlaps(segments, false)
}

func part2(segments []lineSegment) int {
	return getOverlaps(segments, true)
}

func getOverlaps(segments []lineSegment, includeDiagonals bool) int {
	counts := make(map[utils.Point]int)

	for _, segment := range segments {
		points := getPointsOfLineSegment(segment, includeDiagonals)

		for _, p := range points {
			counts[p] += 1
		}
	}

	sum := 0

	for _, i := range counts {
		if i >= 2 {
			sum += 1
		}
	}

	return sum
}

func getPointsOfLineSegment(segment lineSegment, includeDiagonals bool) (points []utils.Point) {
	p1 := segment.start
	p2 := segment.end

	direction := p2.Subtract(p1).Sign()

	if segment.isVertical() || segment.isHorizontal() || includeDiagonals {
		for p := p1; p != p2; p = p.Add(direction) {
			points = append(points, p)
		}

		points = append(points, p2)
	}
	return
}

var regex = regexp.MustCompile(`(?P<Start>\d*,\d*)\s->\s(?P<End>\d*,\d*)`)

func convertToLineSegments(inputs []string) (lineSegments []lineSegment) {
	for _, line := range inputs {
		if line == "" {
			continue
		}
		params := utils.MakeMapFromRegex(regex, line)

		lineSegments = append(lineSegments, lineSegment{start: utils.NewPoint(params["Start"]), end: utils.NewPoint(params["End"])})
	}

	return
}
