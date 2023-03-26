package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

func main() {
	x := sha256.Sum256([]byte("a"))
	y := sha256.Sum256([]byte("XYA"))
	count := countBits(x, y)
	fmt.Printf("%x\n", x)
	fmt.Printf("%x\n", y)
	println(count)
}

func countBits(x, y [32]byte) int {
	if len(x) != sha256.Size || len(y) != sha256.Size {
		panic("inputs must be [32]byte")
	}
	count := 0
	for i := 0; i < sha256.Size; i++ {
		/* population count of array of bits yielded by XOR between the two*/
		diffBits := bits.OnesCount8(x[i] ^ y[i])
		count += diffBits
	}
	return count
}
