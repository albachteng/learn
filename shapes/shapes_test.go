package shapes

import "testing"

func assert(t testing.TB, s Shape, got, want float64) {
  t.Helper()
  if got != want {
    t.Errorf("%#v: got %g, want %g", s, got, want)
  }
}

func TestPerimeter(t *testing.T) {

  perimeterTests := []struct {
    name string
    shape Shape
    hasPerimeter float64
  }{
    {name: "Rectangle", shape: Rectangle{10, 10}, hasPerimeter: 40.0},
    {name: "Circle", shape: Circle{5}, hasPerimeter: 31.41592653589793},
    {name: "Triangle", shape: Triangle{3, 4}, hasPerimeter: 11.54400374531753},
  }
  for _, tt := range perimeterTests {
    t.Run(tt.name, func(t *testing.T) {
      got := tt.shape.Perimeter()
      assert(t, tt.shape, got, tt.hasPerimeter)
    })
  }
}

func TestArea(t *testing.T) {

  areaTests := []struct {
    name string
    shape Shape
    hasArea float64
  }{
    {name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
    {name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
    {name: "Triangle", shape: Triangle{5, 10}, hasArea: 25.0},
  }

  for _, tt := range areaTests {
    t.Run(tt.name, func(t *testing.T) {
      got := tt.shape.Area()
      assert(t, tt.shape, got, tt.hasArea)
    })
  }

}

