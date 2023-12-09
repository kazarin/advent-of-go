package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func onlyZeroes(a []int) bool {
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}

func buildSequences(row string) [][]int {
	re := regexp.MustCompile(`-?(\d)+`)
	values := []int{}

	for _, value := range re.FindAllString(row, -1) {
		n, _ := strconv.Atoi(value)
		values = append(values, n)
	}

	sequences := [][]int{}
	sequences = append(sequences, values)
	for {
		tmp := []int{}
		for i := 0; i < len(values)-1; i++ {
			tmp = append(tmp, values[i+1]-values[i])
		}
		if len(tmp) == 0 {
			tmp = []int{0}
		}
		sequences = append(sequences, tmp)
		values = tmp
		if onlyZeroes(tmp) {
			break
		}
	}
	return sequences
}

func extrapolateLast(sequences [][]int) {
	for i := len(sequences) - 1; i > 0; i-- {
		current := sequences[i]
		previous := &sequences[i-1]
		delta := current[len(current)-1]
		*previous = append(*previous, delta+(*previous)[len(*previous)-1])
	}
}

func extrapolateFirst(sequences [][]int) {
	for i := len(sequences) - 1; i > 0; i-- {
		current := sequences[i]
		previous := &sequences[i-1]
		delta := current[0]
		*previous = append([]int{(*previous)[0] - delta}, *previous...)
	}
}

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	res := 0
	for _, row := range rows[:len(rows)-1] {
		sequences := buildSequences(row)
		extrapolateLast(sequences)
		res += sequences[0][len(sequences[0])-1]

	}
	return res
}

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	res := 0
	for _, row := range rows[:len(rows)-1] {
		sequences := buildSequences(row)
		extrapolateFirst(sequences)
		res += sequences[0][0]

	}
	return res
}

func main() {
	fmt.Println(Part1("input.txt"))
	fmt.Println(Part2("input.txt"))
}
