package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error on read: %s", err)
	}
	r, p := process(string(b))
	log.Printf("Result: %d, %d", r, p)
}

func process(s string) (r int, p int) {
	var set bool
	for i, c := range s {
		if c == '(' {
			r++
		} else if c == ')' {
			r--
		}
		if !set && r == -1 {
			set = true
			p = i + 1
		}
	}
	return
}
