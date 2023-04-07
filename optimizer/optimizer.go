package optimizer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// fineTetromino finds the tetromino in the piece and returns the index of the tetromino
// L = 0, J = 1, T = 2, Z = 3, S = 4, O = 5, I = 6
func identifyTetromino(s string) (string, error) {
	// Split the string into individual rows
	rows := strings.Split(s, "\n")

	// Trim any empty rows or trailing spaces
	for i := 0; i < len(rows); i++ {
		rows[i] = strings.TrimSpace(rows[i])
	}

	fmt.Println("rows", rows)
	// Check each tetromino shape
	if len(rows) == 3 || rows[0] == "####" || rows[1] == "####" || rows[2] == "####" {
		return "I", nil
	} else if rows[0] == "##" && rows[1] == "##" {
		return "O", nil
	} else {

		for i := 0; i < len(rows); i++ {
			if len(rows[i]) == 2 {
				if len(rows[i+1]) == 3 && len(rows[i-1]) == 3 {
					return "Z", nil
				}
			}

		}
		return "", fmt.Errorf("Tetromino not found")

	}
}

// Tetris is the main function for the optimizer
func Tetris(example string) string {

	file, err := os.Open(example)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// lIndex is the line index for counting the lines
	var lIndex int
	var piece string
	var tabel []string

	// Read the file line by line
	for scanner.Scan() {
		lIndex++
		line := scanner.Text()
		piece += line + "\n"

		if lIndex == 3 {
			tetromino, err := identifyTetromino(piece)
			if err != nil {
				fmt.Println("Error finding tetromino:", err)
				return ""
			} else {
				tabel = append(tabel, tetromino)
				piece = ""
				lIndex = 0
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}

	return "fileValue"
}
