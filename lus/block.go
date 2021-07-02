package lus

import (
	"errors"
	"fmt"
)

type Block struct {
	value int
	canPut bool
}

func (b *Block) Print() error {
	switch b.value {
	case 0, 1, 2, 3, 4:
		// a block with number
		fmt.Print(b.value)
	case 5:
		// just a block
		fmt.Print("#")
	case 6:
		// space
		fmt.Print(" ")
	case 7:
		// light
		fmt.Print("*")
	default:
		return errors.New("invalid number of a block")
	}
	return nil
}