package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		A        []int
		expected []int
	}{
		{
			name:     "Random Order",
			A:        []int{15, 21, 20, 2, 15, 24, 5, 19},
			expected: []int{2, 5, 15, 15, 19, 20, 21, 24},
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
		{
			name:     "Empty Slice",
			A:        []int{},
			expected: []int{},
		},
	}

	s := QuickSort{}

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
