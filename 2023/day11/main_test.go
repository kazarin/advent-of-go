package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 374, Part1("test1.txt"))
	assert.Equal(t, 9214785, Part1("input.txt"))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 82000210, Part2("test1.txt"))
	assert.Equal(t, 613686987427, Part2("input.txt"))
}

func Benchmark_Part12(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1("input.txt")
		Part2("input.txt")
	}
}
