package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := 0
	for scanner.Scan() {
		if nice2(scanner.Text()) {
			c++
		}
	}
	log.Printf("Res: %d\n", c)
}

func nice2(s string) bool {
	var pre rune
	var prepre rune
	twom := make(map[string]bool)
	var r1 bool
	for _, r := range s {
		if !r1 {
			if _, ok := twom[string(pre)+string(r)]; ok {
				r1 = true
			} else {
				twom[string(pre)+string(r)] = true
			}

		}
		pre = r

	}
	return false
}
func nice(s string) bool {
	var vowels int
	var double int
	var last rune
	for _, r := range s {
		if strings.ContainsRune("aeiou", r) {
			vowels++
		}
		if last == r {
			double++
		}
		if strings.Contains("ab|cd|pq|xy", string(last)+string(r)) {
			return false
		}
		last = r
	}
	return vowels >= 3 && double > 0
}
