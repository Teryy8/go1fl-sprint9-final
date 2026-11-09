package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	res := generateRandomElements(0)

	assert.Empty(t, res)
	assert.Len(t, res, 0)

	res = generateRandomElements(-1 * SIZE)

	assert.Empty(t, res)
	assert.Len(t, res, 0)

	res = generateRandomElements(SIZE)

	assert.NotEmpty(t, res)
	assert.Len(t, res, SIZE)
	for _, el := range res {
		assert.True(t, el > 0)
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{data: nil, want: 0},
		{data: []int{}, want: 0},
		{data: []int{7}, want: 7},
		{data: []int{9, 2, 3}, want: 9},
		{data: []int{2, 9, 3}, want: 9},
		{data: []int{2, 3, 9}, want: 9},
		{data: []int{5, 5, 5}, want: 5},
		{data: []int{1, 9, 3, 9}, want: 9},
	}

	for _, test := range tests {
		assert.Equal(t, maximum(test.data), test.want)
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{data: nil, want: 0},
		{data: []int{}, want: 0},
		{data: []int{7}, want: 7},
		{data: []int{1, 2, 3, 4, 5, 6, 99}, want: 99},
		{data: []int{1, 2, 3, 4, 5, 6, 7, 99}, want: 99},
		{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, want: 99},
		{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 99}, want: 99},
		{data: []int{5, 5, 5, 5, 5, 5, 5, 5}, want: 5},
		{data: []int{1, 9, 3, 9, 2, 4, 9, 1}, want: 9},
		{data: []int{1, 99, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, want: 99},
		{data: []int{1, 1, 1, 99, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, want: 99},
		{data: []int{1, 1, 1, 1, 1, 99, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, want: 99},
		{data: []int{1, 1, 1, 1, 1, 1, 1, 99, 1, 1, 1, 1, 1, 1, 1, 1}, want: 99},
		{data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 99, 1, 1, 1, 1, 1, 1}, want: 99},
		{data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 99, 1, 1, 1, 1}, want: 99},
		{data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 99, 1, 1}, want: 99},
		{data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 99}, want: 99},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, maxChunks(test.data))
	}
}
