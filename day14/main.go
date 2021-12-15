package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"math"
	"time"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	var input string
	instructions := make(map[string]string)

	if *useTestInputs {
		input = "NNCB"

		instructions["CH"] = "B"
		instructions["HH"] = "N"
		instructions["CB"] = "H"
		instructions["NH"] = "C"
		instructions["HB"] = "C"
		instructions["HC"] = "B"
		instructions["HN"] = "C"
		instructions["NN"] = "C"
		instructions["BH"] = "H"
		instructions["NC"] = "B"
		instructions["NB"] = "B"
		instructions["BN"] = "B"
		instructions["BB"] = "N"
		instructions["BC"] = "B"
		instructions["CC"] = "N"
		instructions["CN"] = "C"
	} else {
		input = "OFSVVSFOCBNONHKFHNPK"

		instructions["HN"] = "C"
		instructions["VB"] = "K"
		instructions["PF"] = "C"
		instructions["BO"] = "F"
		instructions["PB"] = "F"
		instructions["OH"] = "H"
		instructions["OB"] = "N"
		instructions["PN"] = "O"
		instructions["KO"] = "V"
		instructions["CK"] = "V"
		instructions["FP"] = "H"
		instructions["PC"] = "V"
		instructions["PP"] = "N"
		instructions["FN"] = "N"
		instructions["CC"] = "F"
		instructions["FC"] = "N"
		instructions["BP"] = "N"
		instructions["SH"] = "F"
		instructions["NS"] = "V"
		instructions["KK"] = "B"
		instructions["HS"] = "C"
		instructions["NV"] = "N"
		instructions["FO"] = "B"
		instructions["VO"] = "S"
		instructions["KN"] = "F"
		instructions["SC"] = "V"
		instructions["NB"] = "H"
		instructions["CH"] = "B"
		instructions["SF"] = "V"
		instructions["NP"] = "V"
		instructions["FB"] = "P"
		instructions["CV"] = "B"
		instructions["PO"] = "P"
		instructions["SV"] = "P"
		instructions["OO"] = "V"
		instructions["PS"] = "C"
		instructions["CO"] = "N"
		instructions["SP"] = "B"
		instructions["KP"] = "H"
		instructions["KH"] = "S"
		instructions["KS"] = "S"
		instructions["NH"] = "K"
		instructions["SS"] = "P"
		instructions["PV"] = "P"
		instructions["KV"] = "V"
		instructions["ON"] = "N"
		instructions["BS"] = "C"
		instructions["HP"] = "K"
		instructions["SB"] = "P"
		instructions["VC"] = "B"
		instructions["HB"] = "N"
		instructions["FS"] = "V"
		instructions["VP"] = "K"
		instructions["BB"] = "N"
		instructions["FK"] = "S"
		instructions["CS"] = "P"
		instructions["SO"] = "F"
		instructions["HF"] = "F"
		instructions["VV"] = "C"
		instructions["BC"] = "S"
		instructions["SN"] = "K"
		instructions["KB"] = "H"
		instructions["BN"] = "H"
		instructions["HO"] = "S"
		instructions["KC"] = "F"
		instructions["CP"] = "S"
		instructions["HC"] = "S"
		instructions["OS"] = "K"
		instructions["NK"] = "N"
		instructions["BF"] = "S"
		instructions["VN"] = "B"
		instructions["SK"] = "K"
		instructions["HV"] = "B"
		instructions["KF"] = "H"
		instructions["FV"] = "B"
		instructions["VF"] = "H"
		instructions["BH"] = "S"
		instructions["NN"] = "O"
		instructions["HH"] = "K"
		instructions["CN"] = "H"
		instructions["PH"] = "V"
		instructions["NF"] = "S"
		instructions["OV"] = "P"
		instructions["OC"] = "V"
		instructions["OK"] = "H"
		instructions["OF"] = "H"
		instructions["HK"] = "N"
		instructions["FH"] = "P"
		instructions["BK"] = "N"
		instructions["VS"] = "H"
		instructions["NO"] = "V"
		instructions["VK"] = "K"
		instructions["CF"] = "N"
		instructions["CB"] = "N"
		instructions["NC"] = "K"
		instructions["PK"] = "B"
		instructions["VH"] = "F"
		instructions["FF"] = "C"
		instructions["BV"] = "P"
		instructions["OP"] = "K"
	}

	utils.PrintDayResultsWithDuration(14, part1(input, instructions), part2(input, instructions))
}

func part1(polymer string, rules map[string]string) utils.ResultWithTime {
	t1 := time.Now()

	count := synthesisPolymer(polymer, rules, 10)

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    count,
		Duration: diff,
	}
}

func part2(polymer string, rules map[string]string) utils.ResultWithTime {
	t1 := time.Now()

	count := synthesisPolymer(polymer, rules, 40)

	t2 := time.Now()
	diff := t2.Sub(t1)

	return utils.ResultWithTime{
		Value:    count,
		Duration: diff,
	}
}

func synthesisPolymer(polymer string, rules map[string]string, steps int) int {
	pairs := make(map[string]int)
	letterCount := make(map[string]int)

	for i := 0; i < len(polymer)-1; i++ {
		char1 := string(polymer[i])
		char2 := string(polymer[i+1])
		chain := char1 + char2
		pairs[chain] += 1
	}

	for i := 0; i < len(polymer); i++ {
		letterCount[string(polymer[i])] += 1
	}

	for i := 0; i < steps; i++ {
		pairs, letterCount = createNewPolymer(pairs, rules, letterCount)
	}

	minCount := math.MaxInt

	for _, count := range letterCount {
		if count < minCount {
			minCount = count
		}
	}

	maxCount := math.MinInt

	for _, count := range letterCount {
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount - minCount

}

func createNewPolymer(pairs map[string]int, rules map[string]string, letterCount map[string]int) (newPairs map[string]int, newLetterCount map[string]int) {
	newPairs = make(map[string]int)
	newLetterCount = make(map[string]int)

	for v, c := range letterCount {
		newLetterCount[v] = c
	}

	for pair, count := range pairs {
		char1 := string(pair[0])
		char2 := string(pair[1])

		if value, ok := rules[pair]; ok {
			newPairs[char1+value] += count
			newPairs[value+char2] += count
			newLetterCount[value] += count
		} else {
			newPairs[char1+char2] += count
		}
	}

	return
}
