package main

import "testing"

func TestNice(t *testing.T) {
	tests := []struct {
		i string
		o bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, tc := range tests {
		if o := nice(tc.i); o != tc.o {
			t.Errorf("Invalid: %s, expected %t got %t", tc.i, tc.o, o)
		}
	}
}
func TestNice2(t *testing.T) {
	tests := []struct {
		i string
		o bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, tc := range tests {
		if o := nice(tc.i); o != tc.o {
			t.Errorf("Invalid: %s, expected %t got %t", tc.i, tc.o, o)
		}
	}
}
