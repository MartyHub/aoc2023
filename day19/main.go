package main

import (
	"regexp"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 19114, parse("day19/data/sample.txt").Compute1)
	aoc2023.Expect("Part 1", 432788, parse("day19/data/input.txt").Compute1)
}

func part2() {
	// aoc2023.Expect("Sample 2", 167409079868000, parse("day19/data/sample.txt").Compute2)
	// aoc2023.Expect("Part 2", 0, parse("day19/data/input.txt").Compute)
}

func parse(filePath string) System {
	var i int

	res := System{
		Workflows: make(map[string]Workflow),
	}

	lines := aoc2023.MustReadLines(filePath)
	re := regexp.MustCompile(`(\w+){(.*),(\w+)}`)
	reRule := regexp.MustCompile(`(\w)([<>])(\d+):(\w+)`)

	for ; lines[i] != ""; i++ {
		matches := re.FindStringSubmatch(lines[i])
		name := matches[1]
		dft := matches[3]
		wf := Workflow{Default: dft}

		for _, s := range strings.Split(matches[2], ",") {
			matches = reRule.FindStringSubmatch(s)

			wf.Rules = append(wf.Rules, Rule{
				Part: matches[1],
				Op:   matches[2],
				Val:  aoc2023.ToInt(matches[3]),
				Wf:   matches[4],
			})
		}

		if len(wf.Rules) == 1 && wf.Rules[0].Wf == wf.Default {
			wf.Rules = nil
		}

		res.Workflows[name] = wf
	}

	for i++; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		ints := aoc2023.ToInts(lines[i])

		res.Ratings = append(res.Ratings, Rating{
			"x": ints[0],
			"m": ints[1],
			"a": ints[2],
			"s": ints[3],
		})
	}

	return res
}

type (
	System struct {
		Ratings   []Rating
		Workflows map[string]Workflow
	}

	Rating map[string]int

	Workflow struct {
		Rules   []Rule
		Default string
	}

	Rule struct {
		Part string
		Op   string
		Val  int
		Wf   string
	}
)

func (sys System) Compute1() int {
	res := 0

	for _, rtg := range sys.Ratings {
		res += sys.Rate(rtg)
	}

	return res
}

func (sys System) Compute2() int {
	res := 0

	for x := 1; x < 4000; x++ {
		for m := 1; m < 4000; m++ {
			if m == x {
				continue
			}

			for a := 1; a < 4000; a++ {
				if a == x || a == m {
					continue
				}

				for s := 1; s < 4000; s++ {
					if s == x || s == m || s == a {
						continue
					}

					res += sys.Rate(Rating{
						"x": x,
						"m": m,
						"a": a,
						"s": s,
					})
				}
			}
		}
	}

	return res
}

func (sys System) Rate(rtg Rating) int {
	for wf := "in"; ; {
		wf = sys.Workflows[wf].Next(rtg)

		switch wf {
		case "A":
			return rtg.Compute()
		case "R":
			return 0
		}
	}
}

func (wf Workflow) Next(rtg Rating) string {
	for _, rule := range wf.Rules {
		if qty, found := rtg[rule.Part]; found {
			switch rule.Op {
			case "<":
				if qty < rule.Val {
					return rule.Wf
				}
			case ">":
				if qty > rule.Val {
					return rule.Wf
				}
			default:
				panic(rule)
			}
		}
	}

	return wf.Default
}

func (rtg Rating) Compute() int {
	return rtg["x"] + rtg["m"] + rtg["a"] + rtg["s"]
}
