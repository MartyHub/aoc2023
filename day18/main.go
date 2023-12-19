package main

import (
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 62, newDigger(parse1("day18/data/sample.txt")).Compute)
	aoc2023.Expect("Part 1", 48652, newDigger(parse1("day18/data/input.txt")).Compute)
}

func part2() {
	// aoc2023.Expect("Sample 2", 952408144115, newDigger(parse2("day18/data/sample.txt")).Compute)
	// aoc2023.Expect("Part 2", 0, parse("day18/data/input.txt").Compute1)
}

func parse1(filePath string) Plan {
	var res Plan

	for _, line := range aoc2023.MustReadLines(filePath) {
		if line == "" {
			continue
		}

		var ins Instruction

		aoc2023.MustSscanf(line, "%c %d", &ins.Dir, &ins.Len)

		res = append(res, ins)
	}

	return res
}

func parse2(filePath string) Plan {
	var res Plan

	for _, line := range aoc2023.MustReadLines(filePath) {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		color := fields[2]

		var dir uint8

		switch color[7] {
		case '0':
			dir = 'R'
		case '1':
			dir = 'D'
		case '2':
			dir = 'L'
		case '3':
			dir = 'U'
		default:
			panic(color)
		}

		res = append(res, Instruction{
			Dir: dir,
			Len: int(aoc2023.Must[int64](strconv.ParseInt(color[2:7], 16, 64))),
		})
	}

	return res
}

type (
	Plan []Instruction

	Instruction struct {
		Dir uint8
		Len int
	}

	Digger struct {
		outs     map[image.Point]bool
		rect     image.Rectangle
		trenches map[image.Point]bool
	}
)

func newDigger(plan Plan) Digger {
	outs := make(map[image.Point]bool, 150_000)
	trenches, rect := plan.Trenches()

	outs[rect.Min.Add(image.Pt(-1, -1))] = true

	for y := rect.Min.Y - 2; y <= rect.Max.Y+2; y++ {
		for x := rect.Min.X - 2; x <= rect.Max.X+2; x++ {
			pt := image.Pt(x, y)

			if trenches[pt] {
				break
			}

			outs[pt] = true
		}

		for x := rect.Max.X + 2; x >= rect.Min.X-2; x-- {
			pt := image.Pt(x, y)

			if trenches[pt] {
				break
			}

			outs[pt] = true
		}
	}

	return Digger{
		outs:     outs,
		rect:     rect,
		trenches: trenches,
	}
}

func (digger Digger) out(src image.Point) bool {
	cache := make(map[image.Point]struct{})
	dirs := []image.Point{
		image.Pt(-1, 0),
		image.Pt(1, 0),
		image.Pt(0, -1),
		image.Pt(0, 1),
	}
	queue := make([]image.Point, 1)

	cache[src] = struct{}{}
	queue[0] = src

	for len(queue) > 0 {
		pt := queue[0]
		queue = queue[1:]

		if out, found := digger.outs[pt]; found {
			for visited := range cache {
				if !digger.trenches[visited] {
					digger.outs[visited] = out
				}
			}

			return out
		}

		for _, dir := range dirs {
			next := pt.Add(dir)

			if digger.trenches[next] {
				continue
			}

			if _, found := cache[next]; found {
				continue
			}

			if next.X < digger.rect.Min.X-1 || next.Y < digger.rect.Min.Y-1 {
				continue
			}

			if next.X > digger.rect.Max.X+1 || next.Y > digger.rect.Max.Y+1 {
				continue
			}

			cache[next] = struct{}{}

			queue = append(queue, next)
		}
	}

	for pt := range cache {
		digger.outs[pt] = false
	}

	return false
}

func (digger Digger) Compute() int {
	res := 0

	for y := digger.rect.Min.Y; y <= digger.rect.Max.Y; y++ {
		for x := digger.rect.Min.X; x <= digger.rect.Max.X; x++ {
			if digger.trenches[image.Pt(x, y)] || !digger.out(image.Pt(x, y)) {
				res++
			}
		}
	}

	return res
}

func (plan Plan) Trenches() (map[image.Point]bool, image.Rectangle) {
	var pt image.Point

	rect := image.Rectangle{
		Min: image.Pt(math.MaxInt, math.MaxInt),
		Max: image.Pt(math.MinInt, math.MinInt),
	}
	res := make(map[image.Point]bool)

	for _, ins := range plan {
		for i := 0; i < ins.Len; i++ {
			switch ins.Dir {
			case 'U':
				pt = pt.Add(image.Pt(0, -1))
			case 'R':
				pt = pt.Add(image.Pt(1, 0))
			case 'D':
				pt = pt.Add(image.Pt(0, 1))
			case 'L':
				pt = pt.Add(image.Pt(-1, 0))
			default:
				panic(ins.Dir)
			}

			rect.Min.X = min(rect.Min.X, pt.X)
			rect.Min.Y = min(rect.Min.Y, pt.Y)

			rect.Max.X = max(rect.Max.X, pt.X)
			rect.Max.Y = max(rect.Max.Y, pt.Y)

			res[pt] = true
		}
	}

	return res, rect.Canon()
}

func (ins Instruction) String() string {
	return fmt.Sprintf("%c %d", ins.Dir, ins.Len)
}
