package main

import "testing"

func TestIsSafe(t *testing.T) {
	tests := []struct {
		Name     string
		Report   []int
		Expected bool
	}{
		{
			Name:     "[7, 6, 4, 2, 1]",
			Report:   []int{7, 6, 4, 2, 1},
			Expected: true,
		},
		{
			Name:     "[1, 2, 7, 8, 9]",
			Report:   []int{1, 2, 7, 8, 9},
			Expected: false,
		},
		{
			Name:     "[9, 7, 6, 2, 1]",
			Report:   []int{9, 7, 6, 2, 1},
			Expected: false,
		},
		{
			Name:     "[1, 3, 2, 4, 5]",
			Report:   []int{1, 3, 2, 4, 5},
			Expected: false,
		},
		{
			Name:     "[8, 6, 4, 4, 1]",
			Report:   []int{8, 6, 4, 4, 1},
			Expected: false,
		},
		{
			Name:     "[1, 3, 6, 7, 9]",
			Report:   []int{1, 3, 6, 7, 9},
			Expected: true,
		},
	}

	for i, testCase := range tests {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := isSafe(testCase.Report)
			if actual != testCase.Expected {
				t.Errorf("FAILED case %v - %s: Expected %v, actual %v", i, testCase.Name, testCase.Expected, actual)
			}
		})
	}
}
func TestIsSafeWithDampener(t *testing.T) {
	tests := []struct {
		Name     string
		Report   []int
		Expected bool
	}{
		{
			Name:     "[7, 6, 4, 2, 1]",
			Report:   []int{7, 6, 4, 2, 1},
			Expected: true,
		},
		{
			Name:     "[1, 2, 7, 8, 9]",
			Report:   []int{1, 2, 7, 8, 9},
			Expected: false,
		},
		{
			Name:     "[9, 7, 6, 2, 1]",
			Report:   []int{9, 7, 6, 2, 1},
			Expected: false,
		},
		{
			Name:     "[1, 3, 2, 4, 5]",
			Report:   []int{1, 3, 2, 4, 5},
			Expected: true,
		},
		{
			Name:     "[8, 6, 4, 4, 1]",
			Report:   []int{8, 6, 4, 4, 1},
			Expected: true,
		},
		{
			Name:     "[1, 3, 6, 7, 9]",
			Report:   []int{1, 3, 6, 7, 9},
			Expected: true,
		},
	}

	for i, testCase := range tests {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := isSafeWithDampener(testCase.Report)
			if actual != testCase.Expected {
				t.Errorf("FAILED case %v - %s: Expected %v, actual %v", i, testCase.Name, testCase.Expected, actual)
			}
		})
	}
}
