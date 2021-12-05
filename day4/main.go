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
	var cards []bingoCard
	var stamps []int

	if *useTestInputs {
		card1 := newBingoCard([]int{
			22, 13, 17, 11, 0,
			8, 2, 23, 4, 24,
			21, 9, 14, 16, 7,
			6, 10, 3, 18, 5,
			1, 12, 20, 15, 19,
		})

		card2 := newBingoCard([]int{
			3, 15, 0, 2, 22,
			9, 18, 13, 17, 5,
			19, 8, 7, 25, 23,
			20, 11, 10, 24, 4,
			14, 21, 16, 12, 6,
		})
		card3 := newBingoCard([]int{
			14, 21, 17, 24, 4,
			10, 16, 15, 9, 19,
			18, 8, 23, 26, 20,
			22, 11, 13, 6, 5,
			2, 0, 12, 3, 7,
		})

		cards = []bingoCard{card1, card2, card3}
		stamps = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	} else {
		stamps, cards, _ = newBingoReader("./inputs/day4.txt")
	}

	utils.PrintDayResults(4, part1(stamps, cards), part2(stamps, cards))
}

func part1(stamps []int, cards []bingoCard) int {
	winningIndex := 0
	hasWinner := false
	var lastStamp int

	for _, stamp := range stamps {
		lastStamp = stamp
		for i, card := range cards {
			card.stampNumber(stamp)

			if card.hasBingo() {
				winningIndex = i
				hasWinner = true
				break
			}
		}

		if hasWinner {
			break
		}
	}

	return cards[winningIndex].calculateScore(lastStamp)
}

func part2(stamps []int, cards []bingoCard) int {
	winningIndex := 0
	hasWinner := false
	lastStamp := 0

	for {
		hasWinner = false
		for _, stamp := range stamps {
			lastStamp = stamp
			for i, card := range cards {
				card.stampNumber(stamp)

				if card.hasBingo() {
					winningIndex = i
					hasWinner = true
					break
				}
			}

			if hasWinner {
				break
			}
		}

		if len(cards) == 1 {
			break
		} else {
			var filtered []bingoCard

			for i, card := range cards {
				if i != winningIndex {
					filtered = append(filtered, card)
				}
			}
			cards = filtered
		}
	}

	return cards[winningIndex].calculateScore(lastStamp)
}
