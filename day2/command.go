package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile(`(?P<Direction>\w*)\s(?P<Increment>\d*)`)

type command struct {
	direction string
	increment int
}

func newCommand(s string) (command, error) {
	match := regex.FindStringSubmatch(s)

	paramsMap := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	increment, err := strconv.Atoi(paramsMap["Increment"])

	if err != nil {
		return command{}, fmt.Errorf("unabled to parse input: %w", err)
	}

	return command{
		direction: paramsMap["Direction"],
		increment: increment,
	}, nil
}

func (p position) runCommandPart1(c command) position {
	newPos := newPosition(p.horizontal, p.vertical, p.aim)

	switch c.direction {
	case "forward":
		newPos.horizontal += c.increment
	case "down":
		newPos.vertical += c.increment
	case "up":
		newPos.vertical -= c.increment
	}

	return newPos
}

func (p position) runCommandPart2(c command) position {
	newPos := newPosition(p.horizontal, p.vertical, p.aim)

	switch c.direction {
	case "forward":
		newPos.horizontal += c.increment
		newPos.vertical += newPos.aim * c.increment
	case "down":
		newPos.aim += c.increment
	case "up":
		newPos.aim -= c.increment
	}

	return newPos
}
