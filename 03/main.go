package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	partB()
}

func partB() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	s := &Santa{}
	r := &Santa{}
	m := make(map[Santa]struct{})
	instr := 0
	for _, c := range strings.TrimSpace(string(b)) {
		if instr%2 == 0 {
			m[*s] = struct{}{}
			s.move(c)
		} else {
			m[*r] = struct{}{}
			r.move(c)
		}
		instr++
	}
	log.Printf("Instr: %d, Houses: %d\n", instr, len(m))
}
func partA() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	s := &Santa{}
	m := make(map[Santa]struct{})
	instr := 0
	for _, r := range strings.TrimSpace(string(b)) {
		m[*s] = struct{}{}
		s.move(r)
		instr++
	}
	log.Printf("Instr: %d, Houses: %d\n", instr, len(m))
}

type Santa struct {
	x int
	y int
}

func (s *Santa) move(r rune) (x, y int) {
	switch r {
	case '^':
		s.y++
	case 'v':
		s.y--
	case '>':
		s.x++
	case '<':
		s.x--
	default:
		log.Fatalf("Invalid move: %c %d\n", r, r)
	}
	return s.x, s.y
}
