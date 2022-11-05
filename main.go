package main

import (
	"fmt"
)

const englishGreeting = "Hello, "

func Hello(name string) string {
  if name == "" {
    name = "World"
  }
	return englishGreeting + name
}

func main() {
	fmt.Print(Hello("Graham"))
}
