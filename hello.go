package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Chris", ""))
}

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Holla "
	frenchHelloPrefix  = "Bonjur "
)

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
