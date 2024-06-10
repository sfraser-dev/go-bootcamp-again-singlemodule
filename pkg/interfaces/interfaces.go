package interfaces

import (
	"fmt"
)

type LanguageBot interface {
	SayHi()
}

type EnglishBot struct {
	greeting string
}

type SpanishBot struct {
	greeting string
}

// method (has receiver) SayHi() attached to EnglishBot struct
func (eb EnglishBot) SayHi() {
	fmt.Println(eb.greeting)
}

// method (has receiver) SayHi() attached to SpanishBot struct
func (sb SpanishBot) SayHi() {
	fmt.Println(sb.greeting)
}

// func that takes an interface as an argument
func InterfaceAsArgument(lb LanguageBot) {
	lb.SayHi()
}

func RunMe() {
	// instantiate eb
	eb := EnglishBot{
		greeting: "Hello there!",
	}

	// instantiate sb
	sb := SpanishBot{
		greeting: "¡Hola!",
	}

	eb.SayHi()
	sb.SayHi()
	fmt.Println("...Interface as argument...")
	InterfaceAsArgument(eb)
	InterfaceAsArgument(sb)
}
