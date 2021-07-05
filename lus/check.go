package lus

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

func (b *Board) IsWrong() bool {
	for i := 0; i < b.ySize; i++ {
		for j := 0; j < b.xSize; j++ {
			c := b.value[i][j]
			light := b.CountAroundCell(i, j, func(c Cell) bool { return c.value == Light })

			if light > c.value {
				return true
			}
		}
	}

	return false
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
