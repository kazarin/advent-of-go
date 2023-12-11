package main

import (
	"advent-of-go/util"
	"fmt"
	"os"
	"strings"
)

func distance(a, b []int) int {
	return util.Abs(a[0]-b[0]) + util.Abs(a[1]-b[1])
}

func buildUniverse(rows []string, offset int) int {
	galaxies := [][]int{}
	emptyRows := []int{}
	emptyCols := []int{}
	for i, row := range rows[:len(rows)-1] {
		empty := true
		for j, col := range row {
			if col == '#' {
				galaxies = append(galaxies, []int{i, j})
				empty = false
			}
		}
		if empty {
			emptyRows = append(emptyRows, i)
		}

	}

	for i := 0; i < len(rows[0]); i++ {
		empty := true
		for _, v := range rows[:len(rows)-1] {
			if v[i] == '#' {
				empty = false

			}
		}
		if empty {
			emptyCols = append(emptyCols, i)
		}

	}

	answer := 0
	for i := range galaxies {
		j := i + 1
		for j < len(galaxies) {
			d := distance(galaxies[i], galaxies[j])
			x1, x2 := galaxies[i][0], galaxies[j][0]
			y1, y2 := galaxies[i][1], galaxies[j][1]
			if x1 > x2 {
				x2, x1 = x1, x2
			}
			if y1 > y2 {
				y2, y1 = y1, y2
			}
			for _, r := range emptyRows {
				if x1 < r && r < x2 {
					d += offset
				}
			}

			for _, r := range emptyCols {
				if y1 < r && r < y2 {
					d += offset
				}
			}
			answer += d

			j++
		}
	}

	return answer
}
func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	result := buildUniverse(rows, 1)
	return result
}

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	result := buildUniverse(rows, 999999)
	return result
}

func main() {
	fmt.Println(Part1("input.txt"))
	fmt.Println(Part2("input.txt"))
}
