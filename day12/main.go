package main

import (
	"fmt"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 21, parse("day12/data/sample.txt").Compute)
	aoc2023.Expect("Part 1", 7350, parse("day12/data/input.txt").Compute)
}

func part2() {
	aoc2023.Expect("Sample 2", 525152, parse("day12/data/sample.txt").Expand().Compute)
	aoc2023.Expect("Part 2", 200097286528151, parse("day12/data/input.txt").Expand().Compute)
}

func parse(filePath string) Rows {
	var res Rows

	for _, line := range aoc2023.MustReadLines(filePath) {
		if line == "" {
			continue
		}

		res = append(res, NewRow(line))
	}

	return res
}

type (
	Row struct {
		Springs string
		Groups  []int
	}

	Rows []Row
)

func NewRow(s string) Row {
	parts := strings.Fields(s)

	return Row{
		Springs: parts[0],
		Groups:  aoc2023.ToInts(parts[1]),
	}
}

func (row Row) Expand() Row {
	groups := make([]int, 0, len(row.Groups)*5)
	springs := ""

	for i := 0; i < 5; i++ {
		if i > 0 {
			springs += "?"
		}

		springs += row.Springs
		groups = append(groups, row.Groups...)
	}

	return Row{
		Springs: springs,
		Groups:  groups,
	}
}

func (row Row) String() string {
	return fmt.Sprintf("%s %v", row.Springs, row.Groups)
}

func (rows Rows) Compute() int {
	res := 0

	for _, row := range rows {
		res += row.Compute(make(map[string]int))
	}

	return res
}

func (rows Rows) Expand() Rows {
	for i, row := range rows {
		rows[i] = row.Expand()
	}

	return rows
}

func (row Row) Done() (bool, int) {
	if len(row.Springs) == 0 {
		if len(row.Groups) == 0 {
			return true, 1
		}

		return true, 0
	}

	if len(row.Groups) == 0 {
		for _, r := range row.Springs {
			if r == '#' {
				return true, 0
			}
		}

		return true, 1
	}

	return false, 0
}

func (row Row) SkipSprings(length int) Row {
	return Row{
		Springs: row.Springs[length:],
		Groups:  row.Groups,
	}
}

func (row Row) SkipFirstGroup() Row {
	firstGroupLen := row.Groups[0]
	springs := row.Springs[firstGroupLen:]

	if len(row.Groups) == 1 {
		return Row{Springs: springs}
	}

	if len(springs) == 0 || springs[0] == '#' {
		return Row{Groups: row.Groups[1:]}
	}

	return Row{
		Springs: springs[1:],
		Groups:  row.Groups[1:],
	}
}

func (row Row) ReplaceFirstSpring(s string) Row {
	return Row{
		Springs: s + row.Springs[1:],
		Groups:  row.Groups,
	}
}

func (row Row) StartWith(group int) bool {
	if len(row.Springs) < group {
		return false
	}

	for i := 0; i < group; i++ {
		if row.Springs[i] == '.' {
			return false
		}
	}

	return true
}

func (row Row) Compute(cache map[string]int) int {
	key := row.String()

	if res, found := cache[key]; found {
		return res
	}

	if done, res := row.Done(); done {
		cache[key] = res

		return res
	}

	var res int

	switch row.Springs[0] {
	case '.':
		res = row.SkipSprings(1).Compute(cache)
	case '#':
		if row.StartWith(row.Groups[0]) {
			res = row.SkipFirstGroup().Compute(cache)
		} else {
			res = 0
		}
	default:
		res = row.ReplaceFirstSpring(".").Compute(cache) +
			row.ReplaceFirstSpring("#").Compute(cache)
	}

	cache[key] = res

	return res
}
