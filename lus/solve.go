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
	for i := 1; ; i++ {
		update := false
		if col-i >= 0 {
			b.value[row][col-i].canPut = false
			b.value[row][col-i].isBright = true
			update = true
		}
		if col+i < b.xSize {
			b.value[row][col+i].canPut = false
			b.value[row][col+i].isBright = true
			update = true
		}
		if row-i >= 0 {
			b.value[row-i][col].canPut = false
			b.value[row-i][col].isBright = true
			update = true
		}
		if row+i < b.ySize {
			b.value[row+i][col].canPut = false
			b.value[row+i][col].isBright = true
			update = true
		}

		if !update {
			break
		}
	}

	return nil
}
