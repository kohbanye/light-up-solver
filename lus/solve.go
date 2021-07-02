package lus

import "errors"

func (b *Board) IsCorrect() bool {
	for i := 0; i < b.ySize; i++ {
		for j := 0; j < b.xSize; j++ {
			c := b.value[i][j]
			if c.value == Space && !c.isBright {
				return false
			}
		}
	}
	return true
}

func (b *Board) SetLight(row int, col int) error {
	c := &b.value[row][col]
	if c.value != Space {
		return errors.New("cannot set light(cell is not a space)")
	}
	if c.isBright {
		return errors.New("cannot set light(cell is already bright)")
	}

	c.value = Light
	c.canPut = false
	// light the upper side
	for i := 1; ; i++ {
		if col-i < 0 || b.value[row][col-i].isBlock() {
			break
		}
		b.value[row][col-i].canPut = false
		b.value[row][col-i].isBright = true
	}
	// light the under side
	for i := 1; ; i++ {
		if col+i >= b.xSize || b.value[row][col+i].isBlock() {
			break
		}
		b.value[row][col+i].canPut = false
		b.value[row][col+i].isBright = true
	}
	// light the left side
	for i := 1; ; i++ {
		if row-i < 0 || b.value[row-i][col].isBlock() {
			break
		}
		b.value[row-i][col].canPut = false
		b.value[row-i][col].isBright = true
	}
	// light the right side
	for i := 1; ; i++ {
		if row+i >= b.ySize || b.value[row+i][col].isBlock() {
			break
		}
		b.value[row+i][col].canPut = false
		b.value[row+i][col].isBright = true
	}

	return nil
}

func (c *Cell) isBlock() bool {
	switch c.value {
	case Block, Block0, Block1, Block2, Block3, Block4:
		return true
	default:
		return false
	}
}
