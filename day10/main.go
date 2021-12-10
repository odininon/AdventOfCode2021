package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"log"
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
			"[({(<(())[]>[[{[]{<()<>>",
			"[(()[<>])]({[<{<<[]>>(",
			"{([(<{}[<>[]}>{[]{[(<()>",
			"(((({<>}<{<{<>}{[]{[]{}",
			"[[<[([]))<([[{}[[()]]]",
			"[{[{({}]{}}([{[{{{}}([]",
			"{<[[]]>}<{[{[{[]{()[[[]",
			"[<(<(<(<{}))><([]([]()",
			"<{([([[(<>()){}]>(<<{{",
			"<{([{{}}[<[[[<>{}]]]>[]]",
		}
	} else {
		var err error
		lines, err = utils.ReadInputFile(10)

		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	utils.PrintDayResultsWithDuration(10, part1(lines), part2(lines))
}

func part1(lines []string) utils.ResultWithTime {
	t1 := time.Now()

	points := make(map[string]int)
	points[")"] = 3
	points["]"] = 57
	points["}"] = 1197
	points[">"] = 25137

	syntaxErrorScore := 0

	for _, line := range lines {
		if character, isCorrupt := getCorruption(line); isCorrupt {
			// fmt.Printf("line = %v - %v - %v\n", i, character, points[character])
			syntaxErrorScore += points[character]
		}
	}

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    syntaxErrorScore,
		Duration: diff,
	}
}

func getCorruption(line string) (string, bool) {
	c := string(line[0])
	rest := line[1:]

	var t []string

	for _, v := range rest {
		t = append(t, string(v))
	}

	var char string
	var isCorrupt bool
	var newRest []string

	for {
		char, isCorrupt, newRest = findCorruption(c, t)
		if !isCorrupt && len(newRest) > 0 {
			c = newRest[0]
			t = newRest[1:]
		} else {
			break
		}
	}

	return char, isCorrupt
}

func findCorruption(c string, rest []string) (string, bool, []string) {
	pairs := make(map[string]string)
	pairs["("] = ")"
	pairs["["] = "]"
	pairs["{"] = "}"
	pairs["<"] = ">"

	// fmt.Printf("%v - %v\n", c, rest)

	if len(rest) == 0 {
		return "", false, rest
	}

	if _, isOpenning := pairs[rest[0]]; isOpenning {
		if ch, isCorrupt, newRest := findCorruption(rest[0], rest[1:]); isCorrupt {
			return ch, isCorrupt, newRest
		} else {
			return findCorruption(c, newRest)
		}
	} else {
		if pairs[c] != rest[0] {
			return rest[0], true, rest
		} else {
			return "", false, rest[1:]
		}
	}
}

func part2([]string) utils.ResultWithTime {
	return utils.ResultWithTime{}
}
