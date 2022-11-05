package arrays

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sums two collection of any sizes", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9}, []int {0})
		want := []int{3, 9, 0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
