package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, wanted %d, given %v", got, want, numbers)
		}
	})
	t.Run("empty collection", func(t *testing.T) {
		numbers := []int{}
		got := Sum(numbers)
		want := 0

		if got != want {
			t.Errorf("got %d, wanted %d, given %v", got, want, numbers)
		}
	})
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
func TestSumAll(t *testing.T) {
	t.Run("two slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 5})
		want := []int{5, 14}

		checkSums(t, got, want)
	})
	t.Run("no tail", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{0, 9, 5})
		want := []int{0, 14}

		checkSums(t, got, want)
	})
	t.Run("empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 5})
		want := []int{0, 14}

		checkSums(t, got, want)
	})
}
