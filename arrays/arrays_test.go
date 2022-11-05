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
	t.Run("sums any number of collections of any sizes", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9}, []int{0})
		want := []int{3, 9, 0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}
	t.Run("sums the tails of any number of collections", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9} // NOTE: a tail is all values excluding the first
		checkSums(t, got, want)
	})

	t.Run("sums the tails of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}
		checkSums(t, got, want)
	})
}
