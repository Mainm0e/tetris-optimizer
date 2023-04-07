package optimizer

import (
	"bufio"
	"fmt"
	"os"
)

func Tetris(example string) string {

	file, err := os.Open(example)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileValue string
	for scanner.Scan() {
		line := scanner.Text()
		fileValue += line + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	fmt.Println(fileValue)
	return fileValue
}
