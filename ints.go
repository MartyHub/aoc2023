package aoc2023

import (
	"regexp"
	"strconv"
)

func ToInt(s string) int {
	return Must(strconv.Atoi(s))
}

func ToInts(line string) []int {
	strings := regexp.MustCompile(`(-?\d+)`).FindAllString(line, -1)

	if strings == nil {
		return nil
	}

	res := make([]int, len(strings))

	for i, s := range strings {
		res[i] = ToInt(s)
	}

	return res
}

func Delta(i int) int {
	switch {
	case i > 0:
		return 1
	case i < 0:
		return -1
	default:
		return 0
	}
}
