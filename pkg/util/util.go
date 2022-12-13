package util

import (
	"bufio"
	"os"
)

// TODO: Maybe make a channel pkg module

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

type Summable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Summable](c <-chan T) T {
	var ret T
	for v := range c {
		ret += v
	}
	return ret
}

// TODO: Maybe make a slice pkg module

func SliceContains[T comparable](s []T, c T) bool {
	for _, v := range s {
		if c == v {
			return true
		}
	}
	return false
}

func SliceReverse[T any](a []T) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func SliceFilter[T any](a []T, filter func(T) bool) []T {
	ret := []T{}
	for _, v := range a {
		if filter(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func SliceMap[T, V any](a []T, transform func(T) V) []V {
	ret := make([]V, len(a))
	for i, v := range a {
		ret[i] = transform(v)
	}
	return ret
}

// Just generic utils

type Ordered interface {
	Summable | ~string
}

func Min[T Ordered](s []T) T {
	if len(s) == 0 {
		panic("Min called on empty slice")
	}
	min := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}
