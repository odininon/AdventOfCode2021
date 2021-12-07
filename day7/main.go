package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	var crabs []int
	if *useTestInputs {
		crabs = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	} else {
		line, _ := utils.ReadInputFile(7)

		numbers := strings.Split(line[0], ",")

		for _, number := range numbers {
			crabNumber, _ := strconv.Atoi(number)
			crabs = append(crabs, crabNumber)
		}
	}

	utils.PrintDayResultsWithDuration(7, part1(crabs), part2(crabs))
}

func part1(crabs []int) utils.ResultWithTime {
	t1 := time.Now()
	burnTable := make(map[int]int)
	val := minFuelBurn(crabs, 0, constantFuelBurn, burnTable)
	t2 := time.Now()

	diff := t2.Sub(t1)
	return utils.ResultWithTime{
		Value:    val,
		Duration: diff,
	}
}

func part2(crabs []int) utils.ResultWithTime {
	t1 := time.Now()
	summationTable := make(map[int]int)
	burnTable := make(map[int]int)
	val := minFuelBurn(crabs, 0, exponentialFuelBurn(summationTable), burnTable)
	t2 := time.Now()

	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    val,
		Duration: diff,
	}
}

func minFuelBurn(crabs []int, target int, burn func(target int, crab int) int, burnTable map[int]int) int {
	fuelBurn := calculateFuelBurnForPosition(crabs, target, burn, burnTable)

	fuelBurnPrev := calculateFuelBurnForPosition(crabs, target-1, burn, burnTable)
	fuelBurnNext := calculateFuelBurnForPosition(crabs, target+1, burn, burnTable)

	if fuelBurn < fuelBurnPrev && fuelBurn < fuelBurnNext {
		return fuelBurn
	}

	if fuelBurnPrev < fuelBurn && fuelBurnPrev < fuelBurnNext {
		return minFuelBurn(crabs, target-1, burn, burnTable)
	}

	if fuelBurnNext < fuelBurn && fuelBurnNext < fuelBurnPrev {
		return minFuelBurn(crabs, target+1, burn, burnTable)
	}

	return -1
}

func constantFuelBurn(target, crab int) int {
	return int(math.Abs(float64(crab - target)))
}

func exponentialFuelBurn(table map[int]int) func(int, int) int {
	return func(target, crab int) int {
		fuelBurn := constantFuelBurn(target, crab)
		return calculateSummedFuelBurn(fuelBurn, table)
	}
}

func calculateSummedFuelBurn(burn int, table map[int]int) int {
	if val, ok := table[burn]; ok {
		return val
	}

	summed := 0
	if burn != 0 {
		summed = burn + calculateSummedFuelBurn(burn-1, table)
	}

	table[burn] = summed
	return summed
}

func calculateFuelBurnForPosition(crabs []int, target int, burn func(target int, crab int) int, table map[int]int) int {
	totalFuel := 0

	if val, ok := table[target]; ok {
		return val
	}

	for _, crab := range crabs {
		totalFuel += burn(target, crab)
	}

	table[target] = totalFuel

	return totalFuel
}
