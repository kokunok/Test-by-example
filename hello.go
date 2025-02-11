package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	russian = "Russian"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Salut, "
	russianHelloPrefix = "Privet, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetiingPrefix(language) + name
}

func greetiingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
