package optimizer

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func byteToString(data [][][]byte) [][][]string {
	result := make([][][]string, len(data))
	for i := range data {
		result[i] = make([][]string, len(data[i]))
		for j := range data[i] {
			result[i][j] = make([]string, len(data[i][j]))
			for k := range data[i][j] {
				result[i][j][k] = string(data[i][j][k])
			}
		}
	}
	return result
}

// Tetris is the main function of the program. It takes a file path as an argument and returns a string.
func Tetris(example string) (string, error) {
	data, err := ioutil.ReadFile(example)
	if err != nil {
		panic(err)
	}
	coffe, err := checkFormat(data)
	if err != nil {
		return "checkFormat():", err
	}
	if coffe == false {
		return "Invalid file format", nil
	}
	obj := bytes.Split(data, []byte{'\n', '\n'})
	lines := [][][]byte{}
	for _, v := range obj {
		lines = append(lines, bytes.Split(v, []byte{'\n'}))
	}

	expect := [][][]byte{}
	for i := 0; i < len(lines); i++ {
		value1, err := cutShape(lines[i])
		value := rollShape(value1)
		actual, _ := cutShape(value)
		value = rollShape(actual)
		if err != nil {
			fmt.Println("Failed to cut shape:", err)
		}
		expect = append(expect, value)
	}
	stringMap := byteToString(expect)
	tetrisShape := []string{}
	for i := 0; i < len(stringMap); i++ {
		vShape := readShapeMap(stringMap[i])
		tetrisShape = append(tetrisShape, vShape)
	}
	for i := 0; i < len(tetrisShape); i++ {
		if tetrisShape[i] == "Unknown" {
			return "Unknown shape", errors.New("Unknown shape")
		}
	}
	/* 	fmt.Println(tetrisShape) */
	return "Tetris", nil
}

// continue here
// cutShape takes a slice of bytes as an argument and returns a slice of bytes and an error.
// byte # = 35 and byte . = 46 and byte \n = 10

func cutShape(lines [][]byte) ([][]byte, error) {
	var shape [][]byte
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 35 {
				shape = append(shape, lines[i])
				break
			}
		}
	}
	return shape, nil
}

