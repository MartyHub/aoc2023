package main

import (
	"image"
	"log"
	"math"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 102, parse("day17/data/sample1.txt").Compute1)
	aoc2023.Expect("Part 1", 953, parse("day17/data/input.txt").Compute1)
}

func part2() {
	aoc2023.Expect("Sample 2", 94, parse("day17/data/sample1.txt").Compute2)
	aoc2023.Expect("Sample 2", 71, parse("day17/data/sample2.txt").Compute2)
	aoc2023.Expect("Part 2", 1180, parse("day17/data/input.txt").Compute2)
}

func parse(filePath string) Map {
	return aoc2023.MustReadInts(filePath)
}

type (
	Map [][]int

	Path struct {
		Pt       image.Point
		Dir      image.Point
		DirCount byte
		HeatLoss int
	}

	Key [5]byte
)

func (m Map) Compute1() int {
	dirsX := [2]image.Point{image.Pt(-1, 0), image.Pt(1, 0)}
	dirsY := [2]image.Point{image.Pt(0, -1), image.Pt(0, 1)}

	cache := make(map[Key]int, 500_000)   // 235 177
	queue := make([]Path, 0, 150_000_000) // 126 930 489

	res := math.MaxInt
	size := 0
	target := m.Last()

	queue = append(queue, Path{
		Dir: image.Pt(1, 0),
	})

	for len(queue) > 0 {
		path := queue[0]
		key := path.Key()

		queue = queue[1:]
		size++

		if heatLoss, found := cache[key]; found {
			if path.HeatLoss >= heatLoss {
				continue
			}
		}

		if path.Pt == target {
			res = min(res, path.HeatLoss)

			continue
		}

		cache[key] = path.HeatLoss

		if path.DirCount < 3 {
			if next := path.Pt.Add(path.Dir); m.Valid(next) {
				queue = append(queue, path.Straight(next, m[next.Y][next.X]))
			}
		}

		dirs := dirsX

		if path.Dir.Y == 0 {
			dirs = dirsY
		}

		for _, dir := range dirs {
			if next := path.Pt.Add(dir); m.Valid(next) {
				queue = append(queue, path.Turn(next, dir, m[next.Y][next.X]))
			}
		}
	}

	log.Printf("Cache size: %d", len(cache))
	log.Printf("Paths size: %d", size)

	return res
}

func (m Map) Compute2() int {
	dirsX := [2]image.Point{image.Pt(-1, 0), image.Pt(1, 0)}
	dirsY := [2]image.Point{image.Pt(0, -1), image.Pt(0, 1)}

	cache := make(map[Key]int, 1_000_000) // 764 181
	queue := make([]Path, 0, 200_000_000) // 166 004 557

	res := math.MaxInt
	size := 0
	target := m.Last()

	queue = append(queue, Path{
		Dir: image.Pt(1, 0),
	})

	for len(queue) > 0 {
		path := queue[0]
		key := path.Key()

		queue = queue[1:]
		size++

		if heatLoss, found := cache[key]; found {
			if path.HeatLoss >= heatLoss {
				continue
			}
		}

		if path.Pt == target {
			if path.DirCount >= 4 {
				res = min(res, path.HeatLoss)
			}

			continue
		}

		cache[key] = path.HeatLoss

		if path.DirCount < 10 {
			if next := path.Pt.Add(path.Dir); m.Valid(next) {
				queue = append(queue, path.Straight(next, m[next.Y][next.X]))
			}
		}

		if path.DirCount >= 4 {
			dirs := dirsX

			if path.Dir.Y == 0 {
				dirs = dirsY
			}

			for _, dir := range dirs {
				if next := path.Pt.Add(dir); m.Valid(next) {
					queue = append(queue, path.Turn(next, dir, m[next.Y][next.X]))
				}
			}
		}
	}

	log.Printf("Cache size: %d", len(cache))
	log.Printf("Paths size: %d", size)

	return res
}

func (m Map) Last() image.Point {
	return image.Pt(m.Width()-1, m.Height()-1)
}

func (m Map) Height() int {
	return len(m)
}

func (m Map) Width() int {
	return len(m[0])
}

func (m Map) Valid(pt image.Point) bool {
	return pt.X >= 0 && pt.X < m.Width() &&
		pt.Y >= 0 && pt.Y < m.Height()
}

func (path Path) Key() Key {
	var buf Key

	buf[0] = byte(path.Pt.X)
	buf[1] = byte(path.Pt.Y)

	buf[2] = byte(path.Dir.X)
	buf[3] = byte(path.Dir.Y)

	buf[4] = path.DirCount

	return buf
}

func (path Path) Straight(pt image.Point, heatLoss int) Path {
	return Path{
		Pt:       pt,
		Dir:      path.Dir,
		DirCount: path.DirCount + 1,
		HeatLoss: path.HeatLoss + heatLoss,
	}
}

func (path Path) Turn(pt, dir image.Point, heatLoss int) Path {
	return Path{
		Pt:       pt,
		Dir:      dir,
		DirCount: 1,
		HeatLoss: path.HeatLoss + heatLoss,
	}
}
