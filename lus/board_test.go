package lus

import (
	"fmt"
	"testing"
)

func genTestCase() *Board {
	// source: https://www.nikoli.co.jp/ja/app/bj_tutorial/
	s := [][]int{
		{6, 1, 6, 6, 5, 6},
		{6, 5, 6, 6, 6, 6},
		{6, 6, 6, 2, 6, 5},
		{5, 6, 4, 6, 6, 6},
		{6, 6, 6, 6, 0, 6},
		{6, 5, 6, 6, 5, 6},
	}

	b, _ := NewBoard(s)
	return b
}

func TestPrint(t *testing.T) {
	b := genTestCase()
	fmt.Printf("x: %d, y: %d\n", b.xSize, b.ySize)
	err := b.Print()
	if err != nil {
		t.Errorf("Print failed: %w", err)
	}
}
