package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 35, parse("day05/data/sample.txt").MinLoc1)
	aoc2023.Expect("Part 1", 278755257, parse("day05/data/input.txt").MinLoc1)
}

func part2() {
	aoc2023.Expect("Sample 2", 46, parse("day05/data/sample.txt").MinLoc2)
	aoc2023.Expect("Part 2", 26829166, parse("day05/data/input.txt").MinLoc2)
}

func parse(filePath string) *Almanac {
	var ranges Ranges

	res := &Almanac{Maps: make([]Ranges, 0, 7)}

	for _, line := range aoc2023.MustReadLines(filePath) {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds: ") {
			res.Seeds = aoc2023.ToInts(line[7:])

			continue
		}

		if strings.HasSuffix(line, "map:") {
			if len(ranges) > 0 {
				res.Maps = append(res.Maps, ranges)
			}

			ranges = make(Ranges, 0)

			continue
		}

		if ints := aoc2023.ToInts(line); ints != nil {
			ranges = append(ranges, Range{
				Src: ints[1],
				Dst: ints[0],
				Len: ints[2],
			})
		}
	}

	res.Maps = append(res.Maps, ranges)

	return res
}

type (
	Almanac struct {
		Seeds []int
		Maps  []Ranges
	}

	Range struct {
		Src int
		Dst int
		Len int
	}

	Ranges []Range

	Interval struct {
		From, To int
	}

	Intervals []Interval
)

func (alm *Almanac) MinLoc1() int {
	res := math.MaxInt

	for _, seed := range alm.Seeds {
		res = min(res, alm.Map(seed))
	}

	return res
}

func (alm *Almanac) MinLoc2() int {
	var itvs Intervals

	for i := 0; i < len(alm.Seeds); i += 2 {
		firstSeed := alm.Seeds[i]
		length := alm.Seeds[i+1]

		itvs = append(itvs, newInt(firstSeed, firstSeed+length-1))
	}

	for _, rngs := range alm.Maps {
		for i, pt := range rngs.Pts() {
			itvs = itvs.Split(pt, i%2 == 0)
		}

		for i, itv := range itvs {
			itvs[i] = newInt(rngs.Map(itv.From), rngs.Map(itv.To))
		}
	}

	return itvs.Min()
}

func (alm *Almanac) Map(seed int) int {
	res := seed

	for _, rngs := range alm.Maps {
		res = rngs.Map(res)
	}

	return res
}

func (rng Range) Map(i int) int {
	delta := i - rng.Src

	if delta >= 0 && delta < rng.Len {
		return rng.Dst + delta
	}

	return -1
}

func (rng Range) String() string {
	return fmt.Sprintf("%d->%d (+%d)", rng.Src, rng.Dst, rng.Len)
}

func (rngs Ranges) Pts() []int {
	res := make([]int, 0, len(rngs)*2)

	for _, rng := range rngs {
		res = append(res, rng.Src, rng.Src+rng.Len-1)
	}

	return res
}

func (rngs Ranges) Map(i int) int {
	for _, rng := range rngs {
		if res := rng.Map(i); res != -1 {
			return res
		}
	}

	return i
}

func newInt(from, to int) Interval {
	if from > to {
		panic(fmt.Sprintf("Invalid interval [%d, %d]", from, to))
	}

	return Interval{
		From: from,
		To:   to,
	}
}

func (itv Interval) Contains(i int) bool {
	return i > itv.From && i < itv.To
}

func (itv Interval) String() string {
	return fmt.Sprintf("[%d, %d]", itv.From, itv.To)
}

func (itvs Intervals) Min() int {
	res := math.MaxInt

	for _, itv := range itvs {
		res = min(res, itv.From)
	}

	return res
}

func (itvs Intervals) Split(pt int, start bool) Intervals {
	for i, itv := range itvs {
		if itv.Contains(pt) {
			res := make(Intervals, 0, len(itvs)+1)

			res = append(res, itvs[:i]...)

			if start {
				res = append(res, newInt(itv.From, pt-1), newInt(pt, itv.To))
			} else {
				res = append(res, newInt(itv.From, pt), newInt(pt+1, itv.To))
			}

			if i+1 < len(itvs) {
				res = append(res, itvs[i+1:]...)
			}

			return res
		}
	}

	return itvs
}
