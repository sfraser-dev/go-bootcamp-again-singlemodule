package my_structs

import (
	"fmt"
)

type ContactInfo struct {
	Mob  int
	Addr string
}
type PersonType1 struct {
	FirstName string
	LastName  string
}

type PersonType2 struct {
	FirstName string
	LastName  string
	Info      ContactInfo
}

type PersonType3 struct {
	FirstName   string
	LastName    string
	ContactInfo // equivalent to ContactInfo ContactInfo
}

func (p PersonType1) Speak() {
	fmt.Printf("p = %+v\n", p) // fieldnames printed too with %+v
}

func (p PersonType2) Speak() {
	fmt.Printf("p = %+v\n", p) // fieldnames printed too with %+v
}

func (p PersonType3) Speak() {
	fmt.Printf("p = %+v\n", p) // fieldnames printed too with %+v
}

// POINTER RECEIVER
func (ptr *PersonType1) ChangeName(nameIn string) {
	// I'm passing a pointer into this test func
	(*ptr).FirstName = nameIn // can dereference the pointer myself
	ptr.LastName = nameIn     // can let Go auto dereference the pointer
}

// POINTER RECEIVER
func (ptr *PersonType2) ChangeName(nameIn string) {
	// I'm passing a struct by value to this test func
	// Go knows it's a ptr-rec-func so will auto fix and use reference the struct instead (use its address)
	(*ptr).FirstName = nameIn // can dereference the pointer myself
	ptr.LastName = nameIn     // can let Go auto dereference the pointer

}

// POINTER RECEIVER (doing same as with PersonType2's ChangeName func above)
func (ptr *PersonType3) ChangeName(nameIn string) {
	// I'm passing a struct by value to this test func
	// Go knows it's a ptr-rec-func so will auto fix and use reference the struct instead (use its address)
	(*ptr).FirstName = nameIn // can dereference the pointer myself
	ptr.LastName = nameIn // can let Go auto dereference the pointer
}
