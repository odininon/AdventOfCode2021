package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"fmt"
)

var (
	useTestInputs = flag.Bool("test", true, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	var lines []string
	if *useTestInputs {
		lines = []string{
			"[({(<(())[]>[[{[]{<()<>>",
		}
	} else {

	}

	utils.PrintDayResultsWithDuration(10, part1(lines), part2(lines))
}

func part1(lines []string) utils.ResultWithTime {
	line := lines[0]

	points := make(map[rune]int)
	points[')'] = 3
	points[']'] = 57
	points['}'] = 1197
	points['>'] = 25137

	syntaxErrorScore := 0
	if character, isCorrupt := getCorruption(line); isCorrupt {
		syntaxErrorScore += points[character]
	}

	fmt.Println(syntaxErrorScore)

	return utils.ResultWithTime{}
}

func getCorruption(line string) (rune, bool) {
	c := rune(line[0])
	rest := line[1:]

	fmt.Println(rest)
	return c, false
}

func part2([]string) utils.ResultWithTime {
	return utils.ResultWithTime{}
}
