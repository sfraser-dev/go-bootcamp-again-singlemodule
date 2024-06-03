package main

import (
	"fmt"
	"math/rand"

	deck "bitbucket.org/weebucket/gobootcampagain/cards/reusable"
	odd_even "bitbucket.org/weebucket/gobootcampagain/odd_even/reusable" // we import modules which are a collection of packages
)

func main() {
	// Go doesn't need a virtual environment like Python. Instead, the Go module system manages dependencies
	// Go versions: https://go.dev/dl/
	// Go download: cd Downloads; wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
	// Go extract: sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
	// Go add to PATH: export PATH=$PATH:/usr/local/go/bin

	/*
		var card string = "ace of spades" 	// formal declaration
		card := "ace of spades" 			// type inference declaration
	*/
	/*
		card := getOneCard()
		fmt.Printf("card = %s; card=%q\n", card, card) 	// string, double quoted str. snippet "fmpf"
		cardsWee := deck{"Two of Clubs", getOneCard()}  // deck is type slice of string
		cards := append(cardsWee, "Five of Hearts")     // argument cards_wee not changed
	*/
	/*
		// for loop
		for i := 0; i < len(cards); i += 1 {
		 	fmt.Println(cards[i]) // snippet "fp" or "fmpl"
		}
	*/
	/*
		// for loop range keyword
		for i, card := range cards { // range is a keyword in Go, returns arr/slice ind and copy of value
			fmt.Printf("i = %3d, card = %15s", i, card)
			fmt.Println("")
		}
	*/

	cards := deck.NewDeck() // new deck of cards

	fmt.Println("\nprint the whole deck") // receiver functions are LIKE "OOP" in Go
	cards.PrintWholeDeck()
	cards.ShuffleDeck() // shuffle the deck
	fmt.Println("\nprint the shuffled deck")
	cards.PrintWholeDeck()

	fmt.Println("\ndeal two hands of 5 cards")
	hand1, remainingCards := deck.Deal(cards, 5)
	fmt.Println("hand1:")
	hand1.PrintWholeDeck()
	fmt.Println("hand2:")
	hand2, _ := deck.Deal(remainingCards, 5)
	hand2.PrintWholeDeck()

	fmt.Println("\nbyte slice interlude")
	hiStr := "Hi there"
	bs := []byte(hiStr)
	fmt.Println(hiStr) // hi there
	fmt.Println(bs)    // [72 105 32 116 104 101 114 101]

	fmt.Println("\nhand 1 to single string")
	hand1SingleStr := deck.ToSingleString(hand1, ":")
	fmt.Println(hand1SingleStr)

	fmt.Println("\nhand 1 reconstructed to string slice")
	tempReconHand := deck.FromSingleString(hand1SingleStr, deck.MySeperator)
	tempReconHand.PrintWholeDeck()

	fmt.Println("\nwrite the whole deck to file")
	cards.WriteToFile("myfile.txt")

	fmt.Println("\nread the whole deck from file")
	cards = deck.ReadFromFile("myfile.txt")
	cards.PrintWholeDeck()

	fmt.Print("\nrandom number generator: ")
	fmt.Printf("rand.Intn(10) = %3d\n", rand.Intn(10)) // [0,n)

	fmt.Println("\n??Go similarity to Python??")
	var myList = []int{100, 101, 102, 103} // a "list" (slice)
	fmt.Println(len(myList))               // Go uses len function
	fmt.Println(myList[:3])                // slice notation is [start,end), start inclusive, end exclusive

	fmt.Println("\nOddEven")
	// calling Exported function OddEven() from package odd_even in module
	// bitbucket.org/weebucket/gobootcampagain/odd_even/reusable which is in another directory
	// (checking out packages / modules / go.mod and go.work, just seeing how it's done)
	mySlice1, mySlice2 := odd_even.OddEven()
	odd_even.PrintOddAndEven(mySlice1, mySlice2)
}
