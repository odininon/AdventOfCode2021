package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"fmt"
	"log"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	var inputs []string

	if *useTestInputs {
		inputs = []string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}
	} else {
		var err error

		inputs, err = utils.ReadInputFile(2)
		if err != nil {
			log.Fatal(err)
		}
	}

	commands, err := convertToCommands(inputs)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrintDayResults(2, part1(commands), part2(commands))
}

func part1(commands []command) int {
	position := newPosition(0, 0, 0)

	for _, c := range commands {
		position = position.runCommandPart1(c)
	}

	return position.horizontal * position.vertical
}

func part2(commands []command) int {
	position := newPosition(0, 0, 0)

	for _, c := range commands {
		position = position.runCommandPart2(c)
	}

	return position.horizontal * position.vertical
}

func convertToCommands(ss []string) ([]command, error) {
	var commands []command

	for _, s := range ss {
		if s == "" {
			continue
		}
		parsedCommand, err := newCommand(s)

		if err != nil {
			return nil, fmt.Errorf("unable to parse commands: %w", err)
		}
		commands = append(commands, parsedCommand)
	}

	return commands, nil
}
