package lus

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	b := genTestCase()

	bl, ans := Solve(*b)
	if bl {
		ans.Print()
	} else {
		fmt.Printf("couldn't solve")
	}
}