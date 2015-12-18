package main

import "reflect"
import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		i string
		o [3]int
	}{

		{"1x2x3", [3]int{1, 2, 3}},
	}
	for _, tc := range tests {
		if o := parse(tc.i); !reflect.DeepEqual(tc.o, o) {
			t.Errorf("Invalid: %s %s", tc.i, o)
		}
	}

}
func TestProcess(t *testing.T) {
	tests := []struct {
		i string
		o int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}
	for _, tc := range tests {
		if o := processA(parse(tc.i)); o != tc.o {
			t.Errorf("Invalid: %s, expected %d, got %d", tc.i, tc.o, o)
		}

	}
}
func TestBounds(t *testing.T) {
	tests := []struct {
		i  string
		o1 int
		o2 int
	}{
		{"2x3x4", 10, 24},
		{"1x1x10", 4, 10},
	}
	for _, tc := range tests {
		o1, o2 := bound(parse(tc.i))
		if o1 != tc.o1 {
			t.Errorf("Invalid 1: %s, expected %d, got %d", tc.i, tc.o1, o1)
		}
		if o2 != tc.o2 {
			t.Errorf("Invalid 2: %s, expected %d, got %d", tc.i, tc.o2, o2)
		}
	}
}
