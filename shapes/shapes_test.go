package shapes

import "testing"

func assert(t testing.TB, got, want float64) {
  t.Helper()
  if got != want {
    t.Errorf("got %g, want %g", got, want)
  }
}

func TestPerimeter(t *testing.T) {
  t.Run("rectangle perimeter", func(t *testing.T) {
    r := Rectangle{10, 10}
    got := r.Perimeter()
    want := 40.0

    assert(t, got, want)
  })
}

func TestArea(t *testing.T) {

  areaTests := []struct {
    shape Shape
    want float64
  }{
    {Rectangle{12, 6}, 72.0},
    {Circle{10}, 314.1592653589793},
  }

  for _, tt := range areaTests {
    got := tt.shape.Area()
    assert(t, got, tt.want)
  }

}

