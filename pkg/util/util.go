package util

import "github.com/tadeas-vintrlik/AoC22/pkg/generics"

func Abs[T generics.Summable](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
