package clockface

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)
			if c.angle != got {
				t.Fatalf("wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			if c.angle != got {
				t.Fatalf("wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func TestMinuteHandVector(t *testing.T) {
  cases := []struct{
    time time.Time
    point Point
  }{
    {simpleTime(0, 30, 0), Point{0, -1}},
    {simpleTime(0, 45, 0), Point{-1, 0}},
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T){
      got := minuteHandPoint(c.time)
      if !roughlyEqualPoint(c.point, got) {
        t.Fatalf("Wanted %v Point, but got %v", c.point, got)
      }
    })
  }
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("expected to find the second hand line %+v, in SVG %+v", c.line, svg.Line)
			}
		})
	}
}

// func TestSVGWriterMinuteHand(t *testing.T) {
// 	cases := []struct {
// 		time time.Time
// 		line Line
// 	}{
// 		{
// 			simpleTime(0, 0, 0),
// 			Line{150, 150, 150, 70},
// 		},
// 	}
// 	for _, c := range cases {
// 		t.Run(testName(c.time), func(t *testing.T) {
// 			b := bytes.Buffer{}
// 			SVGWriter(&b, c.time)
//
// 			svg := SVG{}
// 			xml.Unmarshal(b.Bytes(), &svg)
//
// 			if !containsLine(c.line, svg.Line) {
// 				t.Errorf("expected to find the minute hand line %+v, in SVG %+v", c.line, svg.Line)
// 			}
// 		})
// 	}
// }
//
func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	const precision = 1e-7
	return math.Abs(a-b) < precision
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}
