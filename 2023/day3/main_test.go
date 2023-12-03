package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 4361, Part1("test1.txt"))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 467835, Part2("test1.txt"))
}

func Benchmark_Part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1("test1.txt")
	}
}

func Benchmark_Part2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part2("test1.txt")
	}
}
