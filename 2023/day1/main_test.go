package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 142, Part1("test1.txt"))
}
func Test_Part2(t *testing.T) {
	assert.Equal(t, 281, Part2("test2.txt"))
}

func Benchmark_Part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1("test1.txt")
	}
}
func Benchmark_Part2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part2("test2.txt")
	}
}
