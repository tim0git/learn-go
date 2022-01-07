package interfaces

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	recieved := Perimeter(rectangle)
	expected := 40.0

	if recieved != expected {
		t.Errorf("recieved %.2f expected %.2f", recieved, expected)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		recieved := shape.Area()

		if recieved != expected {
			t.Errorf("recieved %g expected %g", recieved, expected)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.00, 5.00}
		expected := 25.00

		checkArea(t, rectangle, expected)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.00}

		expected := 314.1592653589793

		checkArea(t, circle, expected)
	})
}
