package lus

import (
	"fmt"
	"testing"
)

func genSampleBoard() *Board {
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

func genCorrectBoard() *Board {
	s := [][]int{
		{Light, Block1, Space, Light, Block, Light},
		{Space, Block, Space, Space, Light, Space},
		{Space, Space, Light, Block2, Space, Block},
		{Block, Light, Block4, Light, Space, Space},
		{Space, Space, Light, Space, Block0, Space},
		{Light, Block, Space, Space, Block, Light},
	}

	b, _ := NewBoard(s)
	return b
}

func TestPrint(t *testing.T) {
	b := genSampleBoard()
	fmt.Printf("x: %d, y: %d\n", b.xSize, b.ySize)
	err := b.Print()
	if err != nil {
		t.Errorf("Print failed: %w", err)
	}
}
