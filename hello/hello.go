package hello

import (
	"fmt"
)

const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "
const frenchGreeting = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {

	switch language {
	case "spanish":
		prefix = spanishGreeting
	case "french":
		prefix = frenchGreeting
	default:
		prefix = englishGreeting
	}

	return
}

func main() {
	fmt.Print(Hello("Graham", ""))
}
