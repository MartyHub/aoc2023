package main

import (
	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 288, parse("day06/data/sample1.txt").Ways)
	aoc2023.Expect("Part 1", 220320, parse("day06/data/input1.txt").Ways)
}

func part2() {
	aoc2023.Expect("Sample 2", 71503, parse("day06/data/sample2.txt").Ways)
	aoc2023.Expect("Part 2", 34454850, parse("day06/data/input2.txt").Ways)
}

func parse(filePath string) Races {
	lines := aoc2023.MustReadLines(filePath)
	times := aoc2023.ToInts(lines[0])
	distances := aoc2023.ToInts(lines[1])

	var res Races

	for i := range times {
		res = append(res, Race{
			Distance: distances[i],
			Time:     times[i],
		})
	}

	return res
}

type (
	Race struct {
		Distance int
		Time     int
	}

	Races []Race
)

func (race Race) Ways() int {
	res := 0

	for i := 1; i < race.Time; i++ {
		if dist := i * (race.Time - i); dist > race.Distance {
			res++
		}
	}

	return res
}

func (races Races) Ways() int {
	res := 1

	for _, race := range races {
		res *= race.Ways()
	}

	return res
}
