package main

import (
	"image"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 46, parse("day16/data/sample.txt").Compute1)
	aoc2023.Expect("Part 1", 7199, parse("day16/data/input.txt").Compute1)
}

func part2() {
	aoc2023.Expect("Sample 2", 51, parse("day16/data/sample.txt").Compute2)
	aoc2023.Expect("Part 2", 7438, parse("day16/data/input.txt").Compute2)
}

func parse(filePath string) Grid {
	return aoc2023.MustReadRunes(filePath)
}

type (
	Grid [][]rune

	Beam struct {
		Pt  image.Point
		Dir image.Point
	}
)

func (grid Grid) Compute1() int {
	return grid.Compute(Beam{Pt: image.Pt(-1, 0), Dir: image.Pt(1, 0)})
}

func (grid Grid) Compute2() int {
	res := 0

	for x := 0; x < grid.Width(); x++ {
		res = max(res, grid.Compute(Beam{Pt: image.Pt(x, -1), Dir: image.Pt(0, 1)}))
	}

	for y := 0; y < grid.Height(); y++ {
		res = max(res, grid.Compute(Beam{Pt: image.Pt(grid.Width(), y), Dir: image.Pt(-1, 0)}))
	}

	for x := 0; x < grid.Width(); x++ {
		res = max(res, grid.Compute(Beam{Pt: image.Pt(x, grid.Height()), Dir: image.Pt(0, -1)}))
	}

	for y := 0; y < grid.Height(); y++ {
		res = max(res, grid.Compute(Beam{Pt: image.Pt(-1, y), Dir: image.Pt(1, 0)}))
	}

	return res
}

func (grid Grid) Compute(start Beam) int {
	cache := make(map[Beam]bool)

	for queue := []Beam{start}; len(queue) > 0; {
		beam := queue[0].Move()
		queue = queue[1:]

		if !grid.Valid(beam.Pt) || cache[beam] {
			continue
		}

		cache[beam] = true

		switch grid[beam.Pt.Y][beam.Pt.X] {
		case '/':
			queue = append(queue, beam.Mirror1())
		case '\\':
			queue = append(queue, beam.Mirror2())
		case '|':
			queue = append(queue, beam.SplitVertical()...)
		case '-':
			queue = append(queue, beam.SplitHorizontal()...)
		default:
			queue = append(queue, beam)
		}
	}

	return Tiles(cache)
}

func (grid Grid) Valid(pt image.Point) bool {
	return pt.X >= 0 && pt.X < grid.Width() &&
		pt.Y >= 0 && pt.Y < grid.Height()
}

func (grid Grid) Height() int {
	return len(grid)
}

func (grid Grid) Width() int {
	return len(grid[0])
}

func (beam Beam) Mirror1() Beam {
	switch beam.Dir {
	case image.Pt(1, 0):
		return beam.WithDir(0, -1)
	case image.Pt(-1, 0):
		return beam.WithDir(0, 1)
	case image.Pt(0, -1):
		return beam.WithDir(1, 0)
	case image.Pt(0, 1):
		return beam.WithDir(-1, 0)
	default:
		panic("Mirror1")
	}
}

func (beam Beam) Mirror2() Beam {
	switch beam.Dir {
	case image.Pt(1, 0):
		return beam.WithDir(0, 1)
	case image.Pt(-1, 0):
		return beam.WithDir(0, -1)
	case image.Pt(0, -1):
		return beam.WithDir(-1, 0)
	case image.Pt(0, 1):
		return beam.WithDir(1, 0)
	default:
		panic("Mirror2")
	}
}

func (beam Beam) SplitHorizontal() []Beam {
	if beam.Dir.Y == 0 {
		return []Beam{beam}
	} else {
		return []Beam{
			beam.WithDir(-1, 0),
			beam.WithDir(1, 0),
		}
	}
}

func (beam Beam) SplitVertical() []Beam {
	if beam.Dir.X == 0 {
		return []Beam{beam}
	} else {
		return []Beam{
			beam.WithDir(0, -1),
			beam.WithDir(0, 1),
		}
	}
}

func (beam Beam) Move() Beam {
	return Beam{
		Pt:  beam.Pt.Add(beam.Dir),
		Dir: beam.Dir,
	}
}

func (beam Beam) WithDir(x, y int) Beam {
	return Beam{
		Pt:  beam.Pt,
		Dir: image.Pt(x, y),
	}
}

func Tiles(beams map[Beam]bool) int {
	pts := make(map[image.Point]bool)

	for beam := range beams {
		pts[beam.Pt] = true
	}

	return len(pts)
}
