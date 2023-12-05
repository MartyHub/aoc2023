package main

import (
	"regexp"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 13, parse("day04/data/sample.txt").Score)
	aoc2023.Expect("Part 1", 20667, parse("day04/data/input.txt").Score)
}

func part2() {
	aoc2023.Expect("Sample 2", 30, parse("day04/data/sample.txt").Scratchcards)
	aoc2023.Expect("Part 2", 5833065, parse("day04/data/input.txt").Scratchcards)
}

func parse(filePath string) Cards {
	var res Cards

	for i, line := range aoc2023.MustReadLines(filePath) {
		line = strings.SplitN(line, ":", 2)[1]
		nums := strings.SplitN(line, "|", 2)

		res = append(res, Card{
			id:      i + 1,
			gotNums: parseNums(nums[1]),
			winNums: parseNums(nums[0]),
		})
	}

	return res
}

func parseNums(s string) map[int]struct{} {
	re := regexp.MustCompile(`(\d+)`)
	res := make(map[int]struct{})

	for _, match := range re.FindAllString(s, -1) {
		res[aoc2023.ToInt(match)] = struct{}{}
	}

	return res
}

type (
	Card struct {
		id      int
		gotNums map[int]struct{}
		winNums map[int]struct{}
	}

	Cards []Card
)

func (card Card) WinCount() int {
	res := 0

	for num := range card.gotNums {
		if _, found := card.winNums[num]; found {
			res++
		}
	}

	return res
}

func (card Card) Score() int {
	res := 0

	for num := range card.gotNums {
		if _, found := card.winNums[num]; found {
			if res == 0 {
				res = 1
			} else {
				res *= 2
			}
		}
	}

	return res
}

func (cards Cards) Score() int {
	res := 0

	for _, card := range cards {
		res += card.Score()
	}

	return res
}

func (cards Cards) Scratchcards() int {
	scratchcards := make(map[int]int)

	for _, card := range cards {
		scratchcards[card.id]++

		for i := 1; i <= card.WinCount(); i++ {
			scratchcards[card.id+i] += scratchcards[card.id]
		}
	}

	res := 0

	for _, nb := range scratchcards {
		res += nb
	}

	return res
}
