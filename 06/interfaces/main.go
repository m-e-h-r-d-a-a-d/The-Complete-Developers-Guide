package main

import "fmt"

type bot interface {
	getGreeting() string
}

type (
	englishBot struct{}
	spanishBot struct{}
)

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	PrintGreeting(eb)
	PrintGreeting(sb)
}

func PrintGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
