package lus

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	b := genSampleBoard()

	bl, ans := Solve(*b)
	if bl {
		ans.Print()
	} else {
		fmt.Println("couldn't solve")
	}
}

func TestIsCorrect(t *testing.T) {
	b := genCorrectBoard()

	b.Print()
	if !b.IsCorrect() {
		t.Errorf("correct board was determined to be wrong")
	}
}