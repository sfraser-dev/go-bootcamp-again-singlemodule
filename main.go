package main

import (
	"fmt"
	"math/rand"

	// A MODULE CONTAINS MULTIPLE PACAKGES (in diff sub-dirs)
	// ONLY ONE NAMED PACKAGE ALLOWED PER FOLDER
	// PACKAGE is a COLLECTION OF SOURCE FILES (all in same dir and all with same pkg name)
	//
	// Import package "pkg_deck_of_cards" AS "pkg_deck"
	// from module gobootcampagain_singlemodule's folder ./pkg/deck_of_cards
	pkg_deck_new_name "go-bootcamp-again-singlemodule/pkg/deck_of_cards"
	// Import package "pkg_interfaces_read_file AS "pkg_interfaces_read_file"
	// from module gobootcampagain_singlemodule's folder ./pkg/interfaces_shapes
	pkg_interfaces_read_file "go-bootcamp-again-singlemodule/pkg/interfaces_read_file"
	// Import package "pkg_slice_fill_rand_odd_even" AS "pkg_odd_even"
	// from module gobootcampagain_singlemodule's folder ./pkg/slice_fill_rand_odd_even
	pkg_odd_even "go-bootcamp-again-singlemodule/pkg/slice_fill_rand_odd_even"
	// Import package "pkg_structs_and_pointers" AS "pkg_structs_and_pointers"
	// from module gobootcampagain_singlemodule's folder ./pkg/structs_and_pointers
	pkg_structs_and_pointers "go-bootcamp-again-singlemodule/pkg/structs_and_pointers"
	// Import "package pkg_maps"
	// from module gobootcampagain_singlemodule's folder ./pkg/maps
	pkg_maps "go-bootcamp-again-singlemodule/pkg/maps"
	// Import "package pkg_interfaces_http_get"
	// from module gobootcampagain_singlemodule's folder ./pkg/interfaces_http_get
	pkg_interfaces_http_get "go-bootcamp-again-singlemodule/pkg/interfaces_http_get"
	// Import "package pkg_interfaces"
	// from module gobootcampagain_singlemodule's folder ./pkg/interfaces
	pkg_interfaces "go-bootcamp-again-singlemodule/pkg/interfaces"
	// Import "package pkg_shapes" AS "package pkg_shapes"
	// from module gobootcampagain_singlemodule's folder ./pkg/interfaces_shapes
	pkg_interfaces_shapes "go-bootcamp-again-singlemodule/pkg/interfaces_shapes"
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

	cards := pkg_deck_new_name.NewDeck() // new deck of cards

	fmt.Println("\nprint the whole deck") // receiver functions are LIKE "OOP" in Go
	cards.PrintWholeDeck()
	cards.ShuffleDeck() // shuffle the deck
	fmt.Println("\nprint the shuffled deck")
	cards.PrintWholeDeck()

	fmt.Println("\ndeal two hands of 5 cards")
	hand1, remainingCards := pkg_deck_new_name.Deal(cards, 5)
	fmt.Println("hand1:")
	hand1.PrintWholeDeck()
	fmt.Println("hand2:")
	hand2, _ := pkg_deck_new_name.Deal(remainingCards, 5)
	hand2.PrintWholeDeck()

	fmt.Println("\nbyte slice interlude")
	hiStr := "Hi there"
	bs := []byte(hiStr)
	fmt.Println(hiStr) // hi there
	fmt.Println(bs)    // [72 105 32 116 104 101 114 101]

	fmt.Println("\nhand 1 to single string")
	hand1SingleStr := pkg_deck_new_name.ToSingleString(hand1, ":")
	fmt.Println(hand1SingleStr)

	fmt.Println("\nhand 1 reconstructed to string slice")
	tempReconHand := pkg_deck_new_name.FromSingleString(hand1SingleStr, pkg_deck_new_name.MySeperator)
	tempReconHand.PrintWholeDeck()

	fmt.Println("\nwrite the whole deck to file")
	cards.WriteToFile("_deck_of_cards.log")

	fmt.Println("\nread the whole deck from file")
	cards = pkg_deck_new_name.ReadFromFile("_deck_of_cards.log")
	cards.PrintWholeDeck()

	fmt.Print("\nrandom number generator: ")
	fmt.Printf("rand.Intn(10) = %3d\n", rand.Intn(10)) // [0,n)

	fmt.Println("\n??Go similarity to Python??")
	var myList = []int{100, 101, 102, 103} // a "list" (slice)
	fmt.Println(len(myList))               // Go uses len function
	fmt.Println(myList[:3])                // slice notation is [start,end), start inclusive, end exclusive

	fmt.Println("\nOddEven")
	// calling Exported function OddEven() from package pkg_odd_even in module
	// bitbucket.org/weebucket/gobootcampagain/odd_even/reusable which is in another directory
	// (checking out packages / modules / go.mod and go.work, just seeing how it's done)
	mySlice1, mySlice2 := pkg_odd_even.OddEven()
	pkg_odd_even.PrintOddAndEven(mySlice1, mySlice2)

	fmt.Println("\nStructs")
	// var p1 pkg_structs_and_pointers.Person // Go auto assigns zero "" values
	// p1.FirstName = "Joe"
	// p1.LastName = "Bloggs"
	p1 := pkg_structs_and_pointers.PersonType1{FirstName: "joe", LastName: "bloggs"}
	fmt.Printf("p1 = %+v\n", p1) // fieldnames printed too with %+v
	fmt.Println("...Embedded struct...")
	p2 := pkg_structs_and_pointers.PersonType2{FirstName: "sam", LastName: "gee", Info: pkg_structs_and_pointers.ContactInfo{Mob: 123, Addr: "ABC"}}
	fmt.Printf("p2 = %+v\n", p2)
	fmt.Printf("p2.Info.Mob = %+v\n", p2.Info.Mob)
	fmt.Printf("p2.Info.Addr= %+v\n", p2.Info.Addr)
	fmt.Println("...NO FIELDNAME embedded struct...") // helps with code reuse
	p3 := pkg_structs_and_pointers.PersonType3{FirstName: "kay", LastName: "nab", ContactInfo: pkg_structs_and_pointers.ContactInfo{Mob: 123, Addr: "ABC"}}
	fmt.Printf("p3 = %+v\n", p3)
	fmt.Printf("p3.Mob = %+v\n", p3.Mob)  // helping with code reuse, INHERITANCE LIKE BEHAVOUR
	fmt.Printf("p3.Addr= %+v\n", p3.Addr) // helping with code reuse, INHERITANCE LIKE BEHAVOUR
	fmt.Println("...Print with receiver functions...")
	p1.Speak()
	p2.Speak()
	p3.Speak()

	fmt.Println("\nPointers")
	// note: dereferencing is accessing the value of the variable stored at the memory location in the ptr
	// note: address-referencing is accessing the address of a variable rather than its value
	// note: ptr-rec-func is a method with a pointer receiver
	var X int = 5
	// var ptr *int; ptr = &X
	ptr := &X
	fmt.Printf("X = %+v\n", X)
	(*ptr) += 1
	fmt.Printf("X = %+v\n", X)
	fmt.Println("...Pointers to Structs...")
	ptrP1 := &p1
	// Passing pointer (holding addr of variable) used to call the ptr-rec-func ChangeName (FINE)
	// "C/C++ sytle" POINTER passing, pointer holding addr of variable passed, just as C/C++ does it
	ptrP1.ChangeName("Terminator")
	ptrP1.Speak()
	ptrP2 := &p2
	// I'm dereferencing the pointer ptrP2 here so actually passing p2 by value to ptr-rec-func ChangeName (FINE)
	// Syntactic sugar and address-referencing
	// Passing struct by value, but Go knows it's a ptr-rec-func so will pass address-reference struct p3 instead
	(*ptrP2).ChangeName("Godzilla") // passing struct p2 by value to ptr-rec-func, address-referenced, not copied
	(*ptrP2).Speak()
	// I'm straight-up passing p3 by value to ptr-rec-func ChangeName (FINE)
	// Syntactic sugar and address-referencing
	// Passing struct by value, but Go knows it's a ptr-rec-func so will pass address-reference struct p3 instead
	p3.ChangeName("Spiderman") // passing struct p3 by value to ptr-rec-func, address-referenced, not copied
	p3.Speak()

	fmt.Println("\nMaps")
	pkg_maps.CreateMaps()

	fmt.Println("\nInterfaces")
	pkg_interfaces.RunMe()

	fmt.Println("\nInterfaces Read Body HTML from Google Homepage")
	pkg_interfaces_http_get.GettingGoogleHomepageHTML()

	fmt.Println("\n\nInterface assignment with shapes")
	pkg_interfaces_shapes.RunIt()

	fmt.Println("\nInterfaces read text file")
	pkg_interfaces_read_file.RunIt()
}
