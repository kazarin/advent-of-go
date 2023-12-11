package main

import (
	"advent-of-go/util"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func distribute(values []int) []int {
	updated := make([]int, len(values))
	copy(updated, values)

	bi := 0
	biggest := slices.Max(values)
	for i := 0; i < len(values); i++ {
		if values[i] == biggest {
			bi = i
			break
		}
	}

	updated[bi] = 0

	for biggest != 0 {
		bi += 1
		bi %= len(values)
		updated[bi]++
		biggest--
	}
	return updated
}
func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	re := regexp.MustCompile(`(\d+)`)

	values := []int{}
	for _, n := range re.FindAllString(rows[0], -1) {
		number, _ := strconv.Atoi(n)
		values = append(values, number)
	}
	seen := map[string]int{}

	steps := 0
	for {
		key := fmt.Sprint(values)
		if seen[key] > 0 {
			break
		}
		seen[key]++
		steps++
		values = distribute(values)
	}

	return steps
}

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	re := regexp.MustCompile(`(\d+)`)

	values := []int{}
	for _, n := range re.FindAllString(rows[0], -1) {
		number, _ := strconv.Atoi(n)
		values = append(values, number)
	}
	slow := values
	fast := distribute(values)

	for {
		slow = distribute(slow)
		fast = distribute(distribute(fast))
		if util.Equal(slow, fast) {
			break
		}
	}

	loops := 0
	for {
		slow = distribute(slow)
		loops++
		if util.Equal(slow, fast) {
			break
		}

	}

	return loops
}

func main() {
	fmt.Println(Part1("input.txt"))
	fmt.Println(Part2("input.txt"))
}
