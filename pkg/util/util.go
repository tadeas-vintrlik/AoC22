package util

import (
	"bufio"
	"os"
)

func Reverse[T any](a []T) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

// Reads file sends lines on the returned channel which is closed on EOF.
// Panics if file could not be read.
func ReadLines(file string) <-chan string {
	c := make(chan string, 50)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	go func() {
		for s.Scan() {
			c <- s.Text()
		}
		if err := s.Err(); err != nil {
			panic(err)
		}
		close(c)
	}()
	return c
}
