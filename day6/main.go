package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"strconv"
	"strings"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	var fish []int

	if *useTestInputs {
		fish = []int{3, 4, 3, 1, 2}
	} else {
		line, _ := utils.ReadInputFile(6)

		numbers := strings.Split(line[0], ",")

		for _, number := range numbers {
			fishNumber, _ := strconv.Atoi(number)
			fish = append(fish, fishNumber)
		}
	}

	utils.PrintDayResults(6, part1(fish), part2(fish))
}

func part1(fish []int) int {
	return countFish(tickDays(fish, 80))
}

func part2(fish []int) int {
	return countFish(tickDays(fish, 256))
}

func tickDays(fish []int, i int) (fishMap map[int]int) {
	fishMap = make(map[int]int)

	for _, i2 := range fish {
		fishMap[i2] += 1
	}

	for range utils.MakeRange(0, i-1) {
		fishMap = tickDay(fishMap)
	}
	return
}

func tickDay(fish map[int]int) (newFish map[int]int) {
	newFish = make(map[int]int)

	for i, i2 := range fish {
		if i == 0 {
			newFish[6] += i2
			newFish[8] += i2
		} else {
			newFish[i-1] += i2
		}
	}

	return
}

func countFish(fish map[int]int) int {
	count := 0
	for _, i := range fish {
		count += i
	}
	return count
}
