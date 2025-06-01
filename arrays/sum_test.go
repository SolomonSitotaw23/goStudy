package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 1, 1, 1, 1}
	got := Sum(numbers)
	expected := 5

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}
