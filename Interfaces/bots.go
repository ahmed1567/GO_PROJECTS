package main

import "fmt"

// bot is the interface: ANY type with a getGreeting() string method
// automatically satisfies it — no "implements" keyword needed in Go.
type bot interface {
	getGreeting() string
}

// greeter has the EXACT same method signature as bot. A type doesn't pick
// one interface to "belong to" — anything satisfying bot's method set
// automatically satisfies greeter too, and vice versa.
type greeter interface {
	getGreeting() string
}

// translator is a second, unrelated interface.
type translator interface {
	translate(text string) string
}

// multilingualBot embeds BOTH bot and translator — a type must satisfy
// every embedded interface's methods to satisfy this one.
type multilingualBot interface {
	bot
	translator
}

type englishBot struct{}
type spanishBot struct{}
type bilingualBot struct{}

func (englishBot) getGreeting() string {
	return "Hello, friend!"
}

func (spanishBot) getGreeting() string {
	return "Hola, amigo!"
}

func (bilingualBot) getGreeting() string {
	return "Hello! / Hola!"
}

func (bilingualBot) translate(text string) string {
	return "[translated] " + text
}

// One printGreeting for ALL bots — replaces per-type duplicate functions.
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Compile-time satisfaction checks: these do nothing at runtime — they
// exist so the compiler fails HERE, with a clear message, the moment a
// type stops satisfying an interface it's supposed to.
var _ bot = englishBot{}
var _ bot = spanishBot{}
var _ greeter = englishBot{}
var _ multilingualBot = bilingualBot{}
