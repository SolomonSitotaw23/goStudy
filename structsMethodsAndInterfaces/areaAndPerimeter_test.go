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

	//  Table driven tests
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 11, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("Expected %g but got %g shape %#v", tt.want, got, tt.shape)
		}
	}

	// this are same

	// checkArea := func(t *testing.T, want float64, shape Shape) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("Expected %g but got %g ", want, got)
	// 	}
	// }

	// t.Run("rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	expected := 100.0
	// 	checkArea(t, expected, rectangle)

	// })

	// t.Run("Circles", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	expected := 314.1592653589793
	// 	checkArea(t, expected, circle)
	// })
}
