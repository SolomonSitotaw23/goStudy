package arrays

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 1})
	expected := []int{15, 7}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v but got %v", expected, got)
	}

}
