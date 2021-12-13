package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"log"
	"regexp"
	"strconv"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

type Direction int

const (
	Horizontal Direction = iota
	Vertical
)

type command struct {
	direction Direction
	point     int
}

var regex = regexp.MustCompile(`fold along (?P<Direction>\w)=(?P<Point>\d*)`)

func newCommand(s string) command {
	parts := utils.MakeMapFromRegex(regex, s)

	var direction Direction

	if parts["Direction"] == "y" {
		direction = Horizontal
	} else {
		direction = Vertical
	}

	point, _ := strconv.Atoi(parts["Point"])

	return command{
		direction: direction,
		point:     point,
	}
}

func main() {
	var paper Paper
	var points []utils.Point
	var commands []command
	if *useTestInputs {
		lines := []string{
			"6,10",
			"0,14",
			"9,10",
			"0,3",
			"10,4",
			"4,11",
			"6,0",
			"6,12",
			"4,1",
			"0,13",
			"10,12",
			"3,4",
			"3,0",
			"8,4",
			"1,10",
			"2,14",
			"8,10",
			"9,0",
		}

		for _, line := range lines {
			points = append(points, utils.NewPoint(line))
		}

		commands = []command{
			newCommand("fold along y=7"),
			newCommand("fold along x=5"),
		}
	} else {
		var err error
		points, commands, err = newOrigamiReader("./inputs/day13.txt")

		if err != nil {
			log.Fatal(err)
		}
	}

	paper = newPaper(points)

	utils.PrintDayResultsWithDuration(13, part1(paper, commands[0]), part2(paper, commands))
}

func part1(paper Paper, command command) utils.ResultWithTime {
	newPaper := paper.RunCommand(command)

	return utils.ResultWithTime{
		Value: len(newPaper.dots),
	}
}

func part2(paper Paper, commad []command) utils.ResultWithTime {
	newPaper := paper

	for _, command := range commad {
		newPaper = newPaper.RunCommand(command)
	}

	newPaper.Draw()

	return utils.ResultWithTime{}
}
