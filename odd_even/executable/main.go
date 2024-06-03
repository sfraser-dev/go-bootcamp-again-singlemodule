package main

import (
	"bitbucket.org/weebucket/gobootcampagain/odd_even/reusable"
)

func main() {
	mySlice1, mySlice2 := odd_even.OddEven()
	odd_even.PrintOddAndEven(mySlice1, mySlice2)
}