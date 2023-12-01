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
