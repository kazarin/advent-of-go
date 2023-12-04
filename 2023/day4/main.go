package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func proceed(row string) int {
	rows := strings.Split(row, ":")
	list := strings.Split(rows[1], "|")
	winningNumbers := strings.Split(strings.Trim(list[0], " "), " ")
	myNumbers := strings.Split(strings.Trim(list[1], " "), " ")
	dict := map[int]int{}
	for _, number := range winningNumbers {
		n, _ := strconv.Atoi(number)
		dict[n]++
	}
	matches := 0

	for _, number := range myNumbers {
		n, _ := strconv.Atoi(number)
		if dict[n] > 0 && n != 0 {
			matches++
		}
	}
	return int(math.Pow(float64(2), float64(matches-1)))

}

func proceed2(row string) []int {
	rows := strings.Split(row, ":")
	id, _ := strconv.Atoi(strings.Fields(rows[0])[1])
	list := strings.Split(rows[1], "|")
	winningNumbers := strings.Split(strings.Trim(list[0], " "), " ")
	myNumbers := strings.Split(strings.Trim(list[1], " "), " ")
	dict := map[int]int{}
	for _, number := range winningNumbers {
		n, _ := strconv.Atoi(number)
		dict[n]++
	}
	matches := 0

	for _, number := range myNumbers {
		n, _ := strconv.Atoi(number)
		if dict[n] > 0 && n != 0 {
			matches++
		}
	}
	return []int{id, matches}

}
func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	res := 0
	for _, row := range rows[:len(rows)-1] {
		res += proceed(row)
	}
	return res
}

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	instances := map[int]int{}
	for _, row := range rows[:len(rows)-1] {
		res := proceed2(row)
		id := res[0]
		matches := res[1]
		instances[id]++
		for j := 0; j < instances[id]; j++ {
			for i := id + 1; i <= id+matches; i++ {
				instances[i]++
			}
		}
	}

	res := 0
	for _, value := range instances {
		res += value
	}

	return res
}

func main() {
	fmt.Println(Part1("day4.txt"))
	fmt.Println(Part2("day4.txt"))
}
