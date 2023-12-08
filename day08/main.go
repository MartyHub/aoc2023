package main

import (
	"fmt"
	"regexp"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", 2, parse("day08/data/sample1.txt").Compute1)
	aoc2023.Expect("Sample 2", 6, parse("day08/data/sample2.txt").Compute1)
	aoc2023.Expect("Part 1", 17287, parse("day08/data/input.txt").Compute1)
}

func part2() {
	aoc2023.Expect("Sample 3", 6, parse("day08/data/sample3.txt").Compute2)
	aoc2023.Expect("Part 2", 18625484023687, parse("day08/data/input.txt").Compute2)
}

func parse(filePath string) Network {
	res := Network{Nodes: make(map[Node]Nodes)}
	re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	for _, line := range aoc2023.MustReadLines(filePath) {
		if line == "" {
			continue
		}

		if res.Instructions == "" {
			res.Instructions = line

			continue
		}

		matches := re.FindStringSubmatch(line)

		res.Nodes[Node(matches[1])] = []Node{Node(matches[2]), Node(matches[3])}
	}

	return res
}

type (
	Network struct {
		Instructions string
		Nodes        map[Node]Nodes
	}

	Node   string
	Nodes  []Node
	Cycles []int
)

func (node Node) IsStart() bool {
	return node[2] == 'A'
}

func (node Node) IsEnd() bool {
	return node[2] == 'Z'
}

func (cycles Cycles) Done() bool {
	for _, cycle := range cycles {
		if cycle == 0 {
			return false
		}
	}

	return true
}

func (nw Network) Instruction(step int) int {
	dir := nw.Instructions[step%len(nw.Instructions)]

	switch dir {
	case 'L':
		return 0
	case 'R':
		return 1
	default:
		panic(fmt.Sprintf("unknown instruction %c", dir))
	}
}

func (nw Network) Compute1() int {
	steps := 0

	for node := Node("AAA"); node != "ZZZ"; steps++ {
		node = nw.Nodes[node][nw.Instruction(steps)]
	}

	return steps
}

func (nw Network) Start() Nodes {
	var res Nodes

	for node := range nw.Nodes {
		if node.IsStart() {
			res = append(res, node)
		}
	}

	return res
}

func (nw Network) Compute2() int {
	nodes := nw.Start()
	cycles := make(Cycles, len(nodes))

	for steps := 0; !cycles.Done(); {
		dir := nw.Instruction(steps)

		steps++

		for i, node := range nodes {
			if cycles[i] != 0 {
				continue
			}

			if nodes[i] = nw.Nodes[node][dir]; nodes[i].IsEnd() {
				cycles[i] = steps
			}
		}
	}

	for len(cycles) > 1 {
		cycles[1] = aoc2023.LCM(cycles[0], cycles[1])
		cycles = cycles[1:]
	}

	return cycles[0]
}
