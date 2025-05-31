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

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}
