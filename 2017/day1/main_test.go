package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert.Equal(t, 3, reviewSequence("1122", 1))
	assert.Equal(t, 4, reviewSequence("1111", 1))
	assert.Equal(t, 0, reviewSequence("1234", 1))
	assert.Equal(t, 9, reviewSequence("91212129", 1))
	assert.Equal(t, 997, Part1("input.txt"))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 6, reviewSequence("1212", 2))
	assert.Equal(t, 0, reviewSequence("1221", 2))
	assert.Equal(t, 4, reviewSequence("123425", 3))
	assert.Equal(t, 12, reviewSequence("123123", 3))
	assert.Equal(t, 4, reviewSequence("12131415", 4))
	assert.Equal(t, 1358, Part2("input.txt"))
}

func Benchmark_Part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1("input.txt")
		Part2("input.txt")
	}
}
