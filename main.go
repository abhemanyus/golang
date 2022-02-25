package main

import (
	"fmt"
)

const english = "English"
const german = "German"
const spanish = "Spanish"
const englishGreetingPrefix = "Hello, "
const englishNamePlaceholder = "World"
const germanGreetingPrefix = "Hallo, "
const germanNamePlaceholder = "Welt"
const spanishGreetingPrefix = "Hola, "
const spanishNamePlaceholder = "Tierra"

func greetingPrefix(language string) (greeting, placeholder string) {
	switch language {
	case english:
		greeting, placeholder = englishGreetingPrefix, englishNamePlaceholder
	case german:
		greeting, placeholder = germanGreetingPrefix, germanNamePlaceholder
	case spanish:
		greeting, placeholder = spanishGreetingPrefix, spanishNamePlaceholder
	default:
		greeting, placeholder = englishGreetingPrefix, englishNamePlaceholder
	}
	return
}

func Hello(name, language string) string {
	greeting, placeholder := greetingPrefix(language)

	if name == "" {
		return greeting + placeholder
	}
	return greeting + name
}

func main() {
	fmt.Println(Hello("Sanndy", "English"))
}