func rollShape(shape [][]byte) [][]byte {
	matrix := shape
	// Find the dimensions of the matrix
	rows := len(matrix)
	cols := len(matrix[0])
	// Initialize the transposed matrix
	transposed := make([][]byte, cols)
	for i := range transposed {
		transposed[i] = make([]byte, rows)
	}
	// Transpose the matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}

// checkFormat checks the format of the input file. It takes a string as an argument and returns a boolean.
func checkFormat(data []byte) (bool, error) {

	str := string(data)
	newLine := strings.Split(str, "\n")
	// slice line by line and check length of each line
	if len(newLine[0]) == 4 {
		for j := 0; j < len(newLine); j += 5 {
			for i := 0; i < 4; i++ {
				if len(newLine[j+i]) != len(newLine[0]) {
					err := fmt.Sprintln("line: ", j+i, "is not valid", newLine[j+i])
					return false, errors.New(err)
				}
			}
		}
	}

	return true, nil
}

func readShapeMap(data [][]string) string {
	// Check for I0 shape
	if len(data) == 4 && len(data[0]) == 1 && len(data[1]) == 1 && len(data[2]) == 1 && len(data[3]) == 1 {
		if data[0][0] == "#" && data[1][0] == "#" && data[2][0] == "#" && data[3][0] == "#" {
			return "I0"
		}
	}
	// Check for I1 shape
	if len(data) == 1 && len(data[0]) == 4 {
		if data[0][0] == "#" && data[0][1] == "#" && data[0][2] == "#" && data[0][3] == "#" {
			return "I1"
		}
	}

	// Check for J0 {{".","#"}, {".","#"}, {"#", "#"}}
	if len(data) == 3 && len(data[0]) == 2 && len(data[1]) == 2 && len(data[2]) == 2 {
		if data[0][1] == "#" && data[1][1] == "#" && data[2][0] == "#" && data[2][1] == "#" {
			return "J0"
		}
	}
	// Check for J1 {{"#", "#", "#"},{".", ".", "#"}}
	if len(data) == 2 && len(data[0]) == 3 && len(data[1]) == 3 {
		if data[0][0] == "#" && data[0][1] == "#" && data[0][2] == "#" && data[1][2] == "#" {
			return "J1"
		}
	}
	// Check for J2 {{"#", "#"}, {"#", "."}, {"#", "."}}
	if len(data) == 3 && len(data[0]) == 2 && len(data[1]) == 2 && len(data[2]) == 2 {
		if data[0][0] == "#" && data[0][1] == "#" && data[1][0] == "#" && data[2][0] == "#" {
			return "J2"
		}
	}

	// Check for J3 {{"#", ".", "."}, {"#", "#", "#"}}
	if len(data) == 2 && len(data[0]) == 3 && len(data[1]) == 3 {
		if data[0][0] == "#" && data[1][0] == "#" && data[1][1] == "#" && data[1][2] == "#" {
			return "J3"
		}
	}
	// Check for L0 {{"#", "."}, {"#", "."}, {"#", "#"}}
	if len(data) == 3 && len(data[0]) == 2 && len(data[1]) == 2 && len(data[2]) == 2 {
		if data[0][0] == "#" && data[1][0] == "#" && data[2][0] == "#" && data[2][1] == "#" {
			return "L0"
		}
	}
	// Check for L1 {{"#", "#", "#"}, {"#", ".", "."}}
	if len(data) == 2 && len(data[0]) == 3 && len(data[1]) == 3 {
		if data[0][0] == "#" && data[0][1] == "#" && data[0][2] == "#" && data[1][0] == "#" {
			return "L1"
		}
	}

	// Check for L2 {{"#", "#"}, {".", "#"}, {".", "#"}}
	if len(data) == 3 && len(data[0]) == 2 && len(data[1]) == 2 && len(data[2]) == 2 {
		if data[0][0] == "#" && data[0][1] == "#" && data[1][1] == "#" && data[2][1] == "#" {
			return "L2"
		}
	}
	// Check for L3 {{".", ".", "#"}, {"#", "#", "#"}}
	if len(data) == 2 && len(data[0]) == 3 && len(data[1]) == 3 {
		if data[0][2] == "#" && data[1][0] == "#" && data[1][1] == "#" && data[1][2] == "#" {
			return "L3"
		}
	}
	// Check for O0 shape
	if len(data) == 2 && len(data[0]) == 2 && len(data[1]) == 2 {
		if data[0][0] == "#" && data[0][1] == "#" && data[1][0] == "#" && data[1][1] == "#" {
			return "O0"
		}
	}

	// Check for S0 {{".","#","#"}, {"#", "#", "."}}
	if len(data) == 2 && len(data[0]) == 3 && len(data[1]) == 3 {
		if data[0][1] == "#" && data[0][2] == "#" && data[1][0] == "#" && data[1][1] == "#" {
			return "S0"
		}
	}
	// Check for S1 {{"#", "."}, {"#", "#"}, {".", "#"}}
	if len(data) == 3 && len(data[0]) == 2 && len(data[1]) == 2 && len(data[2]) == 2 {
		if data[0][0] == "#" && data[1][0] == "#" && data[1][1] == "#" && data[2][1] == "#" {
			return "S1"
		}
	}
	// Check for T0 {{"#", "#", "#"},{".","#", "."}}
	if len(data) == 2 && len(data[0]) == 3 && len(data[1]) == 3 {
		if data[0][0] == "#" && data[0][1] == "#" && data[0][2] == "#" && data[1][1] == "#" {
			return "T0"
		}
	}
	// Check for T1 {{"#", "."}, {"#", "#"}, {"#", "."}}
	if len(data) == 3 && len(data[0]) == 2 && len(data[1]) == 2 && len(data[2]) == 2 {
		if data[0][0] == "#" && data[1][0] == "#" && data[1][1] == "#" && data[2][0] == "#" {
			return "T1"
		}
	}
	// Check for T2 {{".","#", "."}, {"#", "#", "#"}}
	if len(data) == 2 && len(data[0]) == 3 && len(data[1]) == 3 {
		if data[0][1] == "#" && data[1][0] == "#" && data[1][1] == "#" && data[1][2] == "#" {
			return "T2"
		}
	}
	// Check for T3 {{".", "#"}, {"#", "#"}, {".", "#"}}
	if len(data) == 3 && len(data[0]) == 2 && len(data[1]) == 2 && len(data[2]) == 2 {
		if data[0][1] == "#" && data[1][0] == "#" && data[1][1] == "#" && data[2][1] == "#" {
			return "T3"
		}
	}
	// Check for Z0 {{"#", "#", "."}, {".","#", "#"}}
	if len(data) == 2 && len(data[0]) == 3 && len(data[1]) == 3 {
		if data[0][0] == "#" && data[0][1] == "#" && data[1][1] == "#" && data[1][2] == "#" {
			return "Z0"
		}
	}
	// Check for Z1 {{".", "#"}, {"#", "#"}, {"#", "."}}
	if len(data) == 3 && len(data[0]) == 2 && len(data[1]) == 2 && len(data[2]) == 2 {
		if data[0][1] == "#" && data[1][0] == "#" && data[1][1] == "#" && data[2][0] == "#" {
			return "Z1"
		}
	}
	return "Unknown"
}
