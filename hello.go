package main

import "fmt"

type language int

const (
	english language = iota
	spanish
	french
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func greetingPrefix(lang language) (prefix string) {
	prefix = englishHelloPrefix
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return
}

func Hello(name string, lang language) string {
	prefix := greetingPrefix(lang)
	if name == "" {
		name = "World"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", english))
}
