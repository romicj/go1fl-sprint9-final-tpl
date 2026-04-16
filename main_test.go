package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		sizeIn       int
		expectedSize int
	}{
		{-33, 0},
		{-1, 0},
		{0, 0},
		{1, 1},
		{33, 33},
	}
	for _, test := range tests {
		assert.Len(t, generateRandomElements(test.sizeIn), test.expectedSize)
	}

}

func TestMaximum(t *testing.T) {
	tests := []struct {
		data        []int
		expectedMax int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{-11}, 0},
		{[]int{-3, 3}, 3},
		{[]int{0}, 0},
		{[]int{-1, 0, 3}, 3},
		{[]int{111, 22}, 111},
		{[]int{123, 126, 125}, 126},
	}
	for _, test := range tests {
		assert.Equal(t, maximum(test.data), test.expectedMax, fmt.Sprint("input data ", test.data))
	}

}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		data        []int
		expectedMax int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{-11}, 0},
		{[]int{-3, 3}, 3},
		{[]int{0}, 0},
		{[]int{-1, 0, 3}, 3},
		{[]int{111, 22}, 111},
		{[]int{123, 126, 125}, 126},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{[]int{1, 2, 3, 4, 5, 6, 7, 9, 10}, 10},
	}
	for _, test := range tests {
		assert.Equal(t, maxChunks(test.data), test.expectedMax, fmt.Sprint("input data ", test.data))
	}

}
