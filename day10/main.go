package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 4, parse("day10/data/sample1.txt").Compute1)
	aoc2023.Expect("Sample 2", 8, parse("day10/data/sample2.txt").Compute1)
	aoc2023.Expect("Part 1", 6903, parse("day10/data/input.txt").Compute1)
}

func part2() {
	aoc2023.Expect("Sample 3", 4, func() int { return parse("day10/data/sample3.txt").Compute2(4) })
	aoc2023.Expect("Sample 4", 4, func() int { return parse("day10/data/sample4.txt").Compute2(4) })
	aoc2023.Expect("Sample 5", 8, func() int { return parse("day10/data/sample5.txt").Compute2(8) })
	aoc2023.Expect("Sample 6", 10, func() int { return parse("day10/data/sample6.txt").Compute2(10) })
	aoc2023.Expect("Part 2", 265, func() int { return parse("day10/data/input.txt").Compute2(265) })
}

func parse(filePath string) Area {
	res := Area{Tiles: aoc2023.MustReadRunes(filePath)}

	for y, row := range res.Tiles {
		for x, r := range row {
			if r == 'S' {
				res.Start = image.Point{X: x, Y: y}

				break
			}
		}
	}

	return res
}

type (
	Area struct {
		Tiles [][]rune
		Start image.Point
	}
)

func (a Area) Compute1() int {
	return len(a.Loop()) / 2
}

func (a Area) Compute2(expected int) int {
	loop := a.Loop()

	a.MarkOutside(loop)

	for y, row := range a.Tiles {
		for x, r := range row {
			if r == 'O' {
				continue
			}

			if _, found := loop[image.Point{X: x, Y: y}]; found {
				continue
			}

			if a.In(row, x) {
				a.Tiles[y][x] = 'I'
			} else {
				a.Tiles[y][x] = 'O'
			}
		}
	}

	res := a.Enclosed()

	if res != expected {
		a.Print()
	}

	return res
}

func (a Area) In(row []rune, x int) bool {
	s := string(row[:x])

	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "L7", "|")
	s = strings.ReplaceAll(s, "FJ", "|")

	return strings.Count(s, "|")%2 == 1
}

func (a Area) MarkOutside(loop map[image.Point]struct{}) {
	maxX := a.Width() - 1

	for y, row := range a.Tiles {
		for x := range row {
			if _, found := loop[image.Point{X: x, Y: y}]; found {
				break
			}

			a.Tiles[y][x] = 'O'
		}

		for x := maxX; x >= 0; x-- {
			if _, found := loop[image.Point{X: x, Y: y}]; found {
				break
			}

			a.Tiles[y][x] = 'O'
		}
	}
}

func (a Area) Enclosed() int {
	res := 0

	for _, row := range a.Tiles {
		for _, r := range row {
			if r == 'I' {
				res++
			}
		}
	}

	return res
}

func (a Area) Print() {
	for _, row := range a.Tiles {
		for _, r := range row {
			fmt.Printf("%c", r)
		}

		fmt.Println()
	}
}

func (a Area) Loop() map[image.Point]struct{} {
	pt := a.StartNext()
	path := map[image.Point]struct{}{a.Start: {}, pt: {}}

	for found := true; found; pt, found = a.Next(pt, path) {
	}

	return path
}

func (a Area) Pipes(pt image.Point) []image.Point {
	res := make([]image.Point, 0, 2)

	switch a.Tiles[pt.Y][pt.X] {
	case '|':
		res = append(res,
			image.Point{X: pt.X, Y: pt.Y - 1},
			image.Point{X: pt.X, Y: pt.Y + 1},
		)
	case '-':
		res = append(res,
			image.Point{X: pt.X - 1, Y: pt.Y},
			image.Point{X: pt.X + 1, Y: pt.Y},
		)
	case 'L':
		res = append(res,
			image.Point{X: pt.X, Y: pt.Y - 1},
			image.Point{X: pt.X + 1, Y: pt.Y},
		)
	case 'J':
		res = append(res,
			image.Point{X: pt.X, Y: pt.Y - 1},
			image.Point{X: pt.X - 1, Y: pt.Y},
		)
	case '7':
		res = append(res,
			image.Point{X: pt.X, Y: pt.Y + 1},
			image.Point{X: pt.X - 1, Y: pt.Y},
		)
	case 'F':
		res = append(res,
			image.Point{X: pt.X, Y: pt.Y + 1},
			image.Point{X: pt.X + 1, Y: pt.Y},
		)
	}

	return res
}

func (a Area) Next(pt image.Point, path map[image.Point]struct{}) (image.Point, bool) {
	for _, next := range a.Pipes(pt) {
		if !a.Valid(next) {
			continue
		}

		if _, found := path[next]; found {
			continue
		}

		path[next] = struct{}{}

		return next, true
	}

	return image.Point{}, false
}

func (a Area) Neighbors(pt image.Point) []image.Point {
	res := make([]image.Point, 0, 8)

	for _, neighbor := range []image.Point{
		{X: pt.X - 1, Y: pt.Y},
		{X: pt.X - 1, Y: pt.Y - 1},
		{X: pt.X, Y: pt.Y - 1},
		{X: pt.X + 1, Y: pt.Y - 1},
		{X: pt.X + 1, Y: pt.Y},
		{X: pt.X + 1, Y: pt.Y + 1},
		{X: pt.X, Y: pt.Y + 1},
		{X: pt.X - 1, Y: pt.Y + 1},
	} {
		if a.Valid(neighbor) {
			res = append(res, neighbor)
		}
	}

	return res
}

func (a Area) StartNext() image.Point {
	for _, pt := range a.Neighbors(a.Start) {
		for _, next := range a.Pipes(pt) {
			if next == a.Start {
				return pt
			}
		}
	}

	panic("failed to find one start")
}

func (a Area) Height() int {
	return len(a.Tiles)
}

func (a Area) Width() int {
	return len(a.Tiles[0])
}

func (a Area) Valid(pt image.Point) bool {
	return pt.X >= 0 && pt.X < a.Width() && pt.Y >= 0 && pt.Y < a.Height()
}
