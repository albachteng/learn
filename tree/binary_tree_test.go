package tree

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
  b := NewBST[int]()
  b.val = 20
  ints := b.appendValues(int[]{10, 25, 30, 5, 50}, &b)
  fmtPrintln(ints)
}
