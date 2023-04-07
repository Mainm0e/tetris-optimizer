package optimizer

import (
	"testing"
)

func TestTetris(t *testing.T) {
	value := Tetris("../example/example00.txt")
	want := "####\n...#\n....\n....\n"

	if value != want {
		t.Errorf("Tetris() = %q, want %q", value, want)
	}
}

func TestIdentifyTetromino(t *testing.T) {
	value, _ := identifyTetromino("##..\n.##.\n....\n....\n")
	want := "Z"

	if value != want {
		t.Errorf("identifyTetromino() = %q, want %q", value, want)
	}
}
