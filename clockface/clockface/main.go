package main

import (
	"fmt"
	"os"
	"time"

	"github.com/albachteng/learn/clockface"
)

func main() {
  t := time.Now()
  clockface.SVGWriter(os.Stdout, t)
}

func secondHandTag(p clockface.Point) string {
  return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

