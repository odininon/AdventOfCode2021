package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"strings"
	"time"
)

var (
	useTestInputs = flag.Bool("test", true, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

type display struct {
	signals []string
	outputs []string
}

func newDisplay(s string) display {
	parts := strings.Split(s, " | ")

	return display{
		signals: strings.Split(parts[0], " "),
		outputs: strings.Split(parts[1], " "),
	}
}

func main() {
	var lines []string
	if *useTestInputs {
		lines = []string{"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
			"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
			"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
			"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
			"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
			"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
			"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
			"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
			"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
			"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"}

	} else {
		lines, _ = utils.ReadInputFile(8)
	}

	var displays []display

	for _, line := range lines {
		if line == "" {
			continue
		}
		displays = append(displays, newDisplay(line))
	}

	utils.PrintDayResultsWithDuration(8, part1(displays), part2(displays))
}

func part1(displays []display) utils.ResultWithTime {
	t1 := time.Now()
	count := countUniqueNumbersInOutput(displays)
	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    count,
		Duration: diff,
	}
}

func countUniqueNumbersInOutput(displays []display) int {
	count := 0

	for _, display := range displays {
		for _, output := range display.outputs {
			if len(output) == 2 || len(output) == 4 || len(output) == 3 || len(output) == 7 {
				count += 1
			}
		}
	}

	return count
}

func part2(displays []display) utils.ResultWithTime {
	return utils.ResultWithTime{}
}
