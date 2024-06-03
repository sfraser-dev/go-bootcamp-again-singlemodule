package odd_even

import "testing"

func Test_oddEven(t *testing.T) {
	mySlice1, mySlice2 := OddEven()

	if len(mySlice1) != oddEvenSliceSize {
		t.Errorf("length of slice (%v)containing random numbers is not equal to %v\n", len(mySlice1), oddEvenSliceSize)
	}

	if len(mySlice2) != oddEvenSliceSize {
		t.Errorf("length of slice (%v) containing 'odd' and 'even' strings is not equal to %v\n", len(mySlice2), oddEvenSliceSize)
	}
}
