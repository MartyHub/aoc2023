package aoc2023

import (
	"fmt"
	"os"
	"strings"
)

func MustRead(filePath string) string {
	return string(Must(os.ReadFile(filePath)))
}

func MustReadLines(filePath string) []string {
	return strings.Split(MustRead(filePath), "\n")
}

func MustReadRunes(filePath string) [][]rune {
	lines := MustReadLines(filePath)
	res := make([][]rune, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		res = append(res, []rune(line))
	}

	return res
}

func MustReadInts(filePath string) [][]int {
	lines := MustReadLines(filePath)
	res := make([][]int, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		ints := make([]int, len(line))

		for i, r := range line {
			ints[i] = int(r - '0')
		}

		res = append(res, ints)
	}

	return res
}

func MustSscanf(s, format string, a ...any) {
	if _, err := fmt.Sscanf(s, format, a...); err != nil {
		panic(err)
	}
}
