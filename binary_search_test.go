package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{"Existing Element at Beginning", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Existing Element in the Middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Existing Element at the End", []int{1, 2, 3, 4, 5}, 5, 4},
		{"Non-existing Element", []int{1, 2, 3, 4, 5}, 6, -1},
		{"Empty Slice", []int{}, 1, -1},
		{"Single Element Slice - Match", []int{5}, 5, 0},
		{"Single Element Slice - No Match", []int{5}, 3, -1},
		{"Two Element Slice - Match", []int{2, 7}, 7, 1},
		{"Two Element Slice - No Match", []int{2, 7}, 5, -1},
		{"Multiple Iterations", []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 15, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.input, tt.target)

			if result != tt.expected {
				t.Errorf("Expected %d, but got %d for input %v and target %d", tt.expected, result, tt.input, tt.target)
			}
		})
	}
}
