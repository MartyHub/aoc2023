package main

import "testing"

func TestPlatform_Cycle(t *testing.T) {
	pf := parse("data/sample.txt")
	cycle1 := parse("data/cycle1.txt")
	cycle2 := parse("data/cycle2.txt")
	cycle3 := parse("data/cycle3.txt")

	pf.Cycle()

	if !pf.Equals(cycle1) {
		t.Errorf("Cycle 1, expected:\n%v\nGot:\n%v\n", cycle1, pf)
	}

	pf.Cycle()

	if !pf.Equals(cycle2) {
		t.Errorf("Cycle 2, expected:\n%v\nGot:\n%v\n", cycle2, pf)
	}

	pf.Cycle()

	if !pf.Equals(cycle3) {
		t.Errorf("Cycle 3, expected:\n%v\nGot:\n%v\n", cycle3, pf)
	}
}
