package algorithms

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	A := []int{5, 2, 3, 1, 4}
	i := 0
	min := 3

	expectedVal := 1

	if swap(&A, i, min); A[i] != expectedVal {
		t.Errorf("Expected swap to put expectedVal %d into index %d", expectedVal, i)
	}
}

func TestLess(t *testing.T) {
	A := []int{5, 2, 3, 1, 4}
	a := 0
	b := 3

	if less(A[a], A[b]) {
		t.Error("Expected less to return false")
	}

	if !less(A[b], A[a]) {
		t.Error("Expected less to return true")
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want []int
	}{
		{
			name: "Test case 1",
			A:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			want: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			name: "Test case 2",
			A:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A := tt.A
			sort(&A)

			if !reflect.DeepEqual(A, tt.want) {
				t.Errorf("sort() = %v, want %v", A, tt.want)
			}
		})
	}
}
