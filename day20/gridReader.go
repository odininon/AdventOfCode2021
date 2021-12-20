package main

import (
	"bufio"
	"os"
)

func newGridReader(file string) (lines []string, encoding string, err error) {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	encoding = scanner.Text()

	for {
		if !scanner.Scan() {
			return
		}

		if scanner.Text() == "" {
			continue
		}

		lines = append(lines, scanner.Text())
	}
}
