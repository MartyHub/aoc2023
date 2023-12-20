package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 32000000, parse("day20/data/sample1.txt").Compute)
	aoc2023.Expect("Sample 2", 11687500, parse("day20/data/sample2.txt").Compute)
	aoc2023.Expect("Part 1", 883726240, parse("day20/data/input.txt").Compute)
}

func part2() {
	// aoc2023.Expect("Part 2", 0, parse("day20/data/input.txt").Compute)
}

func parse(filePath string) Config {
	res := make(Config)

	re := regexp.MustCompile(`^([&%]?)(\w+) -> (.+)$`)

	for _, line := range aoc2023.MustReadLines(filePath) {
		if line == "" {
			continue
		}

		if matches := re.FindStringSubmatch(line); matches != nil {
			name := matches[2]

			res[name] = &Module{
				Pulses:  make(map[string]Kind),
				Symbol:  matches[1],
				Targets: strings.Split(matches[3], ", "),
			}
		}
	}

	for src, mdlSrc := range res {
		for _, target := range mdlSrc.Targets {
			if mdlTarget := res[target]; mdlTarget != nil && mdlTarget.Symbol == "&" {
				mdlTarget.Pulses[src] = pulseLow
			}
		}
	}

	return res
}

const (
	pulseNone Kind = ""
	pulseHigh Kind = "high"
	pulseLow  Kind = "low"
)

type (
	Config map[string]*Module

	Module struct {
		Targets []string
		Symbol  string

		On bool

		Pulses map[string]Kind
	}

	Kind string

	Pulse struct {
		Src    string
		Kind   Kind
		Target string
	}
)

func (cfg Config) Compute() int {
	countLow := 0
	countHigh := 0
	start := Pulse{
		Src:    "button",
		Kind:   pulseLow,
		Target: "broadcaster",
	}

	var queue []Pulse

	for i := 0; i < 1_000; i++ {
		queue = append(queue, start)

		for len(queue) > 0 {
			pulse := queue[0]
			queue = queue[1:]

			if pulse.Kind == pulseLow {
				countLow++
			} else {
				countHigh++
			}

			mdl := cfg[pulse.Target]
			if mdl == nil {
				continue
			}

			kind := mdl.Process(pulse.Src, pulse.Kind)
			if kind == pulseNone {
				continue
			}

			for _, target := range mdl.Targets {
				queue = append(queue, Pulse{
					Src:    pulse.Target,
					Kind:   kind,
					Target: target,
				})
			}
		}
	}

	return countLow * countHigh
}

func (mdl *Module) Process(src string, kind Kind) Kind {
	switch mdl.Symbol {
	case "":
		return kind
	case "%":
		if kind == pulseHigh {
			return pulseNone
		}

		mdl.On = !mdl.On

		if mdl.On {
			return pulseHigh
		}

		return pulseLow
	case "&":
		mdl.Pulses[src] = kind

		for _, val := range mdl.Pulses {
			if val == pulseLow {
				return pulseHigh
			}
		}

		return pulseLow
	default:
		panic(mdl.Symbol)
	}
}

func (pls Pulse) String() string {
	return fmt.Sprintf("%s -%s-> %s", pls.Src, pls.Kind, pls.Target)
}
