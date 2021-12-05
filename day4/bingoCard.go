package main

import "AdventOfCode2021/utils"

type cell struct {
	value   int
	stamped bool
}

type bingoCard struct {
	cells         []cell
	width, height int
}

func newBingoCard(numbers []int) bingoCard {
	cells := make([]cell, len(numbers))

	for i, number := range numbers {
		cells[i].value = number
	}

	card := bingoCard{height: 5, width: 5, cells: cells}

	return card
}

func (c *bingoCard) stampNumber(number int) {
	for i, c2 := range c.cells {
		if c2.value == number {
			c.cells[i].stamped = true
		}
	}
}

func (c *bingoCard) hasBingo() bool {
	if c.hasBingoHorizontally() {
		return true
	}

	if c.hasBingoVertically() {
		return true
	}

	return false
}

func (c *bingoCard) calculateScore(stamp int) int {
	sum := 0
	for _, c2 := range c.cells {
		if !c2.stamped {
			sum += c2.value
		}
	}

	return stamp * sum
}

func (c *bingoCard) hasBingoVertically() bool {
	span := utils.MakeRange(0, c.width-1)

	for _, i := range span {
		if c.hasBingoInColumn(i) {
			return true
		}
	}
	return false
}

func (c *bingoCard) hasBingoHorizontally() bool {
	span := utils.MakeRange(0, c.width-1)

	for _, i := range span {
		if c.hasBingoInRow(i) {
			return true
		}
	}
	return false
}

func (c *bingoCard) hasBingoInRow(i int) bool {
	cells := c.cells

	return cells[i*c.width].stamped &&
		cells[i*c.width+1].stamped &&
		cells[i*c.width+2].stamped &&
		cells[i*c.width+3].stamped &&
		cells[i*c.width+4].stamped
}

func (c *bingoCard) hasBingoInColumn(i int) bool {
	cells := c.cells

	return cells[0+i].stamped &&
		cells[5+i].stamped &&
		cells[10+i].stamped &&
		cells[15+i].stamped &&
		cells[20+i].stamped
}
