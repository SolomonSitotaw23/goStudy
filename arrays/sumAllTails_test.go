package arrays

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, expected, got []int) {
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	}
	t.Run("make the sums of some slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{4, 3, 3, 1})
		expected := []int{9, 7}
		checkSums(t, expected, got)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{})
		expected := []int{0}
		checkSums(t, expected, got)
	})

}
