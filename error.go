package aoc2023

import (
	"fmt"
	"time"
)

func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}

func Expect[T comparable](part string, expected T, fn func() T) {
	start := time.Now()
	got := fn()

	if got == expected {
		fmt.Printf("%s: %v in %v\n", part, got, time.Since(start))
	} else {
		fmt.Printf("%s: expected %v, got %v\n", part, expected, got)
	}
}
