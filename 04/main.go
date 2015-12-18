package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
	"strings"
	"sync"
)

func main() {
	worker := 4
	close := make(chan struct{})
	input := make(chan int, 10)
	go func() {
		for i := 0; ; i++ {
			input <- i
			if i%10000 == 0 {
				log.Printf("%d\n", i)
			}
		}
	}()
	wg := &sync.WaitGroup{}
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go func() {
			work(input, close)
			wg.Done()
		}()
	}
	wg.Wait()
}

func work(in chan int, cch chan struct{}) {
	var i int
	for {
		select {
		case i = <-in:
		case <-cch:
			return
		}
		r := hash(strconv.Itoa(i))
		if strings.HasPrefix(r, "000000") {
			log.Printf("Found: %d", i)
			close(cch)
			return
		}
	}

}

func hash(i string) string {
	h := md5.Sum([]byte("iwrupvqb" + i))
	return hex.EncodeToString(h[:])
}
