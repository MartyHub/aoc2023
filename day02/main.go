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
	bag := Set{
		red:   12,
		green: 13,
		blue:  14,
	}

	aoc2023.Expect("Sample 1", 8, func() int { return parse("day02/data/sample.txt").sum(bag) })
	aoc2023.Expect("Part 1", 2377, func() int { return parse("day02/data/input.txt").sum(bag) })
}

func part2() {
	aoc2023.Expect("Sample 2", 2286, parse("day02/data/sample.txt").min)
	aoc2023.Expect("Part 2", 71220, parse("day02/data/input.txt").min)
}

func parse(filePath string) Games {
	var res Games

	for i, line := range aoc2023.MustReadLines(filePath) {
		var sets []Set

		for _, lineSet := range strings.Split(line, ";") {
			sets = append(sets, parseSet(lineSet))
		}

		res = append(res, Game{
			id:   i + 1,
			sets: sets,
		})
	}

	return res
}

func parseSet(s string) Set {
	var res Set

	re := regexp.MustCompile(`(\d+) (red|blue|green)`)

	for _, match := range re.FindAllStringSubmatch(s, -1) {
		count := aoc2023.ToInt(match[1])

		switch match[2] {
		case "red":
			res.red += count
		case "green":
			res.green += count
		case "blue":
			res.blue += count
		default:
			panic(match[2])
		}
	}

	return res
}

type (
	Games []Game

	Game struct {
		id   int
		sets []Set
	}

	Set struct {
		red   int
		green int
		blue  int
	}
)

func (games Games) sum(bag Set) int {
	res := 0

	for _, game := range games {
		ok := true

		for _, set := range game.sets {
			if !bag.contains(set) {
				ok = false

				break
			}
		}

		if ok {
			res += game.id
		}
	}

	return res
}

func (games Games) min() int {
	res := 0

	for _, game := range games {
		var minSet Set

		for _, set := range game.sets {
			minSet.red = max(minSet.red, set.red)
			minSet.green = max(minSet.green, set.green)
			minSet.blue = max(minSet.blue, set.blue)
		}

		res += minSet.power()
	}

	return res
}

func (set Set) contains(other Set) bool {
	return other.red <= set.red &&
		other.green <= set.green &&
		other.blue <= set.blue
}

func (set Set) power() int {
	return set.red * set.green * set.blue
}
