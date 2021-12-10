package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"log"
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
		if character, isCorrupt, _ := getCorruption(line); isCorrupt {
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

func getCorruption(line string) (string, bool, []string) {
	c := string(line[0])
	rest := line[1:]

	var t []string

	for _, v := range rest {
		t = append(t, string(v))
	}

	var char string
	var isCorrupt bool
	var newRest []string

	var completions []string

	for {
		char, isCorrupt, newRest, completions = findCorruption(c, t, completions)
		if !isCorrupt && len(newRest) > 0 {
			c = newRest[0]
			t = newRest[1:]
		} else {
			break
		}
	}

	return char, isCorrupt, completions
}

func findCorruption(c string, rest []string, completions []string) (string, bool, []string, []string) {
	pairs := make(map[string]string)
	pairs["("] = ")"
	pairs["["] = "]"
	pairs["{"] = "}"
	pairs["<"] = ">"

	if len(rest) == 0 {
		return "", false, rest, append(completions, pairs[c])
	}

	if _, isOpenning := pairs[rest[0]]; isOpenning {
		if ch, isCorrupt, newRest, complets := findCorruption(rest[0], rest[1:], completions); isCorrupt {
			return ch, isCorrupt, newRest, complets
		} else {
			return findCorruption(c, newRest, complets)
		}
	} else {
		if pairs[c] != rest[0] {
			return rest[0], true, rest, completions
		} else {
			return "", false, rest[1:], completions
		}
	}
}

func part2(lines []string) utils.ResultWithTime {
	t1 := time.Now()

	points := make(map[string]int)
	points[")"] = 1
	points["]"] = 2
	points["}"] = 3
	points[">"] = 4

	var scores []int

	for _, line := range lines {
		if _, isCorrupt, completions := getCorruption(line); !isCorrupt {
			score := 0
			for _, com := range completions {
				score *= 5
				score += points[com]
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    scores[(len(scores)-1)/2],
		Duration: diff,
	}
}
