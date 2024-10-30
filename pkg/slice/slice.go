package slice

import "github.com/tadeas-vintrlik/AoC22/pkg/generics"

func Contains[T comparable](s []T, c T) bool {
	for _, v := range s {
		if c == v {
			return true
		}
	}
	return false
}

func Reverse[T any](a []T) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func Filter[T any](a []T, filter func(T) bool) []T {
	ret := []T{}
	for _, v := range a {
		if filter(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func Map[T, V any](a []T, transform func(T) V) []V {
	ret := make([]V, len(a))
	for i, v := range a {
		ret[i] = transform(v)
	}
	return ret
}

func Flatten[T any](a [][]T) []T {
	ret := []T{}
	for _, slice := range a {
		ret = append(ret, slice...)
	}
	return ret
}

func Min[T generics.Ordered](s []T) T {
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

func Max[T generics.Ordered](s []T) T {
	if len(s) == 0 {
		panic("Min called on empty slice")
	}
	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}
