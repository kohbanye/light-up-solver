package lus

import (
	"fmt"
	"testing"
)

func genSampleBoard() *Board {
	// source: https://www.nikoli.co.jp/ja/app/bj_tutorial/
	s := [][]int{
		{Space, Block1, Space, Space, Block, Space},
		{Space, Block, Space, Space, Space, Space},
		{Space, Space, Space, Block2, Space, Block},
		{Block, Space, Block4, Space, Space, Space},
		{Space, Space, Space, Space, Block0, Space},
		{Space, Block, Space, Space, Block, Space},
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

func genDifficultBoard() *Board {
	// source: https://web.archive.org/web/20181027164307/http://www.nikoli.com/ja/puzzles/bijutsukan/
	s := [][]int{
		{Space, Space, Block2, Space, Space, Space, Space, Space, Space, Space, Space, Space, Space, Space, Block, Space, Space, Space, Space, Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Space, Space, Space, Space},
		{Space, Space, Space, Block2, Space, Space, Space, Space, Space, Block3, Space, Space, Space, Space, Space, Block, Space, Space, Space, Space, Space, Block3, Space, Space, Space, Space, Space, Block, Space, Space, Space, Space, Space, Block1, Space, Space},
		{Block, Space, Space, Space, Block2, Space, Space, Space, Block3, Space, Space, Space, Block, Space, Space, Space, Block, Space, Space, Space, Block1, Space, Space, Space, Block2, Space, Space, Space, Block2, Space, Space, Space, Block1, Space, Space, Space},
		{Space, Space, Space, Space, Space, Space, Block1, Block1, Space, Space, Block, Space, Space, Space, Space, Space, Space, Space, Block1, Block1, Space, Space, Block1, Space, Space, Space, Space, Space, Space, Space, Block1, Block1, Space, Space, Block1, Space},
		{Space, Block, Space, Space, Block2, Block1, Space, Space, Space, Space, Space, Space, Space, Block, Space, Space, Block2, Block1, Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space, Block1, Block1, Space, Space, Space, Space, Space, Space},
		{Space, Space, Space, Block2, Space, Space, Space, Block, Space, Space, Space, Block0, Space, Space, Space, Block2, Space, Space, Space, Block, Space, Space, Space, Block, Space, Space, Space, Block1, Space, Space, Space, Block1, Space, Space, Space, Block},
		{Space, Space, Block1, Space, Space, Space, Space, Space, Block0, Space, Space, Space, Space, Space, Block3, Space, Space, Space, Space, Space, Block2, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Block1, Space, Space, Space},
		{Space, Block2, Space, Space, Space, Space, Space, Space, Space, Block, Space, Space, Space, Block3, Space, Space, Space, Space, Space, Space, Space, Block, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space},
		{Block2, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Block, Space, Block2, Space, Space, Space, Space, Space, Block0, Space, Space, Space, Block3, Space, Block1, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Block2, Space},
		{Block1, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Block, Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Space, Space},
		{Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Block1},
		{Space, Block, Space, Space, Space, Block, Space, Space, Space, Space, Space, Block1, Space, Block, Space, Space, Space, Block, Space, Space, Space, Space, Space, Block1, Space, Block3, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Block1},
		{Space, Space, Block1, Space, Space, Space, Space, Space, Space, Space, Block3, Space, Space, Space, Block, Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Block3, Space, Space, Space, Space, Space, Space, Space, Block1, Space},
		{Space, Space, Space, Block, Space, Space, Space, Space, Space, Block2, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Block1, Space, Space},
		{Block1, Space, Space, Space, Block2, Space, Space, Space, Block1, Space, Space, Space, Block2, Space, Space, Space, Block0, Space, Space, Space, Block1, Space, Space, Space, Block2, Space, Space, Space, Block, Space, Space, Space, Block3, Space, Space, Space},
		{Space, Space, Space, Space, Space, Space, Block2, Block1, Space, Space, Block, Space, Space, Space, Space, Space, Space, Space, Block1, Block1, Space, Space, Block, Space, Space, Space, Space, Space, Space, Space, Block1, Block2, Space, Space, Block, Space},
		{Space, Block1, Space, Space, Block2, Block1, Space, Space, Space, Space, Space, Space, Space, Block2, Space, Space, Block1, Block1, Space, Space, Space, Space, Space, Space, Space, Block0, Space, Space, Block2, Block1, Space, Space, Space, Space, Space, Space},
		{Space, Space, Space, Block1, Space, Space, Space, Block, Space, Space, Space, Block, Space, Space, Space, Block1, Space, Space, Space, Block1, Space, Space, Space, Block, Space, Space, Space, Block2, Space, Space, Space, Block, Space, Space, Space, Block},
		{Space, Space, Block2, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Block, Space, Space, Space, Space, Space, Block0, Space, Space, Space, Space, Space, Block1, Space, Space, Space},
		{Space, Space, Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Space, Space, Space, Space, Space, Space, Block1, Space, Space, Space, Space, Space, Space, Space, Space, Space, Space, Space, Block2, Space, Space},
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
