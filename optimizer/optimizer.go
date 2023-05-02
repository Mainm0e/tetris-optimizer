package optimizer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

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
	moreCoffe, err := cutShape(data)
	fmt.Println(moreCoffe)
	return "Tetris", nil
}

// continue here
// cutShape takes a slice of bytes as an argument and returns a slice of bytes and an error.
// byte # = 35 and byte . = 46 and byte \n = 10

func cutShape(data []byte) ([]byte, error) {
	for i := 0; i < len(data); i++ {

	}
	return []byte{35, 35, 35, 35}, nil
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
