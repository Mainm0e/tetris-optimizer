package main

import (
	"fmt"
	"tetris-optimizer/optimizer"
)

func main() {
	file := "example/example00.txt"
	returnVal := optimizer.Tetris(file)
	fmt.Println(returnVal)
}
