package intset

import (
	"bytes"
	"fmt"
)

/* TODO:
* Len(),
* Remove(x int),
* Clear(),
* Copy() *IntSet,
* AddAll(xs ...int)
* IntersectWith(t *IntSet)
* DifferenceWith(t *IntSet)
* SymmetricDifference(t *IntSet)
* Elems() []int
* 32/64 bit modification using 32<<(^uint(0) >> 63)
* */

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Union(t *IntSet) *IntSet {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
	return s
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i*j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
