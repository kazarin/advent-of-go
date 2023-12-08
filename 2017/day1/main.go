package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reviewSequence(s string, offset int) int {
	i := 0
	res := 0
	for i < len(s) {
		nextIndex := (i + offset) % len(s)
		if s[i] == s[nextIndex] {
			n, _ := strconv.Atoi(string(s[i]))
			res += n
		}
		i += 1
	}
	return res
}

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	answer := reviewSequence(rows[0], 1)
	return answer
}

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	answer := reviewSequence(rows[0], len(rows[0])/2)
	return answer
}

func main() {
	fmt.Println(Part1("input.txt"))
	fmt.Println(Part2("input.txt"))
}
