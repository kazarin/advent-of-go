package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 18, Part1("test1.txt"))
	assert.Equal(t, 43074, Part1("input.txt"))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 9, Part2("test2.txt"))
	assert.Equal(t, 280, Part2("input.txt"))
}

func Benchmark_Part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1("input.txt")
		Part2("input.txt")
	}
}
