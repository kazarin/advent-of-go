package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 6440, Part1("test1.txt"))
	assert.Equal(t, 250120186, Part1("day7.txt"))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 5905, Part2("test1.txt"))
	assert.Equal(t, 250665248, Part2("day7.txt"))
}

func Benchmark_Part12(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1("day7.txt")
		Part2("day7.txt")
	}
}

//func Benchmark_Part2(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		Part2("day4.txt")
//	}
//}
