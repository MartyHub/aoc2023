package main

import (
	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 114, parse("day09/data/sample.txt").Compute1)
	aoc2023.Expect("Part 1", 1819125966, parse("day09/data/input.txt").Compute1)
}

func part2() {
	aoc2023.Expect("Sample 2", 2, parse("day09/data/sample.txt").Compute2)
	aoc2023.Expect("Part 2", 1140, parse("day09/data/input.txt").Compute2)
}

func parse(filePath string) Report {
	var res Report

	for _, line := range aoc2023.MustReadLines(filePath) {
		res = append(res, aoc2023.ToInts(line))
	}

	return res
}

type (
	History []int
	Report  []History
)

func (hist History) Compute1() int {
	report := Report{hist}

	for !report.IsZero() {
		report = append(report, report.Last().Diff())
	}

	res := 0

	for i := len(report) - 2; i >= 0; i-- {
		res += report[i].Last()
	}

	return res
}

func (hist History) Compute2() int {
	report := Report{hist}

	for !report.IsZero() {
		report = append(report, report.Last().Diff())
	}

	res := 0

	for i := len(report) - 2; i >= 0; i-- {
		res = report[i].First() - res
	}

	return res
}

func (hist History) Diff() History {
	var res History

	for i := 0; i < len(hist)-1; i++ {
		res = append(res, hist[i+1]-hist[i])
	}

	return res
}

func (hist History) IsZero() bool {
	for _, v := range hist {
		if v != 0 {
			return false
		}
	}

	return true
}

func (hist History) First() int {
	return hist[0]
}

func (hist History) Last() int {
	return hist[len(hist)-1]
}

func (report Report) Compute1() int {
	var res int

	for _, hist := range report {
		res += hist.Compute1()
	}

	return res
}

func (report Report) Compute2() int {
	var res int

	for _, hist := range report {
		res += hist.Compute2()
	}

	return res
}

func (report Report) IsZero() bool {
	return report.Last().IsZero()
}

func (report Report) Last() History {
	return report[len(report)-1]
}
