package lus

import (
	"fmt"
	"strings"
)

type Board struct {
	xSize int
	ySize int
	value [][]Block
}

func (b *Board) Print() error {
	fmt.Println("┏", strings.Repeat("━", b.xSize), "┓")

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

	fmt.Println("┗", strings.Repeat("━", b.xSize), "┛")

	return nil
}
