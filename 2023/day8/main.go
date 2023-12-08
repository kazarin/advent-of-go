package main

import (
	"advent-of-go/util"
	"fmt"
	"os"
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

type Step struct {
	Left  string
	Right string
	Key   string
}

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	instructions := rows[0]
	dict := map[string]Step{}

	for _, row := range rows[2 : len(rows)-1] {
		key := row[:3]
		left := row[7:10]
		right := row[12:15]
		dict[key] = Step{Left: left, Right: right}
	}
	var next string
	current := dict["AAA"]

	i := 0
	for next != "ZZZ" {
		for _, step := range instructions {
			if step == 76 {
				next = current.Left
			} else {
				next = current.Right
			}
			current = dict[next]
			i++
		}
	}
	return i
}
func cycle(slice string) func() byte {
	i := 0
	return func() byte {
		value := slice[i]
		i = (i + 1) % len(slice)
		return value
	}

}

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	instructions := rows[0]
	dict := map[string]Step{}

	for _, row := range rows[2 : len(rows)-1] {
		key := row[:3]
		left := row[7:10]
		right := row[12:15]
		dict[key] = Step{Key: key, Left: left, Right: right}
	}
	ghosts := []string{}
	for _, step := range dict {
		if step.Key[2] == 65 {
			ghosts = append(ghosts, step.Key)
		}
	}
	values := []int{}

	for _, ghost := range ghosts {
		current := dict[ghost]
		i := 0
		step := util.Cycle(instructions)
		for {
			var next string
			if step() == 76 {
				next = current.Left
			} else {
				next = current.Right

			}
			current = dict[next]

			i++
			if next[2] == 90 {
				values = append(values, i)
				break
			}
		}
	}

	result := values[0]
	for _, num := range values[1:] {
		result = util.Lcm(result, num)
	}
	return result

}

func main() {
	fmt.Println(Part1("input.txt"))
	fmt.Println(Part2("input.txt"))
}
