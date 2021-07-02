package lus

import (
	"errors"
	"fmt"
)

const (
	Block0 = iota
	Block1
	Block2
	Block3
	Block4
	Block
	Space
	Light
)

type Cell struct {
	value    int
	canPut   bool
	isBright bool
}

func (c *Cell) Print() error {
	switch c.value {
	case Block0, Block1, Block2, Block3, Block4:
		fmt.Print(c.value)
	case Block:
		fmt.Print("#")
	case Space:
		fmt.Print(" ")
	case Light:
		fmt.Print("*")
	default:
		return errors.New("invalid number of a block")
	}
	return nil
}
