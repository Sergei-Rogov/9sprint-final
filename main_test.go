package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"zero size", 0},
		{"small size", 5},
		{"medium size", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			// проверка длины
			assert.Equal(t, tt.size, len(got))

			if tt.size == 0 {
				assert.Empty(t, got)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{10}, 10},
		{"normal case", []int{1, 5, 3, 9, 2}, 9},
		{"all same", []int{7, 7, 7, 7}, 7},
		{"negative not expected but safe", []int{-1, -5, -3}, -1},
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
		{"empty slice", []int{}, 0},
		{"one element", []int{42}, 42},
		{"small slice", []int{1, 2, 3, 4, 5}, 5},
		{"normal case", []int{10, 50, 3, 99, 7, 2, 88, 1}, 99},
		{"large uneven", []int{5, 1, 9, 3, 100, 2, 8}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			assert.Equal(t, tt.want, got)
		})
	}
}
