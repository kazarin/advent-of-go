package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	values := []int{}
	rows := strings.Split(string(input), "\n")
	for _, n := range rows[:len(rows)-1] {
		number, _ := strconv.Atoi(n)
		values = append(values, number)

	}
	i := 0
	steps := 0

	for {
		instruction := values[i]
		values[i]++
		steps++
		i = i + instruction
		if i > len(values)-1 {
			return steps
		}
	}
	return steps
}

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	values := []int{}
	rows := strings.Split(string(input), "\n")
	for _, n := range rows[:len(rows)-1] {
		number, _ := strconv.Atoi(n)
		values = append(values, number)

	}
	i := 0
	steps := 0

	for {
		instruction := values[i]
		if values[i] >= 3 {
			values[i]--
		} else {
			values[i]++
		}
		steps++
		i = i + instruction
		if i > len(values)-1 {
			return steps
		}
	}
	return steps
}
func main() {
	fmt.Println(Part1("input.txt"))
	fmt.Println(Part2("input.txt"))
}
