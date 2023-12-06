package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func proceed(time int, distance int) int {
	count := 0
	for i := 1; i <= time; i++ {

		rem := time - i
		if rem*i > distance {
			count++

		}

	}
	return count
}

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	re := regexp.MustCompile(`(\d+)`)
	times := re.FindAllString(rows[0], -1)
	distances := re.FindAllString(rows[1], -1)
	res := 1
	for i := 0; i < len(distances); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		res *= proceed(t, d)
	}
	return res
}

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	re := regexp.MustCompile(`(\d+)`)
	times := re.FindAllString(rows[0], -1)
	distances := re.FindAllString(rows[1], -1)
	t, _ := strconv.Atoi(strings.Join(times, ""))
	d, _ := strconv.Atoi(strings.Join(distances, ""))
	return proceed(t, d)
}

func main() {
	fmt.Println(Part1("day6.txt"))
	fmt.Println(Part2("day6.txt"))
}
