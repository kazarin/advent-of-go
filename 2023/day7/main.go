package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type Hand struct {
	Kind  int
	Cards string
	Bid   int
}

func proceed(row string) *Hand {
	parts := strings.Fields(row)
	hand := parts[0]
	bid, _ := strconv.Atoi(parts[1])
	cards := map[string]int{}
	for _, card := range hand {
		cards[string(card)]++
	}
	values := []int{}
	for _, v := range cards {
		values = append(values, v)
	}
	sort.Ints(values)
	if equal(values, []int{5}) {
		return &Hand{Bid: bid, Kind: 6, Cards: hand}
	} else if equal(values, []int{1, 4}) {
		return &Hand{Bid: bid, Kind: 5, Cards: hand}
	} else if equal(values, []int{2, 3}) {
		return &Hand{Bid: bid, Kind: 4, Cards: hand}
	} else if equal(values, []int{1, 1, 3}) {
		return &Hand{Bid: bid, Kind: 3, Cards: hand}
	} else if equal(values, []int{1, 2, 2}) {
		return &Hand{Bid: bid, Kind: 2, Cards: hand}
	} else if equal(values, []int{1, 1, 1, 2}) {
		return &Hand{Bid: bid, Kind: 1, Cards: hand}
	} else {
		return &Hand{Bid: bid, Kind: 0, Cards: hand}
	}
}

var dict = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}
var dict2 = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

func proceed2(row string) *Hand {
	parts := strings.Fields(row)
	hand := parts[0]
	bid, _ := strconv.Atoi(parts[1])
	cards := map[string]int{}
	jokers := 0
	for _, card := range hand {
		if string(card) == "J" {
			jokers++
		}
		cards[string(card)]++
	}
	values := []int{}
	for _, v := range cards {
		values = append(values, v)
	}
	sort.Ints(values)
	if jokers == 4 {
		return &Hand{Bid: bid, Kind: 6, Cards: hand}
	} else if equal(values, []int{5}) {
		return &Hand{Bid: bid, Kind: 6, Cards: hand}
	} else if equal(values, []int{1, 4}) {
		kind := 5
		if jokers > 0 {
			kind = 6
		}

		return &Hand{Bid: bid, Kind: kind, Cards: hand}
	} else if equal(values, []int{2, 3}) {
		kind := 4
		if jokers > 0 {
			kind = 6
		}
		return &Hand{Bid: bid, Kind: kind, Cards: hand}
	} else if equal(values, []int{1, 1, 3}) {
		kind := 3
		if jokers > 0 {
			kind = 5
		}
		return &Hand{Bid: bid, Kind: kind, Cards: hand}
	} else if equal(values, []int{1, 2, 2}) {
		kind := 2
		if jokers > 0 {
			kind = 3 + jokers
		} else {
			kind = 2
		}

		return &Hand{Bid: bid, Kind: kind, Cards: hand}
	} else if equal(values, []int{1, 1, 1, 2}) {
		kind := 1
		if jokers > 0 {
			kind = 3
		}
		return &Hand{Bid: bid, Kind: kind, Cards: hand}
	} else {

		return &Hand{Bid: bid, Kind: 0 + jokers, Cards: hand}
	}
}
func strongestCardWithJokers(hand1, hand2 *Hand) bool {
	for i := 0; i < len(hand1.Cards); i++ {
		v1 := hand1.Cards[i]
		v2 := hand2.Cards[i]
		if v1 == v2 {
			continue
		} else {
			return dict2[string(v1)] < dict2[string(v2)]
		}
	}
	return false

}

func strongestCard(hand1, hand2 *Hand) bool {
	for i := 0; i < len(hand1.Cards); i++ {
		v1 := hand1.Cards[i]
		v2 := hand2.Cards[i]
		if v1 == v2 {
			continue
		} else {
			return dict[string(v1)] < dict[string(v2)]
		}
	}
	return false

}

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	hands := []*Hand{}
	for _, row := range rows[:len(rows)-1] {
		hand := proceed(row)
		hands = append(hands, hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Kind == hands[j].Kind {
			return strongestCard(hands[i], hands[j])

		} else {
			return hands[i].Kind < hands[j].Kind
		}
	})
	res := 0
	for i, hand := range hands {
		res += hand.Bid * (i + 1)

	}

	return res
}
func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	hands := []*Hand{}
	for _, row := range rows[:len(rows)-1] {
		hand := proceed2(row)
		hands = append(hands, hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Kind == hands[j].Kind {
			return strongestCardWithJokers(hands[i], hands[j])

		} else {
			return hands[i].Kind < hands[j].Kind
		}
	})
	res := 0
	for i, hand := range hands {
		res += hand.Bid * (i + 1)

	}

	return res
}

func main() {
	fmt.Println(Part1("day7.txt"))
	fmt.Println(Part2("day7.txt"))
}
