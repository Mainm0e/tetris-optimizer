package optimizer

import (
	"io/ioutil"
	"testing"
)

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
	expected := []byte{35, 35, 35, 35}
	actual, err := cutShape(data)
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Test failed, expected: '%v', got:  '%v' , err: '%v'", expected, actual, err)
		}
	}
}
