package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func splitLinesByNewLine(lines string) (parts []string) {
	pieces := strings.Split(strings.ReplaceAll(lines, "\r\n", "\n"), "\n")

	for _, piece := range pieces {
		if piece != "" {
			parts = append(parts, piece)
		}
	}

	return
}

func ReadInputFile(day int) ([]string, error) {
	data, err := os.ReadFile(fmt.Sprintf("./inputs/day%d.txt", day))

	if err != nil {
		return nil, fmt.Errorf("unable to open input file: %w", err)
	}

	return splitLinesByNewLine(string(data)), err
}

func PrintDayResults(day int, part1 int, part2 int) {
	fmt.Printf("==Day%v==\nPart1: %v, Part2: %v\n", day, part1, part2)
}

type ResultWithTime struct {
	Value    int
	Duration time.Duration
}

func PrintDayResultsWithDuration(day int, part1 ResultWithTime, part2 ResultWithTime) {
	fmt.Printf("==Day%v==\nPart1: %v, Part2: %v\n", day, part1, part2)
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func MakeMapFromRegex(regex *regexp.Regexp, s string) map[string]string {
	match := regex.FindStringSubmatch(s)

	paramsMap := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	return paramsMap
}
