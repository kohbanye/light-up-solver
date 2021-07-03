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
		t.Error("couldn't solve")
	}
}

func TestIsCorrect(t *testing.T) {
	b := genCorrectBoard()

	b.Print()
	if !b.IsCorrect() {
		t.Errorf("correct board was determined to be wrong")
	}
}

func TestSetLight(t *testing.T) {
	b := genSampleBoard()

	b.SetLight(4, 2)
	b.Print()

	for i := 0; i < b.ySize; i++ {
		for j := 0; j < b.xSize; j++ {
			if b.value[i][j].isBright {
				fmt.Printf("(%d, %d) is lit\n", i, j)
			}
		}
	}
}
