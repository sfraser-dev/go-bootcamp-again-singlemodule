package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// new type called deck of type sting slice
type theDeck []string

const (
	MySeperator string = ":"
)

// receiver: "extend" the functionality of the deck type via receiver
func (d theDeck) PrintWholeDeck() {
	for _, card := range d {
		fmt.Printf("%25s\n", card)
	}
}

// receiver
func (d theDeck) ShuffleDeck() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

// receiver
func (d theDeck) WriteToFile(filename string) {
	err := os.WriteFile(filename, []byte(ToSingleString(d, MySeperator)), 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// helper functions
func ReadFromFile(filename string) (myStrSlice []string) {
	var myByteSlice = []byte{}
	myByteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return FromSingleString(string(myByteSlice), MySeperator)
}

func NewDeck() theDeck {
	cards := theDeck{}
	// creating two LISTS, suits and numbers
	suits := [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}
	numbers := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range suits {
		for _, number := range numbers {
			tempStr := number + " of " + suit
			cards = append(cards, tempStr)
		}
	}
	return cards
}

func Deal(d theDeck, handSize int) (hand theDeck, remainingCards theDeck) {
	return d[:handSize], d[handSize:]
}

func ToSingleString(d theDeck, sep string) (theStr string) {
	// var s string = ""
	// for _, card := range d {
	// 	s = s + card + sep
	// }
	theStr = strings.Join(d, sep)
	return theStr
}

func FromSingleString(s string, sep string) (d theDeck) {
	d = strings.Split(s, sep)
	return d
}
