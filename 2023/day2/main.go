package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var blueRe, _ = regexp.Compile("([0-9]+) blue")
var redRe, _ = regexp.Compile("([0-9]+) red")
var greenRe, _ = regexp.Compile("([0-9]+) green")

func proceed(row string) int {
	r, _ := regexp.Compile("Game (?P<id>[0-9]+)")
	id, _ := strconv.Atoi(r.FindStringSubmatch(row)[1])
	blue := 0
	for _, b := range blueRe.FindAllStringSubmatch(row, -1) {
		value, _ := strconv.Atoi(b[1])
		if value > 14 {
			return 0
		}
		blue += value
	}
	red := 0
	for _, r := range redRe.FindAllStringSubmatch(row, -1) {
		value, _ := strconv.Atoi(r[1])
		if value > 12 {
			return 0
		}
		red += value
	}
	green := 0
	for _, g := range greenRe.FindAllStringSubmatch(row, -1) {
		value, _ := strconv.Atoi(g[1])
		if value > 13 {
			return 0
		}
		green += value
	}
	return id
}

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	ids := 0
	for _, row := range rows[:len(rows)-1] {
		ids += proceed(row)
	}
	return ids
}

func proceed2(row string) int {
	blue := []int{}
	for _, b := range blueRe.FindAllStringSubmatch(row, -1) {
		value, _ := strconv.Atoi(b[1])
		blue = append(blue, value)
	}
	red := []int{}
	for _, r := range redRe.FindAllStringSubmatch(row, -1) {
		value, _ := strconv.Atoi(r[1])
		red = append(red, value)
	}
	green := []int{}
	for _, g := range greenRe.FindAllStringSubmatch(row, -1) {
		value, _ := strconv.Atoi(g[1])
		green = append(green, value)
	}
	return slices.Max(blue) * slices.Max(green) * slices.Max(red)
}
func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	power := 0
	for _, row := range rows[:len(rows)-1] {
		power += proceed2(row)
	}
	return power
}

func main() {
	fmt.Println(Part1("day2.txt"))
	fmt.Println(Part2("day2.txt"))
}
