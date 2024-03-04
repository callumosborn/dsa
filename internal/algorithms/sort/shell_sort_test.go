package sort

import (
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	tests := []struct {
		name     string
		A        []int
		expected []int
	}{
		{
			name:     "Empty Slice",
			A:        []int{},
			expected: []int{},
		},
		{
			name:     "Random Order",
			A:        []int{5, 1, 4, 2, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Decreasing Order",
			A:        []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Sorted Order",
			A:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	s := ShellSort{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A := tt.A
			s.sort(&A)

			if !reflect.DeepEqual(A, tt.expected) {
				t.Errorf("sort() = %v, expected %v", A, tt.expected)
			}
		})
	}
}
