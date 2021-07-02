package lus

import (
	"errors"
	"fmt"
	"strings"
)

type Board struct {
	xSize int
	ySize int
	value [][]Cell
}

func NewBoard(v [][]int) (*Board, error) {
	b := new(Board)

	b.xSize = len(v[0])
	b.ySize = len(v)

	b.value = [][]Cell{}
	for _, arr := range v {
		var s []Cell
		for _, x := range arr {
			switch x {
			case Block, Block0, Block1, Block2, Block3, Block4, Light:
				s = append(s, Cell{x, false})
			case Space:
				s = append(s, Cell{x, true})
			default:
				return nil, errors.New("invalid number of a cell")
			}
		}
		b.value = append(b.value, s)
	}

	return b, nil
}

func (b *Board) Print() error {
	fmt.Println("┏" + strings.Repeat("━", b.xSize*2-1) + "┓")

	for _, arr := range b.value {
		fmt.Print("┃")
		for i, block := range arr {
			err := block.Print()
			if err != nil {
				return fmt.Errorf("Print failed: %w", err)
			}
			if i < b.xSize-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println("┃")
	}

	fmt.Println("┗" + strings.Repeat("━", b.xSize*2-1) + "┛")

	return nil
}
