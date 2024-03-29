package main

import "fmt"

type bot interface {
	getGreeting(string, int) (string, error)
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// very custom logic for generating greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
