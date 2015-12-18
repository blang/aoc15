package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	parttwo()
}

func parttwo() {

	scanner := bufio.NewScanner(os.Stdin)
	var count int
	for scanner.Scan() {
		p := parse(scanner.Text())
		a, b := bound(p)
		r := a + b
		fmt.Printf("%dx%dx%d = %d\n", p[0], p[1], p[2], r)
		count += r
	}
	log.Printf("Res: %d\n", count)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
func partone() {
	scanner := bufio.NewScanner(os.Stdin)
	var count int
	for scanner.Scan() {
		p := parse(scanner.Text())
		r := processA(p)
		fmt.Printf("%dx%dx%d = %d\n", p[0], p[1], p[2], r)
		count += r
	}
	log.Printf("Res: %d\n", count)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

type Dim [3]int

func parse(s string) [3]int {
	p := strings.Split(s, "x")
	if len(p) != 3 {
		log.Fatalf("Error could not split: %s", p)
	}
	var err error
	l, err := strconv.Atoi(p[0])
	if err != nil {
		log.Fatalf("Could not parse number: %s", p[0])
	}
	w, err := strconv.Atoi(p[1])
	if err != nil {
		log.Fatalf("Could not parse number: %s", p[1])
	}
	h, err := strconv.Atoi(p[2])
	if err != nil {
		log.Fatalf("Could not parse number: %s", p[2])
	}

	return [3]int{l, w, h}
}
func bound(dims [3]int) (int, int) {
	a, b := min2(dims[0], dims[1], dims[2])

	return a*2 + b*2, dims[0] * dims[1] * dims[2]
}
func min2(a, b, c int) (o1 int, o2 int) {
	if a < b {
		if b < c {
			o2 = b
		} else {
			o2 = c
		}
		o1 = a
	} else {
		if a < c {
			o2 = a
		} else {
			o2 = c
		}
		o1 = b
	}
	return
}
func min(a, b, c int) int {
	p := 0
	if a < b {
		p = a
	} else {
		p = b
	}

	if c < p {
		return c
	}
	return p
}
func processA(dims [3]int) int {
	a := dims[0] * dims[1]
	b := dims[1] * dims[2]
	c := dims[0] * dims[2]
	m := min(a, b, c)
	return 2*(a+b+c) + m

}
