package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"fmt"
	"log"
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
		inputs, err := utils.ReadInputFile(1)

		if err != nil {
			log.Fatal(err)
		}

		depths, err = convertToDepths(inputs)

		if err != nil {
			log.Fatal(err)
		}
	}

	utils.PrintDayResults(1, part1(depths), part2(depths))
}

func part1(depths []int) int {
	return calculateDepthWithSpan(depths, 1)
}

func part2(depths []int) int {
	return calculateDepthWithSpan(depths, 3)
}

func calculateDepthWithSpan(depths []int, spanCount int) int {
	count := 0
	span := utils.MakeRange(0, spanCount-1)

	for i := 0; i < len(depths)-spanCount; i++ {
		sum1 := 0
		sum2 := 0

		for k := range span {
			sum1 += depths[i+k]
		}

		for k := range span {
			sum2 += depths[i+k+1]
		}

		if sum2 > sum1 {
			count += 1
		}
	}

	return count
}

func convertToDepths(inputs []string) ([]int, error) {
	var depths []int

	for _, input := range inputs {
		if input == "" {
			continue
		}
		depth, err := strconv.Atoi(input)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to depths: %w", err)
		}

		depths = append(depths, depth)
	}

	return depths, nil
}
