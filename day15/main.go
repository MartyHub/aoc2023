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
	aoc2023.Expect("Sample 1", 52, parse("day15/data/sample1.txt").Compute1)
	aoc2023.Expect("Sample 2", 1320, parse("day15/data/sample2.txt").Compute1)
	aoc2023.Expect("Part 1", 515210, parse("day15/data/input.txt").Compute1)
}

func part2() {
	aoc2023.Expect("Sample 2", 145, parse("day15/data/sample2.txt").Compute2)
	aoc2023.Expect("Part 2", 246762, parse("day15/data/input.txt").Compute2)
}

func parse(filePath string) Steps {
	steps := strings.Split(strings.ReplaceAll(aoc2023.MustRead(filePath), "\n", ""), ",")

	res := make(Steps, len(steps))

	for i, s := range steps {
		res[i] = Step(s)
	}

	return res
}

type (
	Step  string
	Steps []Step

	Lens struct {
		Label string
		Focal int
	}

	Box   []Lens
	Boxes []Box
)

func (step Step) Compute() int {
	res := 0

	for _, r := range step {
		res += int(r)
		res *= 17
		res %= 256
	}

	return res
}

func (step Step) Lens() Lens {
	l := len(step)

	if step[l-1] == '-' {
		return Lens{Label: string(step[:l-1])}
	}

	return Lens{
		Label: string(step[:l-2]),
		Focal: int(step[l-1] - '0'),
	}
}

func (steps Steps) Compute1() int {
	res := 0

	for _, step := range steps {
		res += step.Compute()
	}

	return res
}

func (steps Steps) Compute2() int {
	boxes := make(Boxes, 256)

	for _, step := range steps {
		boxes.Process(step.Lens())
	}

	return boxes.Compute()
}

func (box Box) Compute(id int) int {
	res := 0

	for i, lens := range box {
		res += (id + 1) * (i + 1) * lens.Focal
	}

	return res
}

func (boxes Boxes) Compute() int {
	res := 0

	for i, box := range boxes {
		res += box.Compute(i)
	}

	return res
}

func (boxes Boxes) Process(lens Lens) {
	hash := Step(lens.Label).Compute()

	if lens.Focal == 0 {
		for i, l := range boxes[hash] {
			if l.Label == lens.Label {
				boxes[hash] = slices.Delete(boxes[hash], i, i+1)

				return
			}
		}

		return
	}

	for i, l := range boxes[hash] {
		if l.Label == lens.Label {
			boxes[hash][i] = lens

			return
		}
	}

	boxes[hash] = append(boxes[hash], lens)
}
