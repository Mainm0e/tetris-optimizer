package optimizer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestTetris(t *testing.T) {
	input := "../example2/example04.txt"
	expected := "Tetris"
	actual, err := Tetris(input)
	if err != nil {
		t.Errorf("Test failed with error: %v", err)
	}
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
func TestCheckFormat(t *testing.T) {
	input := "../example2/example01.txt"
	data, err := ioutil.ReadFile(input)
	expected := true
	actual, err := checkFormat(data)
	if actual != expected {

		t.Errorf("Test failed, expected: '%t', got:  '%t' , err: '%v'", expected, actual, err)
	}
}
func TestCutShape(t *testing.T) {
	input := "../example2/example00.txt"
	data, err := ioutil.ReadFile(input)
	if err != nil {
		t.Errorf("Failed to read input file: %v", err)
	}
	obj := bytes.Split(data, []byte{'\n', '\n'})
	lines := [][][]byte{}
	for _, v := range obj {
		lines = append(lines, bytes.Split(v, []byte{'\n'}))
	}
	expected := [][][]byte{
		{{35}, {35}, {35}, {35}},
		{{35, 35, 35, 35}},
		{{35, 35, 35}, {46, 46, 35}},
		{{46, 35, 35}, {35, 35, 46}},
	}

	expect := [][][]byte{}
	for i := 0; i < len(lines); i++ {
		value1, err := cutShape(lines[i])
		value := rollShape(value1)
		actual, _ := cutShape(value)
		value = rollShape(actual)
		if err != nil {
			t.Errorf("Failed to cut shape: %v", err)
		}
		expect = append(expect, value)
	}
	if !reflect.DeepEqual(expect, expected) {
		t.Errorf("Test failed, expected: '%v', got:  '%v' , err: '%v'", expected, expect, err)
	}
}

func TestShapeMap(t *testing.T) {
	input := [][][]string{{{"#"}, {"#"}, {"#"}, {"#"}}, {{"#", "#", "#", "#"}}, {{".", "#"}, {".", "#"}, {"#", "#"}}, {{"#", "#", "#"}, {".", ".", "#"}}, {{"#", "#"}, {"#", "."}, {"#", "."}}, {{"#", ".", "."}, {"#", "#", "#"}}, {{"#", "."}, {"#", "."}, {"#", "#"}}, {{"#", "#", "#"}, {"#", ".", "."}}, {{"#", "#"}, {".", "#"}, {".", "#"}}, {{".", ".", "#"}, {"#", "#", "#"}}, {{"#", "#"}, {"#", "#"}}, {{".", "#", "#"}, {"#", "#", "."}}, {{"#", "."}, {"#", "#"}, {".", "#"}}, {{"#", "#", "#"}, {".", "#", "."}}, {{"#", "."}, {"#", "#"}, {"#", "."}}, {{".", "#", "."}, {"#", "#", "#"}}, {{".", "#"}, {"#", "#"}, {".", "#"}}, {{"#", "#", "."}, {".", "#", "#"}}, {{".", "#"}, {"#", "#"}, {"#", "."}}}
	for i := 0; i < len(input); i++ {
		value := readShapeMap(input[i])
		fmt.Println(value)
	}
	input2 := [][]string{{"#"}, {"#"}, {"#"}, {"#"}}
	str := readShapeMap(input2)
	expected := "I0"
	if str != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, str)
	}
}
