package lus

func (b *Board) IsCorrect() bool {
	for i := 0; i < b.ySize; i++ {
		for j := 0; j < b.xSize; j++ {
			c := b.value[i][j]
			if c.value == Space && c.canPut {
				return false
			}
		}
	}
	return true
}
