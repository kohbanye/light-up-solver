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
	var lights [][2]int
	for i := 0; i < b.ySize; i++ {
		var s []Cell
		for j := 0; j < b.xSize; j++ {
			switch v[i][j] {
			case Block, Block0, Block1, Block2, Block3, Block4:
				s = append(s, Cell{v[i][j], false, false})
			case Space:
				s = append(s, Cell{v[i][j], true, false})
			case Light:
				s = append(s, Cell{v[i][j], true, false})
				lights = append(lights, [2]int{i, j})
			default:
				return nil, errors.New("invalid number of a cell")
			}
		}
		b.value = append(b.value, s)
	}

	for _, coord := range lights {
		b.SetLight(coord[0], coord[1])
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
