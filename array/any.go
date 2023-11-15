/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-11-13 11:26
**/

package array

import (
	"github.com/lemonyxk/utils/slice"
)

type Slice[T any, P any] struct {
	src []T
}

func Any[T any, P any](src []T) Slice[T, P] {
	return Slice[T, P]{src}
}

func (a Slice[T, P]) Slice() slice.Slice[[]T, T] {
	var s = slice.Slice[[]T, T]{}
	s.Set(a.src)
	return s
}

func (a Slice[T, P]) Map(fn func(T, int) P) []P {
	var res []P
	for i := 0; i < len(a.src); i++ {
		res = append(res, fn(a.src[i], i))
	}
	return res
}
