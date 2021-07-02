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
				c := b.value[i][j]
				if c.value == Space && c.canPut {
					b_ := b
					b.SetLight(i, j)

					if bl, ans := Solve(b); bl {
						return true, ans
					}

					b = b_
					b.value[i][j].canPut = false
					if bl, ans := Solve(b); bl {
						return true, ans
					}
				}
			}
		}
		return false, b
	}
}

func (b *Board) IsCorrect() bool {
	for i := 0; i < b.ySize; i++ {
		for j := 0; j < b.xSize; j++ {
			c := b.value[i][j]

			// check if all space cells are bright
			if c.value == Space && !c.isBright {
				return false
			}

			// check if numbers of lights around number-blocks are correct
			diff := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			if c.isNumberBlock() {
				cnt := 0
				for _, d := range diff {
					if 0 <= i+d[0] && i+d[0] < b.ySize && 0 <= j+d[1] && j+d[1] < b.xSize &&
						b.value[i+d[0]][j+d[1]].value == Light {
						cnt++
					}
				}

				if cnt != c.value {
					return false
				}
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
	// light the left side
	for i := 1; ; i++ {
		if col-i < 0 || b.value[row][col-i].isBlock() {
			break
		}
		b.value[row][col-i].canPut = false
		b.value[row][col-i].isBright = true
	}
	// light the right side
	for i := 1; ; i++ {
		if col+i >= b.xSize || b.value[row][col+i].isBlock() {
			break
		}
		b.value[row][col+i].canPut = false
		b.value[row][col+i].isBright = true
	}
	// light the upper side
	for i := 1; ; i++ {
		if row-i < 0 || b.value[row-i][col].isBlock() {
			break
		}
		b.value[row-i][col].canPut = false
		b.value[row-i][col].isBright = true
	}
	// light the under side
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

func (c *Cell) isNumberBlock() bool {
	switch c.value {
	case Block0, Block1, Block2, Block3, Block4:
		return true
	default:
		return false
	}
}
