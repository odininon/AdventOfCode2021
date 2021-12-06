---
to: day<%=day%>/main.go
---
package main

import (
	"AdventOfCode2021/utils"
	"flag"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	if *useTestInputs {

	} else {

	}

	utils.PrintDayResults(<%=day%>, part1(), part2())
}

func part1() int {
	return 0
}

func part2() int {
	return 0
}
