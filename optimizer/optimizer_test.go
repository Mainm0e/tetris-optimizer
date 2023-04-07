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
