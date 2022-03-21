/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2020-01-07 21:38
**/

package utils

import "github.com/lemoyxk/utils/v3/constraints"

func Array[T constraints.Integer](s T, src []T) Arr[T] {
	return Arr[T]{s, src}
}

type Arr[T constraints.Integer] struct {
	s   T
	src []T
}

func (a *Arr[T]) In() bool {
	for _, v := range a.src {
		if v == a.s {
			return true
		}
	}
	return false
}

func (a *Arr[T]) Sum() T {
	var sum T
	for _, v := range a.src {
		sum += v
	}
	return sum
}
