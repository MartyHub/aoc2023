package main

import (
	"fmt"
	"image"
	"regexp"
	"strconv"
	"unicode"

	"github.com/MartyHub/aoc2023"
)

type (
	Engine struct {
		area    image.Rectangle
		numbers []Number
		symbols map[image.Point]rune
	}

	Number struct {
		rect  image.Rectangle
		value int
	}
)

func (num Number) String() string {
	return fmt.Sprintf("%3d, X=[%3d, %3d[, Y=%3d", num.value, num.rect.Min.X, num.rect.Max.X, num.rect.Min.Y)
}

func (num Number) neighbors() []image.Point {
	res := []image.Point{
		image.Pt(num.rect.Min.X-1, num.rect.Min.Y-1),
		image.Pt(num.rect.Min.X-1, num.rect.Min.Y),
		image.Pt(num.rect.Min.X-1, num.rect.Min.Y+1),
		image.Pt(num.rect.Max.X, num.rect.Max.Y-1),
		image.Pt(num.rect.Max.X, num.rect.Max.Y),
		image.Pt(num.rect.Max.X, num.rect.Max.Y+1),
	}

	for x := num.rect.Min.X; x < num.rect.Max.X; x++ {
		res = append(res, image.Pt(x, num.rect.Min.Y-1))
		res = append(res, image.Pt(x, num.rect.Min.Y+1))
	}

	return res
}

func main() {
	part1()
	part2()
}

func part1() {
	aoc2023.Expect("Sample 1", sum1("day03/data/sample.txt"), 4361)
	aoc2023.Expect("Part 1", sum1("day03/data/input.txt"), 526404)
}

func sum1(filePath string) int {
	engine := parse(filePath)
	res := 0

	for _, num := range engine.numbers {
		part := false

		for _, nb := range num.neighbors() {
			if !nb.In(engine.area) {
				continue
			}

			if _, found := engine.symbols[nb]; found {
				part = true

				break
			}
		}

		if part {
			res += num.value
		}
	}

	return res
}

func part2() {
	aoc2023.Expect("Sample 2", sum2("day03/data/sample.txt"), 467835)
	aoc2023.Expect("Part 2", sum2("day03/data/input.txt"), 84399773)
}

func sum2(filePath string) int {
	engine := parse(filePath)
	res := 0

	for pt, r := range engine.symbols {
		if r != '*' {
			continue
		}

		gears := make([]Number, 0, 2)

		for _, num := range engine.numbers {
			for _, nb := range num.neighbors() {
				if !nb.In(engine.area) {
					continue
				}

				if nb == pt {
					gears = append(gears, num)

					if len(gears) == 2 {
						break
					}
				}
			}
		}

		if len(gears) == 2 {
			res += gears[0].value * gears[1].value
		}
	}

	return res
}

func parse(filePath string) Engine {
	res := Engine{
		area:    image.Rect(0, 0, 0, 0),
		symbols: make(map[image.Point]rune),
	}

	maxX, maxY := 0, 0
	re := regexp.MustCompile(`(\d+)`)

	for i, line := range aoc2023.MustReadLines(filePath) {
		maxX = max(maxX, len(line)-1)
		maxY = max(maxY, i)

		indexes := re.FindAllStringIndex(line, -1)

		for j, value := range re.FindAllString(line, -1) {
			res.numbers = append(res.numbers, Number{
				rect:  image.Rect(indexes[j][0], i, indexes[j][1], i),
				value: aoc2023.Must(strconv.Atoi(value)),
			})
		}

		for j, r := range line {
			if r == '.' || unicode.IsDigit(r) {
				continue
			}

			res.symbols[image.Pt(j, i)] = r
		}
	}

	res.area = image.Rect(0, 0, maxX, maxY)

	return res
}
