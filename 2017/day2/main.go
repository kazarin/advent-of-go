package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	res := 0
	for _, row := range rows[:len(rows)-1] {
		re := regexp.MustCompile(`(\d)+`)
		numbers := []int{}
		for _, n := range re.FindAllString(row, -1) {
			number, _ := strconv.Atoi(n)
			numbers = append(numbers, number)

		}
		res += (slices.Max(numbers) - slices.Min(numbers))

	}
	return res
}

func findEvenlyDivisible(numbers []int) int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			a, b := numbers[i], numbers[j]
			if b > a {
				a, b = b, a
			}
			if a%b == 0 {
				return a / b
			}
		}
	}
	return 0
}
func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	res := 0
	for _, row := range rows[:len(rows)-1] {
		re := regexp.MustCompile(`(\d)+`)
		numbers := []int{}
		for _, n := range re.FindAllString(row, -1) {
			number, _ := strconv.Atoi(n)
			numbers = append(numbers, number)

		}
		res += findEvenlyDivisible(numbers)

	}
	return res
}

func main() {
	fmt.Println(Part1("input.txt"))
	fmt.Println(Part2("input.txt"))
}
