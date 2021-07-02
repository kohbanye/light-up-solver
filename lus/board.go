package lus

import (
	"fmt"
	"strings"
)

type Board struct {
	xSize int
	ySize int
	value [][]Cell
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
