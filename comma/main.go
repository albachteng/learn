package main

import (
	"bytes"
	"fmt"
)

func main() {
	thousand := "1000"
	tenMillion := "10000000"
	hundredBillion := "100000000000"
	fmt.Printf("%q, %q\n", comma(thousand), commaTwo(thousand))
	fmt.Printf("%q, %q\n", comma(tenMillion), commaTwo(tenMillion))
	fmt.Printf("%q, %q\n", comma(hundredBillion), commaTwo(hundredBillion))
}

// recursive solution
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaTwo(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	b := bytes.Buffer{}
	offset := n % 3
	b.Write([]byte(s[:offset]))
	for i := offset; i < n; i += 3 {
		if offset == 0 && i == offset {
			b.Write([]byte(s[i : i+3]))
		} else {
			b.Write([]byte("," + s[i:i+3]))
		}
	}
	return b.String()
}
