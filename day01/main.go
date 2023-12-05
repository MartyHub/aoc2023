package main

import (
	"fmt"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	integers := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	aoc2023.Expect("Sample 1", 142, parse("day01/data/sample1.txt", integers).sum)
	aoc2023.Expect("Part 1", 55834, parse("day01/data/input.txt", integers).sum)
}

func part2() {
	integers := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	aoc2023.Expect("Sample 2", 281, parse("day01/data/sample2.txt", integers).sum)
	aoc2023.Expect("Part 2", 53221, parse("day01/data/input.txt", integers).sum)
}

func parse(filePath string, integers []string) Numbers {
	var res Numbers

	for _, line := range aoc2023.MustReadLines(filePath) {
		num := new(Number)

		for i, s := range integers {
			num.Add(Digit{
				index: strings.Index(line, s),
				value: i%9 + 1,
			})
			num.Add(Digit{
				index: strings.LastIndex(line, s),
				value: i%9 + 1,
			})
		}

		res = append(res, num)
	}

	return res
}

type (
	Digit struct {
		index int
		value int
	}

	Number [2]Digit

	Numbers []*Number
)

func (num *Number) Add(digit Digit) {
	if digit.index == -1 {
		return
	}

	if digit.index < num[0].index || num[0].value == 0 {
		num[0] = digit
	}

	if digit.index > num[1].index || num[1].value == 0 {
		num[1] = digit
	}
}

func (num *Number) Value() int {
	return aoc2023.ToInt(fmt.Sprintf("%d%d",
		num[0].value,
		num[1].value,
	))
}

func (nums Numbers) sum() int {
	res := 0

	for _, num := range nums {
		res += num.Value()
	}

	return res
}
