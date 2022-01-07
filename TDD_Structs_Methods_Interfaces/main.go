package interfaces

import "math"

// Rectangle interface
type Rectangle struct {
	width  float64
	height float64
}

// Area calc for rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Circle interface
type Circle struct {
	Radius float64
}

// Area calc for circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter of a rectangel or square when passed the length & height
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

// Shape interface
type Shape interface {
	Area() float64
}

// Area given two lengths perpendicular to each other of a rectangle or square returns the area
func Area(rectangle Rectangle) (area float64) {

	area = rectangle.width * rectangle.height

	return area
}
