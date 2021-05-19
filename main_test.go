package main

import (
	"io/ioutil"
	"testing"
)

type AddResult struct {
	x        int
	y        int
	expected int
}

var addResults = []AddResult{
	{1, 1, 2},
	{-3, -3, -6},
}

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{-1, 1},
		{0, 2},
		{999, 1001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("For the input {}, expected {} value didn't match the output {}", test.input, test.expected, output)
		}
	}
}

func TestAdd(t *testing.T) {
	for _, test := range addResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected result do not match the function output")
		}
	}
}

func TestFileData(t *testing.T) {
	data, err := ioutil.ReadFile("test-data/test.data")

	if err != nil {
		t.Fatal("Error while reading the file")
	}

	if string(data) != "Lorem ipsum dolor sit amet\n" {
		t.Fatal("The expected value did not match the file contents.")
	}
}
