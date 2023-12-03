package aoc2023

import (
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
	res := make([][]rune, len(lines))

	for i, line := range lines {
		res[i] = []rune(line)
	}

	return res
}
