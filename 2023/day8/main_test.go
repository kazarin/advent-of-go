package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 2, Part1("test1.txt"))
	assert.Equal(t, 6, Part1("test2.txt"))
	assert.Equal(t, 19667, Part1("input.txt"))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 6, Part2("test3.txt"))
	assert.Equal(t, 19185263738117, Part2("input.txt"))
}

func Benchmark_Part12(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1("input.txt")
		Part2("input.txt")
	}
}
