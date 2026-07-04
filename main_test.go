package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(10)

	assert.Len(t, data, 10)

	data = generateRandomElements(0)
	assert.Len(t, data, 0)
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty", []int{}, 0},
		{"one element", []int{5}, 5},
		{"normal", []int{1, 3, 2, 10, 4}, 10},
		{"all same", []int{7, 7, 7}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty", []int{}, 0},
		{"one element", []int{42}, 42},
		{"normal", []int{1, 5, 2, 99, 3, 7, 8}, 99},
		{"big chunk", []int{10, 20, 5, 100, 1}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			assert.Equal(t, tt.want, got)
		})
	}
}
