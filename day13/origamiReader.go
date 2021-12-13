package main

import (
	"AdventOfCode2021/utils"
	"bufio"
	"os"
	"strings"
)

func newOrigamiReader(file string) (points []utils.Point, commands []command, err error) {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for {
		if !scanner.Scan() {
			return
		}

		if scanner.Text() == "" {
			continue
		}

		if strings.Contains(scanner.Text(), ",") {
			points = append(points, utils.NewPoint(scanner.Text()))
		} else {
			commands = append(commands, newCommand(scanner.Text()))
		}
	}

}
