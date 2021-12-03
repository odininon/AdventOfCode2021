package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"fmt"
	"strconv"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	var readings []string

	if *useTestInputs {
		readings = []string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}
	} else {
		readings, _ = utils.ReadInputFile(3)
	}

	utils.PrintDayResults(3, part1(readings), part2(readings))
}

func part1(readings []string) int {
	return calculateEpsilon(readings) * calculateGamma(readings)
}

func part2(readings []string) int {
	return calculateOxygen(readings) * calculateC02(readings)
}

func filterByBit(r []string, i int, mcb bool) string {
	if len(r) == 1 {
		return r[0]
	}

	filterValue := sigBitAtLocation(r, i, mcb)

	var filtered []string

	for _, s := range r {
		if s == "" {
			continue
		}

		bit := fmt.Sprintf("%c", s[i])

		if bit == filterValue {
			filtered = append(filtered, s)
		}
	}

	return filterByBit(filtered, i+1, mcb)
}

func sigBitAtLocation(r []string, i int, mcb bool) string {
	countOf1s := 0
	countOf0s := 0

	for _, s := range r {
		if s == "" {
			continue
		}
		bit := fmt.Sprintf("%c", s[i])

		if bit == "1" {
			countOf1s += 1
		} else {
			countOf0s += 1
		}
	}

	var filterValue string

	if countOf0s == countOf1s {
		if mcb {
			filterValue = "1"
		} else {
			filterValue = "0"
		}
	} else if countOf0s > countOf1s {
		if mcb {
			filterValue = "0"
		} else {
			filterValue = "1"
		}
	} else {
		if mcb {
			filterValue = "1"
		} else {
			filterValue = "0"
		}
	}
	return filterValue
}

func calculateOxygen(r []string) int {
	return int(convertToDecimal(filterByBit(r, 0, true)))
}

func calculateC02(r []string) int {
	return int(convertToDecimal(filterByBit(r, 0, false)))
}

func combineSigBits(r []string, mcb bool) int {
	var combined string

	length := len(r[0])

	for i := 0; i < length; i++ {
		value := sigBitAtLocation(r, i, mcb)
		combined += value
	}

	return int(convertToDecimal(combined))
}

func calculateEpsilon(r []string) int {
	return combineSigBits(r, true)
}

func calculateGamma(r []string) int {
	return combineSigBits(r, false)
}

func convertToDecimal(b string) int64 {
	output, _ := strconv.ParseInt(b, 2, 64)
	return output
}
