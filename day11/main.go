package main

import (
	"fmt"
	"image"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 374, parse("day11/data/sample.txt").Compute1)
	aoc2023.Expect("Part 1", 9536038, parse("day11/data/input.txt").Compute1)
}

func part2() {
	aoc2023.Expect("Sample 2", 1030, parse("day11/data/sample.txt").Compute10)
	aoc2023.Expect("Sample 2", 8410, parse("day11/data/sample.txt").Compute100)
	aoc2023.Expect("Part 2", 447744640566, parse("day11/data/input.txt").Compute2)
}

func parse(filePath string) Image {
	return Image{data: aoc2023.MustReadRunes(filePath)}
}

type (
	Image struct {
		data [][]rune
	}

	Pair struct {
		From image.Point
		To   image.Point
	}
)

func (img Image) Compute1() int {
	return img.Compute(2)
}

func (img Image) Compute10() int {
	return img.Compute(10)
}

func (img Image) Compute100() int {
	return img.Compute(100)
}

func (img Image) Compute2() int {
	return img.Compute(1_000_000)
}

func (img Image) Compute(expansion int) int {
	res := 0
	spaceCols := img.SpaceCols()
	spaceRows := img.SpaceRows()

	for _, pair := range img.Pairs() {
		res += pair.Compute(spaceCols, spaceRows, expansion)
	}

	return res
}

func (img Image) SpaceCols() map[int]bool {
	res := make(map[int]bool)

	for x := 0; x < img.Width(); x++ {
		if img.IsColSpace(x) {
			res[x] = true
		}
	}

	return res
}

func (img Image) SpaceRows() map[int]bool {
	res := make(map[int]bool)

	for y := 0; y < img.Height(); y++ {
		if img.IsRowSpace(y) {
			res[y] = true
		}
	}

	return res
}

func (img Image) IsColSpace(x int) bool {
	for _, row := range img.data {
		if row[x] == '#' {
			return false
		}
	}

	return true
}

func (img Image) IsRowSpace(y int) bool {
	for _, r := range img.data[y] {
		if r == '#' {
			return false
		}
	}

	return true
}

func (img Image) Pairs() []Pair {
	var res []Pair

	gal := img.Galaxies()

	for i := 0; i < len(gal)-1; i++ {
		for j := i + 1; j < len(gal); j++ {
			res = append(res, Pair{
				From: gal[i],
				To:   gal[j],
			})
		}
	}

	return res
}

func (img Image) Galaxies() []image.Point {
	var res []image.Point

	for y, row := range img.data {
		for x, r := range row {
			if r == '#' {
				res = append(res, image.Point{
					X: x,
					Y: y,
				})
			}
		}
	}

	return res
}

func (img Image) Height() int {
	return len(img.data)
}

func (img Image) Width() int {
	return len(img.data[0])
}

func (img Image) Print() {
	for _, row := range img.data {
		for _, r := range row {
			fmt.Print(string(r))
		}

		fmt.Println()
	}
}

func (pair Pair) Compute(spaceCols, spaceRows map[int]bool, expansion int) int {
	res := 0
	dx := aoc2023.Delta(pair.To.X - pair.From.X)
	dy := aoc2023.Delta(pair.To.Y - pair.From.Y)

	for pt := pair.From; pt != pair.To; res++ {
		if pt.X != pair.To.X {
			pt.X += dx

			if spaceCols[pt.X] {
				res += expansion - 1
			}

			continue
		}

		if pt.Y != pair.To.Y {
			pt.Y += dy

			if spaceRows[pt.Y] {
				res += expansion - 1
			}

			continue
		}
	}

	return res
}
