package structsmethodsandinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, got)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, want float64, shape Shape) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Expected %g but got %g ", want, got)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		expected := 100.0
		checkArea(t, expected, rectangle)

	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.0}
		expected := 314.1592653589793
		checkArea(t, expected, circle)
	})
}
