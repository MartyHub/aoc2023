package main

import (
	"image"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 405, parse("day13/data/sample.txt").Compute1)
	aoc2023.Expect("Part 1", 30802, parse("day13/data/input.txt").Compute1)
}

func part2() {
	aoc2023.Expect("Sample 2", 400, parse("day13/data/sample.txt").Compute2)
	aoc2023.Expect("Part 2", 37876, parse("day13/data/input.txt").Compute2)
}

func parse(filePath string) Patterns {
	var (
		cur Pattern
		res Patterns
	)

	for _, line := range aoc2023.MustReadLines(filePath) {
		if line == "" {
			if cur != nil {
				res = append(res, cur)
				cur = nil
			}

			continue
		}

		cur = append(cur, []rune(line))
	}

	if cur != nil {
		res = append(res, cur)
	}

	return res
}

type (
	Pattern [][]rune

	Patterns []Pattern
)

func (ptt Pattern) String() string {
	sb := strings.Builder{}

	for i, runes := range ptt {
		if i > 0 {
			sb.WriteRune('\n')
		}

		sb.WriteString(string(runes))
	}

	return sb.String()
}

func (ptt Pattern) Compute1() int {
	if res := ptt.ComputeHorizontal(-1); res > 0 {
		return res * 100
	}

	if res := ptt.ComputeVertical(-1); res > 0 {
		return res
	}

	panic("no reflection")
}

func (ptt Pattern) Compute2() int {
	var (
		pt1    image.Point
		smudge image.Point
	)

	if res := ptt.ComputeHorizontal(-1); res > 0 {
		pt1.Y = res
	} else if res = ptt.ComputeVertical(-1); res > 0 {
		pt1.X = res
	} else {
		panic("no reflection")
	}

	for y := 0; y < ptt.Height(); y++ {
		for x := 0; x < ptt.Width(); x++ {
			if x != 0 || y != 0 {
				ptt.Toggle(smudge)
			}

			smudge = image.Point{X: x, Y: y}

			ptt.Toggle(smudge)

			if res := ptt.ComputeHorizontal(pt1.Y - 1); res > 0 {
				return res * 100
			}

			if res := ptt.ComputeVertical(pt1.X - 1); res > 0 {
				return res
			}
		}
	}

	panic("no reflection")
}

func (ptt Pattern) Toggle(pt image.Point) {
	if ptt[pt.Y][pt.X] == '.' {
		ptt[pt.Y][pt.X] = '#'
	} else {
		ptt[pt.Y][pt.X] = '.'
	}
}

func (ptt Pattern) ComputeHorizontal(exclude int) int {
	for y := 0; y < ptt.Height()-1; y++ {
		if y == exclude {
			continue
		}

		if ptt.Horizontal(y) {
			return y + 1
		}
	}

	return 0
}

func (ptt Pattern) Horizontal(y int) bool {
	for dy := 0; y-dy >= 0 && y+1+dy < ptt.Height(); dy++ {
		if string(ptt[y-dy]) != string(ptt[y+1+dy]) {
			return false
		}
	}

	return true
}

func (ptt Pattern) ComputeVertical(exclude int) int {
	for x := 0; x < ptt.Width()-1; x++ {
		if x == exclude {
			continue
		}

		if ptt.Vertical(x) {
			return x + 1
		}
	}

	return 0
}

func (ptt Pattern) Vertical(x int) bool {
	for dx := 0; x-dx >= 0 && x+1+dx < ptt.Width(); dx++ {
		for y := 0; y < ptt.Height(); y++ {
			if ptt[y][x-dx] != ptt[y][x+1+dx] {
				return false
			}
		}
	}

	return true
}

func (ptt Pattern) Height() int {
	return len(ptt)
}

func (ptt Pattern) Width() int {
	return len(ptt[0])
}

func (pts Patterns) Compute1() int {
	res := 0

	for _, ptt := range pts {
		res += ptt.Compute1()
	}

	return res
}

func (pts Patterns) Compute2() int {
	res := 0

	for _, ptt := range pts {
		res += ptt.Compute2()
	}

	return res
}
