package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 5, Part1("test1.txt"))
	assert.Equal(t, 7864, Part1("input.txt"))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 4, Part2("test1.txt"))
	assert.Equal(t, 1695, Part2("input.txt"))
}

func Benchmark_Part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1("input.txt")
		//Part2("input.txt")
	}
}
