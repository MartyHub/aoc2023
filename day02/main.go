package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/MartyHub/aoc2023"
)

type set struct {
	red   int
	green int
	blue  int
}

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", sum("day02/data/sample.txt"), 8)
	aoc2023.Expect("Part 1", sum("day02/data/input.txt"), 2377)
}

func part2() {
	aoc2023.Expect("Sample 2", min("day02/data/sample.txt"), 2286)
	aoc2023.Expect("Part 2", min("day02/data/input.txt"), 71220)
}

func sum(filePath string) int {
	bag := set{
		red:   12,
		green: 13,
		blue:  14,
	}
	res := 0

	for i, line := range aoc2023.MustReadLines(filePath) {
		lineSets := strings.SplitN(line, ":", 2)[1]
		ok := true

		for _, lineSet := range strings.Split(lineSets, ";") {
			s := newSet(lineSet)

			if !bag.contains(s) {
				ok = false

				break
			}
		}

		if ok {
			res += i + 1
		}
	}

	return res
}

func min(filePath string) int {
	res := 0

	for _, line := range aoc2023.MustReadLines(filePath) {
		lineSets := strings.SplitN(line, ":", 2)[1]
		minSet := set{}

		for _, lineSet := range strings.Split(lineSets, ";") {
			s := newSet(lineSet)

			minSet.red = max(minSet.red, s.red)
			minSet.green = max(minSet.green, s.green)
			minSet.blue = max(minSet.blue, s.blue)
		}

		res += minSet.power()
	}

	return res
}

func newSet(s string) set {
	re := regexp.MustCompile(`(\d+)`)
	res := set{}

	for _, cube := range strings.Split(s, ",") {
		count := aoc2023.Must(strconv.Atoi(re.FindString(cube)))

		switch {
		case strings.Contains(cube, "red"):
			res.red += count
		case strings.Contains(cube, "green"):
			res.green += count
		case strings.Contains(cube, "blue"):
			res.blue += count
		default:
			panic(cube)
		}
	}

	return res
}

func (st set) contains(other set) bool {
	return other.red <= st.red &&
		other.green <= st.green &&
		other.blue <= st.blue
}

func (st set) power() int {
	return st.red * st.green * st.blue
}
