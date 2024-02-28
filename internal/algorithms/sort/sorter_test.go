package sort

import (
	"errors"
	"reflect"
	"testing"
)

func TestLess(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected bool
	}{
		{
			name:     "A is less than B, positive numbers",
			a:        1,
			b:        5,
			expected: true,
		},
		{
			name:     "A is greater than B, positive numbers",
			a:        5,
			b:        1,
			expected: false,
		},
		{
			name:     "A is less than B, negative numbers",
			a:        -10,
			b:        -2,
			expected: true,
		},
		{
			name:     "A is greater than B, negative numbers",
			a:        -1,
			b:        -5,
			expected: false,
		},
		{
			name:     "A is equal to B",
			a:        1,
			b:        1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if less(tt.a, tt.b) != tt.expected {
				t.Errorf("Expected %v to be less than %v", tt.a, tt.b)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		name     string
		A        *[]int
		a        int
		b        int
		expected *[]int
		err      error
	}{
		{
			name:     "Basic Swap",
			A:        &[]int{1, 2, 3, 4},
			a:        1,
			b:        2,
			expected: &[]int{1, 3, 2, 4},
			err:      nil,
		},
		{
			name:     "Same Index Swap",
			A:        &[]int{1, 2, 3, 4},
			a:        1,
			b:        1,
			expected: &[]int{1, 2, 3, 4},
			err:      nil,
		},
		{
			name:     "First and Last Swap",
			A:        &[]int{1, 2, 3, 4},
			a:        0,
			b:        3,
			expected: &[]int{4, 2, 3, 1},
			err:      nil,
		},
		{
			name:     "Negative Index",
			A:        &[]int{1, 2, 3, 4},
			a:        -1,
			b:        3,
			expected: &[]int{1, 2, 3, 4},
			err:      errors.New("invalid index: indices cannot be negative"),
		},
		{
			name:     "Index Out-of-Range",
			A:        &[]int{1, 2, 3, 4},
			a:        1,
			b:        6,
			expected: &[]int{1, 2, 3, 4},
			err:      errors.New("invalid index: indices cannot be out-of-range"),
		},
		{
			name:     "Slice Empty",
			A:        &[]int{},
			a:        0,
			b:        0,
			expected: &[]int{},
			err:      errors.New("invalid slice: slices cannot have less than 2 elements"),
		},
		{
			name:     "Slice Too Small",
			A:        &[]int{1},
			a:        0,
			b:        0,
			expected: &[]int{1},
			err:      errors.New("invalid slice: slices cannot have less than 2 elements"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := swap(tt.A, tt.a, tt.b)

			if err != nil {

				if err.Error() != tt.err.Error() {
					t.Errorf("Expected %v, but got %v", tt.err.Error(), err.Error())
				}
			}

			if !reflect.DeepEqual(tt.A, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, tt.A)
			}
		})
	}
}
