package interfaces

import (
	"fmt"
)

// interface defining REQUIRED signature "menu"
type Bot interface {
	GetGreeting() string
}

type EnglishBot struct {
	greeting string
}

type SpanishBot struct {
	greeting string
}

// method (has receiver) attached to EnglishBot struct (returns English greeting)
func (eb EnglishBot) GetGreeting() string {
	return eb.greeting
}

// method (has receiver) attached to SpanishBot struct (returns Spanish greeting)
func (sb SpanishBot) GetGreeting() string {
	return sb.greeting
}

// func that takes a bot INTERFACE as an argument
func InterfaceAsArgument(lb Bot) {
	var uniqueGreeting string = lb.GetGreeting()
	fmt.Println(uniqueGreeting)
}

func RunMe() {
	// instantiate an english bot
	eb := EnglishBot{
		greeting: "Hello there!",
	}

	// instantiate a spanish bot
	sb := SpanishBot{
		greeting: "¡Hola!",
	}

	fmt.Println("...Interface as argument...")
	InterfaceAsArgument(eb)
	InterfaceAsArgument(sb)
}
