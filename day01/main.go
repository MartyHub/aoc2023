package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MartyHub/aoc2023"
)

type digit struct {
	index int
	value int
}

func main() {
	part1()
	part2()
}

func part1() {
	integers := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	aoc2023.Expect("Sample 1", sum("day01/data/sample1.txt", integers), 142)
	aoc2023.Expect("Part 1", sum("day01/data/input.txt", integers), 55834)
}

func part2() {
	integers := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	aoc2023.Expect("Sample 2", sum("day01/data/sample2.txt", integers), 281)
	aoc2023.Expect("Part 2", sum("day01/data/input.txt", integers), 53221)
}

func sum(filePath string, integers []string) int {
	res := 0

	for _, line := range aoc2023.MustReadLines(filePath) {
		var first, last digit

		for i, s := range integers {
			if idx := strings.Index(line, s); idx != -1 {
				if idx < first.index || first.value == 0 {
					first.index = idx
					first.value = i%9 + 1
				}
			}

			if idx := strings.LastIndex(line, s); idx != -1 {
				if idx > last.index || last.value == 0 {
					last.index = idx
					last.value = i%9 + 1
				}
			}
		}

		res += aoc2023.Must(
			strconv.Atoi(fmt.Sprintf("%d%d",
				first.value,
				last.value,
			)),
		)
	}

	return res
}
