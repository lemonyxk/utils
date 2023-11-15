/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-10-16 14:44
**/

package array

import "github.com/lemonyxk/utils/hash"

func Compare[T any, P comparable](src []T) Comparable[T, P] {
	return Comparable[T, P]{Slice[T, P]{src: src}}
}

type Comparable[T any, P comparable] struct {
	Slice[T, P]
}

func (a Comparable[T, P]) Hash(fn func(T, int) P) hash.Hash[P, T] {
	var h = hash.Hash[P, T]{}
	for i := 0; i < len(a.src); i++ {
		h.Set(fn(a.src[i], i), a.src[i])
	}
	return h
}
