package Utils

import (
	"fmt"
	"os"
	"strings"
)

func splitLinesByNewLine(lines string) []string {
	return strings.Split(strings.ReplaceAll(lines, "\r\n", "\n"), "\n")
}

func ReadInputFile(day int) ([]string, error) {
	data, err := os.ReadFile(fmt.Sprintf("./inputs/day%d.txt", day))

	if err != nil {
		return nil, err
	}

	return splitLinesByNewLine(string(data)), err
}

func PrintDayResults(day int, part1 int, part2 int) {
	fmt.Printf("==Day%v==\nPart1: %v, Part2: %v\n", day, part1, part2)
}
