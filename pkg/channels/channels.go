package channels

import (
	"bufio"
	"os"

	"github.com/tadeas-vintrlik/AoC22/pkg/generics"
)

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

// Reads file sends paragraphs (chunks of text seperated by empty lines) on the returned channel which is closed on EOF.
// Panics if file could not be read.
func ReadParagraphs(file string) <-chan string {
	c := make(chan string, 50)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	go func() {
		t := ""
		for s.Scan() {
			cur := s.Text()
			if cur == "" {
				c <- t
				t = ""
			} else {
				t += cur + "\n"
			}
		}
		c <- t
		if err := s.Err(); err != nil {
			panic(err)
		}
		close(c)
	}()
	return c
}

func Map[T, V any](c <-chan T, transform func(T) V) <-chan V {
	ret := make(chan V, 50)
	go func() {
		for val := range c {
			ret <- transform(val)
		}
		close(ret)
	}()
	return ret
}

func Flatten[T any](c <-chan []T) <-chan T {
	ret := make(chan T, 50)
	go func() {
		for slice := range c {
			for _, val := range slice {
				ret <- val
			}
		}
		close(ret)
	}()
	return ret
}

func Collect[T any](c <-chan T) []T {
	var ret []T
	for v := range c {
		ret = append(ret, v)
	}
	return ret
}

func Sum[T generics.Summable](c <-chan T) T {
	var ret T
	for v := range c {
		ret += v
	}
	return ret
}
