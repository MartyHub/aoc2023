package aoc2023

func GCD(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
