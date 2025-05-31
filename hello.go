package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Chris", ""))
}

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Holla "
const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjur "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}
