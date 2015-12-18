package main

import "testing"

func TestA(t *testing.T) {
	tests := []struct {
		i string
		o int
	}{
		{"()()", 0},
		{"(())", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, tc := range tests {
		if o, _ := process(tc.i); o != tc.o {
			t.Errorf("Input: %s, expected %d, got %d", tc.i, tc.o, o)
		}
	}
}

func TestB(t *testing.T) {
	tests := []struct {
		i string
		o int
	}{
		{")", 1},
		{"()())", 5},
		{"((())))", 7},
	}

	for _, tc := range tests {
		if _, o := process(tc.i); o != tc.o {
			t.Errorf("Input: %s, expected %d, got %d", tc.i, tc.o, o)
		}
	}

}
