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
func validate(current []rune, previousRow string, currentRow string, nextRow string, right int, left int) int {
	n, _ := strconv.Atoi(string(current))
	if previousRow != "" {
		for _, t := range previousRow[left : right+1] {
			if !isDigit(byte(t)) && t != '.' {
				return n
			}
		}

	}
	if nextRow != "" {
		for _, t := range nextRow[left : right+1] {
			if !isDigit(byte(t)) && t != '.' {
				return n
			}
		}
	}
	if left > 1 {
		if currentRow[left] != '.' && !isDigit(currentRow[left]) {
			return n
		}
	}

	if right+1 < len(currentRow) {
		if currentRow[right] != '.' && !isDigit(currentRow[right]) {
			return n
		}
	}
	return 0

}

func Part1(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	res := 0
	for iRow, row := range rows {
		var nextRow, previousRow string
		if iRow+1 < len(rows) {
			nextRow = rows[iRow+1]
		}
		if iRow-1 >= 0 {
			previousRow = rows[iRow-1]
		}
		current := []rune{}
		number := false
		right := 0
		for _, sym := range row {
			if isDigit(byte(sym)) {
				current = append(current, sym)
				number = true

			} else {
				if number {
					left := max(0, right-len(current)-1)
					res += validate(current, previousRow, row, nextRow, right, left)
					current = nil
					number = false

				}
			}
			right++
		}
		if number {
			left := max(0, right-len(current)-1)
			res += validate(current, previousRow, row, nextRow, right-1, left)
		}

	}
	return res
}

func Part2(filename string) int {
	input, _ := os.ReadFile(filename)
	rows := strings.Split(string(input), "\n")
	res := 0
	for iRow, row := range rows {
		var nextRow, previousRow string
		if iRow+1 < len(rows) {
			nextRow = rows[iRow+1]
		}
		if iRow-1 >= 0 {
			previousRow = rows[iRow-1]
		}
		right := 0
		tmps := []string{}
		for _, sym := range row {
			if sym == '*' {
				if previousRow[right] == '.' {
					if isDigit(previousRow[right-1]) {
						tmp := []byte{}
						for i := right - 1; i >= 0; i-- {
							if isDigit(previousRow[i]) {
								tmp = append([]byte{previousRow[i]}, tmp...)
							} else {
								break
							}
						}
						tmps = append(tmps, string(tmp))

					}
					if isDigit(previousRow[right+1]) {
						tmp := []byte{}
						for i := right + 1; i <= len(row)-1; i++ {
							if isDigit(previousRow[i]) {
								tmp = append(tmp, previousRow[i])
							} else {
								break
							}
						}
						tmps = append(tmps, string(tmp))

					}
				} else if isDigit(previousRow[right]) {
					tmp := []byte{}
					for i := right - 1; i >= 0; i-- {
						if isDigit(previousRow[i]) {
							tmp = append([]byte{previousRow[i]}, tmp...)

						} else {
							break
						}
					}
					tmp = append(tmp, previousRow[right])
					for i := right + 1; i <= len(row)-1; i++ {
						if isDigit(previousRow[i]) {
							tmp = append(tmp, previousRow[i])

						} else {
							break
						}
					}
					tmps = append(tmps, string(tmp))
				}
				if nextRow[right] == '.' {
					if isDigit(nextRow[right-1]) {
						tmp := []byte{}
						for i := right - 1; i >= 0; i-- {
							if isDigit(nextRow[i]) {
								tmp = append([]byte{nextRow[i]}, tmp...)
							} else {
								break
							}
						}
						tmps = append(tmps, string(tmp))

					}
					if isDigit(nextRow[right+1]) {
						tmp := []byte{}
						for i := right + 1; i <= len(row)-1; i++ {
							if isDigit(nextRow[i]) {
								tmp = append(tmp, nextRow[i])
							} else {
								break
							}
						}
						tmps = append(tmps, string(tmp))

					}
				} else if isDigit(nextRow[right]) {
					tmp := []byte{}
					for i := right - 1; i >= 0; i-- {
						if isDigit(nextRow[i]) {
							tmp = append([]byte{nextRow[i]}, tmp...)

						} else {
							break
						}
					}
					tmp = append(tmp, nextRow[right])
					for i := right + 1; i <= len(row)-1; i++ {
						if isDigit(nextRow[i]) {
							tmp = append(tmp, nextRow[i])

						} else {
							break
						}
					}
					tmps = append(tmps, string(tmp))
				}

				if isDigit(row[right-1]) {
					tmp := []byte{}
					for i := right - 1; i >= 0; i-- {
						if isDigit(row[i]) {
							tmp = append([]byte{row[i]}, tmp...)
						} else {
							break
						}
					}
					tmps = append(tmps, string(tmp))

				}
				if isDigit(row[right+1]) {
					tmp := []byte{}
					for i := right + 1; i <= len(row)-1; i++ {
						if isDigit(row[i]) {
							tmp = append(tmp, row[i])
						} else {
							break
						}
					}
					tmps = append(tmps, string(tmp))

				}

			}

			right++
			if len(tmps) == 2 {
				a1, _ := strconv.Atoi(tmps[0])
				a2, _ := strconv.Atoi(tmps[1])
				res += a1 * a2
			}
			tmps = nil
		}
	}
	return res
}

func main() {
	fmt.Println(Part1("day3.txt"))
	fmt.Println(Part2("day3.txt"))
}
