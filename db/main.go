package main

import (
	"fmt"
	"log"

	"github.com/albachteng/learn/db/store"
)

func main() {
	s, err := store.NewStore("./test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Set("a", "123")
	if err != nil {
		log.Fatal(err)
	}
	err = s.Set("b", "789")
	if err != nil {
		log.Fatal(err)
	}
	out, err := s.Get("a")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
