package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func fetchDigit(row string) int {
	left := 0
	right := len(row) - 1
	digits := make([]byte, 2)
	for left <= right {
		if digits[0] != 0 && digits[1] != 0 {
			digit, _ := strconv.Atoi(string(digits))
			return digit
		}
		if isDigit(row[left]) {
			digits[0] = row[left]
		} else {
			left++
		}
		if isDigit(row[right]) {
			digits[1] = row[right]
		} else {
			right--
		}
	}
	return -1
}

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Fields(string(input))
	total := 0
	for _, row := range rows {
		digit := fetchDigit(row)
		total += digit
	}
	return total
}

var r1 = strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
var r2 = strings.NewReplacer("eno", "1", "owt", "2", "eerht", "3", "ruof", "4", "evif", "5", "xis", "6", "neves", "7", "thgie", "8", "enin", "9")

func reverse(str string) (result string) {
	for _, s := range str {
		result = string(s) + result
	}
	return
}

func fetchDigit2(row string) int {
	left := 0
	digits := make([]byte, 2)
	row = r1.Replace(row)
	reversedRow := r2.Replace(reverse(row))
	for left < len(row) {
		if isDigit(row[left]) {
			digits[0] = row[left]
			break
		}
		left++
	}

	left = 0

	for left < len(reversedRow) {
		if isDigit(reversedRow[left]) {
			digits[1] = reversedRow[left]
			break
		}
		left++

	}
	digit, _ := strconv.Atoi(string(digits))
	return digit
}
func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Fields(string(input))
	total := 0
	for _, row := range rows {
		digit := fetchDigit2(row)
		total += digit
	}
	return total
}

func main() {
	fmt.Println(Part1("day1.txt"))
}
