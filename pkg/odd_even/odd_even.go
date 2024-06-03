package odd_even

import (
	"fmt"
	"math/rand"
)

const (
	oddEvenSliceSize int = 10
)

func OddEven() (myOddEvenSlice1 []int, myOddEvenSlice2[]string) {
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
		fmt.Printf("%v is %v\n", mySlice1[i], mySlice2[i])
	}
}