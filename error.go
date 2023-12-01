package aoc2023

import "fmt"

func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}

func Expect(part string, got, expected any) {
	if got == expected {
		fmt.Printf("%s: %v\n", part, got)
	} else {
		fmt.Printf("%s: expected %v, got %v\n", part, expected, got)
	}
}
