package main

import (
	"fmt"
	"os"
	"strings"
)

type direction int

const (
	Up    direction = 0
	Down            = 1
	Left            = 2
	Right           = 3
)

type NextPosition struct {
	i, j      int
	Direction direction
}

func nextPosition(current NextPosition, s string) *NextPosition {
	switch s {
	case "|":
		if current.Direction == Up {
			return &NextPosition{i: current.i - 1, j: current.j, Direction: current.Direction}
		} else {
			return &NextPosition{i: current.i + 1, j: current.j, Direction: current.Direction}
		}
	case "-":
		if current.Direction == Left {
			return &NextPosition{i: current.i, j: current.j - 1, Direction: current.Direction}
		} else {
			return &NextPosition{i: current.i, j: current.j + 1, Direction: current.Direction}
		}
	case "L":
		if current.Direction == Down {
			return &NextPosition{i: current.i, j: current.j + 1, Direction: Right}
		} else {
			return &NextPosition{i: current.i - 1, j: current.j, Direction: Up}
		}
	case "J":
		if current.Direction == Down {
			return &NextPosition{i: current.i, j: current.j - 1, Direction: Left}
		} else {
			return &NextPosition{i: current.i - 1, j: current.j, Direction: Up}
		}
	case "7":
		if current.Direction == Right {
			return &NextPosition{i: current.i + 1, j: current.j, Direction: Down}
		} else {
			return &NextPosition{i: current.i, j: current.j - 1, Direction: Left}
		}
	case "F":
		if current.Direction == Up {
			return &NextPosition{i: current.i, j: current.j + 1, Direction: Right}
		} else {
			return &NextPosition{i: current.i + 1, j: current.j, Direction: Down}
		}
	}
	return &NextPosition{}
}

func buildGrid(rows []string, grid [][]string) (si, sj int) {
	for i, row := range rows[:len(rows)-1] {
		cols := strings.Split(row, "")
		grid[i] = make([]string, len(cols))
		for j, col := range cols {
			if col == "S" {
				si, sj = i, j
			}
			grid[i][j] = col
		}
	}

	return
}

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")

	grid := make([][]string, len(rows)-1)
	si, sj := buildGrid(rows, grid)

	next := &NextPosition{i: si, j: sj + 1, Direction: Right}
	steps := 1
	for {
		next = nextPosition(*next, grid[next.i][next.j])
		steps++
		if grid[next.i][next.j] == "S" {
			break
		}
	}
	return (steps + 1) / 2
}

type Point struct {
	I, J int
}

type Area []Point

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	grid := make([][]string, len(rows)-1)

	si, sj := buildGrid(rows, grid)

	next := &NextPosition{i: si, j: sj + 1, Direction: Right}

	area := Area{}

	for {
		next = nextPosition(*next, grid[next.i][next.j])
		area = append(area, Point{I: next.i, J: next.j})
		if grid[next.i][next.j] == "S" {
			break
		}
	}

	inside := 0
	for i := range grid {
		for j := range grid[i] {
			point := Point{I: i, J: j}
			if surrounded(point, area) == true {
				inside++
			}

		}
	}

	return inside
}

func surrounded(p Point, area Area) bool {
	surr := false
	for i, j := 0, len(area)-1; i < len(area); j, i = i, i+1 {
		if border(p, area[i], area[j]) {
			return false
		}
		if ((area[i].J > p.J) != (area[j].J > p.J)) && (p.I < (area[j].I-area[i].I)*(p.J-area[i].J)/(area[j].J-area[i].J)+area[i].I) {
			surr = !surr
		}
	}
	return surr
}

func border(p, a, b Point) bool {
	minI := min(a.I, b.I)
	maxI := max(a.I, b.I)
	minJ := min(a.J, b.J)
	maxJ := max(a.J, b.J)
	return p.I >= minI && p.I <= maxI && p.J >= minJ && p.J <= maxJ &&
		(a.I-p.I)*(b.J-p.J) == (b.I-p.I)*(a.J-p.J)
}

func main() {
	fmt.Println(Part1("input.txt"))
	fmt.Println(Part2("input.txt"))
}
