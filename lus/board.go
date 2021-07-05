package lus

import (
	"errors"
	"fmt"
	"strings"
)

type Board struct {
	xSize int
	ySize int
	value [][]Cell
}

func NewBoard(v [][]int) (*Board, error) {
	b := new(Board)

	b.xSize = len(v[0])
	b.ySize = len(v)

	b.value = [][]Cell{}
	var lights [][2]int
	for i := 0; i < b.ySize; i++ {
		var s []Cell
		for j := 0; j < b.xSize; j++ {
			switch v[i][j] {
			case Block, Block0, Block1, Block2, Block3, Block4:
				s = append(s, Cell{v[i][j], false, false})
			case Space:
				s = append(s, Cell{v[i][j], true, false})
			case Light:
				s = append(s, Cell{Space, true, false})
				lights = append(lights, [2]int{i, j})
			default:
				return nil, errors.New("invalid number of a cell")
			}
		}
		b.value = append(b.value, s)
	}

	for _, coord := range lights {
		b.SetLight(coord[0], coord[1])
	}

	return b, nil
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

// CopyBoard copies elements from a source Board and returns a copied Board.
func CopyBoard(src Board) Board {
	dst := Board{src.xSize, src.ySize, make([][]Cell, len(src.value))}
	for i := range src.value {
		dst.value[i] = make([]Cell, len(src.value[i]))
		copy(dst.value[i], src.value[i])
	}

	return dst
}

type boolFunc func(c Cell) bool

func (b *Board) CountAroundCell(i, j int, f boolFunc) int {
	cnt := 0
	diff := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, d := range diff {
		if 0 <= i+d[0] && i+d[0] < b.ySize && 0 <= j+d[1] && j+d[1] < b.xSize &&
			f(b.value[i+d[0]][j+d[1]]) {
			cnt++
		}
	}

	return cnt
}

func (b *Board) SpaceAroundCell(i, j int) [][2]int {
	var s [][2]int
	diff := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, d := range diff {
		if 0 <= i+d[0] && i+d[0] < b.ySize && 0 <= j+d[1] && j+d[1] < b.xSize &&
			b.value[i+d[0]][j+d[1]].value == Space {
			s = append(s, [2]int{i + d[0], j + d[1]})
		}
	}

	return s
}

func (b *Board) ForbidAroundCell(i, j int) {
	for _, coord := range b.SpaceAroundCell(i, j) {
		b.value[coord[0]][coord[1]].canPut = false
	}
}

func (b *Board) LightAroundCell(i, j int) {
	for _, coord := range b.SpaceAroundCell(i, j) {
		b.SetLight(coord[0], coord[1])
	}
}
