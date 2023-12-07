package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/MartyHub/aoc2023"
)

func main() {
	part1()
	part2()
}

func part1() {
	const cardOrder = "23456789TJQKA"

	aoc2023.Expect("Sample 1", 6440, func() int {
		return parse("day07/data/sample.txt").Compute(1, cardOrder)
	})
	aoc2023.Expect("Part 1", 253638586, func() int {
		return parse("day07/data/input.txt").Compute(1, cardOrder)
	})
}

func part2() {
	const cardOrder = "J23456789TQKA"

	aoc2023.Expect("Sample 2", 5905, func() int {
		return parse("day07/data/sample.txt").Compute(2, cardOrder)
	})
	aoc2023.Expect("Part 2", 253253225, func() int {
		return parse("day07/data/input.txt").Compute(2, cardOrder)
	})
}

func parse(filePath string) Hands {
	var res Hands

	for _, line := range aoc2023.MustReadLines(filePath) {
		res = append(res, Hand{
			Bid:   aoc2023.ToInt(line[6:]),
			Cards: line[:5],
		})
	}

	return res
}

type (
	Hand struct {
		Cards string
		Bid   int
	}

	Hands []Hand

	HandType int
)

func (hand Hand) Kinds() map[uint8]int {
	res := make(map[uint8]int, 5)

	for i := 0; i < 5; i++ {
		res[hand.Cards[i]]++
	}

	return res
}

func (hand Hand) Strength(part int) HandType {
	switch part {
	case 1:
		return hand.Strength1(false)
	case 2:
		return hand.Strength2()
	default:
		panic(fmt.Sprintf("invalid part %d", part))
	}
}

func (hand Hand) Strength1(joker bool) HandType {
	kinds := hand.Kinds()
	three := false
	pairs := 0

	for card, kind := range kinds {
		if joker && card == 'J' {
			continue
		}

		if kind == 5 {
			return FiveKind
		}

		if kind == 4 {
			return FourKind
		}

		if kind == 3 {
			three = true
		}

		if kind == 2 {
			pairs++
		}
	}

	if three && pairs == 1 {
		return FullHouse
	}

	if three {
		return ThreeKind
	}

	if pairs == 2 {
		return TwoPairs
	}

	if pairs == 1 {
		return Pair
	}

	return HighCard
}

func (hand Hand) Strength2() HandType {
	joker := strings.Count(hand.Cards, "J")

	if joker == 0 {
		return hand.Strength1(false)
	}

	strength1 := hand.Strength1(true)

	switch strength1 {
	case FiveKind:
		return FiveKind
	case FourKind:
		return FiveKind
	case FullHouse:
		return FullHouse
	case ThreeKind:
		switch joker {
		case 1:
			return FourKind
		case 2:
			return FiveKind
		}
	case TwoPairs:
		return FullHouse
	case Pair:
		switch joker {
		case 1:
			return ThreeKind
		case 2:
			return FourKind
		case 3:
			return FiveKind
		}
	default:
		switch joker {
		case 1:
			return Pair
		case 2:
			return ThreeKind
		case 3:
			return FourKind
		case 4, 5:
			return FiveKind
		}
	}

	panic(fmt.Sprintf("invalid case: cards=%v, strength1=%v, joker=%d", hand.Cards, strength1, joker))
}

func (hands Hands) Compute(part int, cardOrder string) int {
	slices.SortFunc(hands, func(a, b Hand) int {
		if res := a.Strength(part) - b.Strength(part); res != 0 {
			return int(res)
		}

		for i := 0; i < 5; i++ {
			if res := strings.IndexRune(cardOrder, rune(a.Cards[i])) -
				strings.IndexRune(cardOrder, rune(b.Cards[i])); res != 0 {
				return res
			}
		}

		panic(fmt.Sprintf("Same hand %v and %v", a, b))
	})

	return hands.Total()
}

func (hands Hands) Total() int {
	res := 0

	for i, hand := range hands {
		res += (i + 1) * hand.Bid
	}

	return res
}

func (ht HandType) String() string {
	switch ht {
	case HighCard:
		return "High Card"
	case Pair:
		return "Pair"
	case TwoPairs:
		return "Two Pairs"
	case ThreeKind:
		return "Three of a Kind"
	case FullHouse:
		return "Full House"
	case FourKind:
		return "Four of a Kind"
	case FiveKind:
		return "Five of a Kind"
	default:
		return ""
	}
}

const (
	HighCard HandType = iota
	Pair
	TwoPairs
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)
