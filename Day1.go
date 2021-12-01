package main

import (
	"AdventOfCode2021/Utils"
	"flag"
	"strconv"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	var depths []int

	if *useTestInputs {
		depths = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	} else {
		inputs, _ := Utils.ReadInputFile(1)
		depths, _ = convertToDepths(inputs)
	}

	Utils.PrintDayResults(1, part1(depths), part2(depths))
}

func part1(depths []int) int {
	count := 0

	for i := 0; i < len(depths)-1; i++ {
		if depths[i+1] > depths[i] {
			count += 1
		}
	}

	return count
}

func part2(depths []int) int {
	count := 0

	for i := 0; i < len(depths)-3; i++ {
		sum1 := depths[i] + depths[i+1] + depths[i+2]
		sum2 := depths[i+1] + depths[i+2] + depths[i+3]

		if sum2 > sum1 {
			count += 1
		}
	}

	return count
}

func convertToDepths(inputs []string) ([]int, error) {
	var depths []int

	for _, input := range inputs {
		depth, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}

		depths = append(depths, depth)
	}

	return depths, nil
}
