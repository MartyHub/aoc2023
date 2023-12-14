package main

import (
	"slices"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 136, parse("day14/data/sample.txt").Compute1)
	aoc2023.Expect("Part 1", 110274, parse("day14/data/input.txt").Compute1)
}

func part2() {
	aoc2023.Expect("Sample 2", 64, parse("day14/data/sample.txt").Compute2)
	aoc2023.Expect("Part 2", 90982, parse("day14/data/input.txt").Compute2)
}

func parse(filePath string) Platform {
	return aoc2023.MustReadRunes(filePath)
}

type (
	Platform  [][]rune
	Platforms []Platform
)

func (pf Platform) RollNorth(x, y int) {
	for ; y > 0 && pf[y-1][x] == '.'; y-- {
		pf[y][x] = '.'
		pf[y-1][x] = 'O'
	}
}

func (pf Platform) RollWest(x, y int) {
	for ; x > 0 && pf[y][x-1] == '.'; x-- {
		pf[y][x] = '.'
		pf[y][x-1] = 'O'
	}
}

func (pf Platform) RollSouth(x, y int) {
	height := pf.Height()

	for ; y < height-1 && pf[y+1][x] == '.'; y++ {
		pf[y][x] = '.'
		pf[y+1][x] = 'O'
	}
}

func (pf Platform) RollEast(x, y int) {
	width := pf.Width()

	for ; x < width-1 && pf[y][x+1] == '.'; x++ {
		pf[y][x] = '.'
		pf[y][x+1] = 'O'
	}
}

func (pf Platform) TiltNorth() Platform {
	height := pf.Height()
	width := pf.Width()

	for y := 1; y < height; y++ {
		for x := 0; x < width; x++ {
			if pf[y][x] == 'O' {
				pf.RollNorth(x, y)
			}
		}
	}

	return pf
}

func (pf Platform) TiltWest() Platform {
	height := pf.Height()
	width := pf.Width()

	for y := 0; y < height; y++ {
		for x := 1; x < width; x++ {
			if pf[y][x] == 'O' {
				pf.RollWest(x, y)
			}
		}
	}

	return pf
}

func (pf Platform) TiltSouth() Platform {
	height := pf.Height()
	width := pf.Width()

	for y := height - 2; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if pf[y][x] == 'O' {
				pf.RollSouth(x, y)
			}
		}
	}

	return pf
}

func (pf Platform) TiltEast() Platform {
	height := pf.Height()
	width := pf.Width()

	for y := 0; y < height; y++ {
		for x := width - 2; x >= 0; x-- {
			if pf[y][x] == 'O' {
				pf.RollEast(x, y)
			}
		}
	}

	return pf
}

func (pf Platform) Cycle() {
	pf.TiltNorth()
	pf.TiltWest()
	pf.TiltSouth()
	pf.TiltEast()
}

func (pf Platform) Load() int {
	height := pf.Height()
	width := pf.Width()
	res := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if pf[y][x] == 'O' {
				res += height - y
			}
		}
	}

	return res
}

func (pf Platform) Compute1() int {
	return pf.TiltNorth().Load()
}

func (pf Platform) Compute2() int {
	const cycles = 1_000_000_000

	var (
		cycle       int
		countCycles int
		hist        Platforms
	)

	for i := 0; i <= cycles; i++ {
		hist = append(hist, pf.Clone())

		pf.Cycle()

		if match, idx := hist.Match(pf); match {
			if cycle == 0 {
				cycle = len(hist) - idx
			} else if cycle != len(hist)-idx {
				panic("cycle")
			}

			countCycles++

			if countCycles == 10 {
				return hist[len(hist)-cycle+(cycles-i-1)%cycle].Load()
			}
		}
	}

	panic("no cycle")
}

func (pf Platform) Height() int {
	return len(pf)
}

func (pf Platform) Width() int {
	return len(pf[0])
}

func (pf Platform) Equals(other Platform) bool {
	height := pf.Height()
	width := pf.Width()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if pf[y][x] != other[y][x] {
				return false
			}
		}
	}

	return true
}

func (pf Platform) Clone() Platform {
	height := pf.Height()
	res := make(Platform, height)

	for y := 0; y < height; y++ {
		res[y] = slices.Clone(pf[y])
	}

	return res
}

func (pf Platform) String() string {
	sb := strings.Builder{}
	height := pf.Height()

	for y := 0; y < height; y++ {
		if y > 0 {
			sb.WriteRune('\n')
		}

		sb.WriteString(string(pf[y]))
	}

	return sb.String()
}

func (hist Platforms) Match(pf Platform) (bool, int) {
	for i := len(hist) - 1; i >= 0; i-- {
		if pf.Equals(hist[i]) {
			return true, i
		}
	}

	return false, 0
}
