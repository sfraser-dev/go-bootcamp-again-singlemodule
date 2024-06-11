package pkg_structs_and_pointers

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

// ptr-rec-funcs can take the pointer (as declared)
// **OR** they can use a SHORTCUT
// and pass the value/obj by value that the pointer points to
//
// It's auto-fixing as it was expecting a pointer but got a value/obj
// It then auto-references, grabs address of the value/obj, and uses this
//
// Dereferencing is accessing the value of the variable stored at the memory location in the ptr
// Address-referencing is accessing the address of a variable rather than its value

// POINTER RECEIVER (OPTION POINTER)
func (ptr *PersonType1) ChangeName(nameIn string) {
	// I'm passing a pointer into this test func
	(*ptr).FirstName = nameIn // can dereference the pointer myself
	ptr.LastName = nameIn     // can let Go auto dereference the pointer
}

// POINTER RECEIVER (OPTION VALUE)
func (ptr *PersonType2) ChangeName(nameIn string) {
	// I'm passing a struct by value to this test func
	// Go knows it's a ptr-rec-func so will auto-fix and address-reference the struct instead
	(*ptr).FirstName = nameIn // can dereference the pointer myself
	ptr.LastName = nameIn     // can let Go auto dereference the pointer automatically

}

// POINTER RECEIVER (OPTION VALUE) (doing same as with PersonType2's ChangeName func above)
func (ptr *PersonType3) ChangeName(nameIn string) {
	// I'm passing a struct by value to this test func
	// Go knows it's a ptr-rec-func so will auto-fix and address-reference the struct instead
	(*ptr).FirstName = nameIn // can dereference the pointer myself
	ptr.LastName = nameIn     // can let Go auto dereference the pointer
}
