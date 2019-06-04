package filesystem

import (
	"fmt"
	"testing"
)

func TestIsExists(t *testing.T) {
	if !IsExists(".") {
		t.Error(". must exist")
		return
	}
}

func TestIsDir(t *testing.T) {
	if !IsDir(".") {
		t.Error(". should be a directory")
		return
	}
}

func TestReadLine(t *testing.T) {
	lchan, err := ReadLine("../testdata/data.txt")
	if err != nil {
		t.Fatal(err)
	}

	for line := range lchan {
		fmt.Println(string(line))
	}
}
