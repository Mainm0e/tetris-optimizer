package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatalf("failed opening file : %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	for _, eachline := range txtlines {
		fmt.Printf(eachline)
	}
}
