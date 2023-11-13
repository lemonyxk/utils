/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-11-13 11:26
**/

package slice

type MapSlice[T any, P any] struct {
	src []T
}

func From[T any, P any](src []T) MapSlice[T, P] {
	return MapSlice[T, P]{src}
}

func (a MapSlice[T, P]) Any() Slice[[]T, T] {
	return Slice[[]T, T]{a.src}
}

func (a MapSlice[T, P]) Map(fn func(T, int) P) []P {
	var res []P
	for i := 0; i < len(a.src); i++ {
		res = append(res, fn(a.src[i], i))
	}
	return res
}
