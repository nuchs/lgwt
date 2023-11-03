package main

import "fmt"

const (
	french             = "French"
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	englishHelloPrefix = "Hello, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case "French":
		return frenchHelloPrefix
	case "Spanish":
		return spanishHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("World", ""))
}
