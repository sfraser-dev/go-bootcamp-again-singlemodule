package pkg_slice_fill_rand_odd_even

import (
	"fmt"
	"math/rand"
)

const (
	oddEvenSliceSize int = 10
)

// fill a slice with random values (between 0 and 99 inclusive)
// report if each random value is odd or even
func OddEven() (myOddEvenSlice1 []int, myOddEvenSlice2 []string) {
	for i := 0; i < oddEvenSliceSize; i += 1 {
		myOddEvenSlice1 = append(myOddEvenSlice1, rand.Intn(100))
	}
	for i := 0; i < len(myOddEvenSlice1); i += 1 {
		if myOddEvenSlice1[i]%2 == 0 {
			myOddEvenSlice2 = append(myOddEvenSlice2, "Even")
		} else {
			myOddEvenSlice2 = append(myOddEvenSlice2, "Odd")
		}
	}
	return myOddEvenSlice1, myOddEvenSlice2
}

func PrintOddAndEven(mySlice1 []int, mySlice2 []string) {
	for i := 0; i < oddEvenSliceSize; i += 1 {
		fmt.Printf("%03v is %v\n", mySlice1[i], mySlice2[i])
	}
}
