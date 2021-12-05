package main

import (
	"AdventOfCode2021/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func newBingoReader(file string) (stamps []int, cards []bingoCard, err error) {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	t := strings.Split(scanner.Text(), ",")

	for _, s := range t {
		v, err := strconv.Atoi(s)
		if err == nil {
			stamps = append(stamps, v)
		}
	}

	for {
		var cells []int

		for range utils.MakeRange(0, 24) {
			if !scanner.Scan() {
				return
			}
			number, _ := strconv.Atoi(scanner.Text())
			cells = append(cells, number)
		}

		cards = append(cards, newBingoCard(cells))
	}
}
