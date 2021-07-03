package lus

import (
	"errors"
)

func Solve(b Board) (bool, Board) {
	if b.IsCorrect() {
		return true, b
	} else {
		for i := 0; i < b.ySize; i++ {
			for j := 0; j < b.xSize; j++ {
				if b.value[i][j].canPut {
					b_ := CopyBoard(b)
					b.SetLight(i, j)
					if bl, ans := Solve(b); bl {
						return true, ans
					}

					b = CopyBoard(b_)
					b.value[i][j].canPut = false
					if bl, ans := Solve(b); bl {
						return true, ans
					}

					return false, b
				}
			}
		}
		return false, b
	}
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
	c.isBright = true
	// light the left side
	for i := 1; col-i >= 0 && !b.value[row][col-i].isBlock(); i++ {
		b.value[row][col-i].canPut = false
		b.value[row][col-i].isBright = true
	}
	// light the right side
	for i := 1; col+i < b.xSize && !b.value[row][col+i].isBlock(); i++ {
		b.value[row][col+i].canPut = false
		b.value[row][col+i].isBright = true
	}
	// light the upper side
	for i := 1; row-i >= 0 && !b.value[row-i][col].isBlock(); i++ {
		b.value[row-i][col].canPut = false
		b.value[row-i][col].isBright = true
	}
	// light the under side
	for i := 1; row+i < b.ySize && !b.value[row+i][col].isBlock(); i++ {
		b.value[row+i][col].canPut = false
		b.value[row+i][col].isBright = true
	}

	return nil
}
