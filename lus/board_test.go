package lus

import "testing"

var sample = Board{
	xSize: 10,
	ySize: 10,
	// source: https://www.nikoli.co.jp/ja/app/bj_tutorial/
	value: [][]Block{
		{{6, true}, {1, false}, {6, true}, {6, true}, {5, false}, {6, true}},
		{{6, true}, {5, false}, {6, true}, {6, true}, {6, true}, {6, true}},
		{{6, true}, {6, true}, {6, true}, {2, false}, {6, true}, {5, false}},
		{{5, false}, {6, true}, {4, false}, {6, true}, {6, true}, {6, true}},
		{{6, true}, {6, true}, {6, true}, {6, true}, {0, false}, {6, true}},
		{{6, true}, {5, false}, {6, true}, {6, true}, {5, true}, {6, true}},
	},
}

func TestPrint(t *testing.T) {
	err := sample.Print()
	if err != nil {
		t.Errorf("Print failed: %w", err)
	}
}
