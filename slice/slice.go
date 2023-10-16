/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2020-01-07 21:38
**/

package slice

import (
	"sort"
)

type Slice[T ~[]E, E any] struct {
	src T
}

func (a Slice[T, E]) First() E {
	if len(a.src) == 0 {
		var zero E
		return zero
	}
	return a.src[0]
}

func (a Slice[T, E]) Last() E {
	if len(a.src) == 0 {
		var zero E
		return zero
	}
	return a.src[len(a.src)-1]
}

func (a Slice[T, E]) Slice(start, end int) T {

	var res T

	if start < 0 {
		start = len(a.src) + start
	}

	if end < 0 {
		end = len(a.src) + end
	}

	if start > end {
		panic("start must less than end")
	}

	for i := start; i < end; i++ {
		res = append(res, a.src[i])
	}

	return res
}

func (a Slice[T, E]) Splice(start int, count int, elem ...E) T {

	if start < 0 {
		panic("start must be greater than 0")
	}

	if start > len(a.src) {
		panic("start must be less than length of array")
	}

	if start+count > len(a.src) {
		count = len(a.src) - start
	}

	var p1 = a.src[:start]
	var p2 = a.src[start+count:]

	var p3 = a.src[start : start+count]

	a.src = a.src[0:0]

	a.src = append(a.src, p1...)
	a.src = append(a.src, elem...)
	a.src = append(a.src, p2...)

	return p3
}

func (a Slice[T, E]) Insert(start int, elem ...E) {

	if start < 0 {
		panic("start must be greater than 0")
	}

	if start > len(a.src) {
		panic("start must be less than length of array")
	}

	var p1 = a.src[:start]
	var p2 = a.src[start:]

	a.src = a.src[0:0]

	a.src = append(a.src, p1...)
	a.src = append(a.src, elem...)
	a.src = append(a.src, p2...)
}

func (a Slice[T, E]) Delete(start int, count int) {

	if start < 0 {
		panic("start must be greater than 0")
	}

	if start > len(a.src) {
		panic("start must be less than length of array")
	}

	if start+count > len(a.src) {
		count = len(a.src) - start
	}

	var p1 = a.src[:start]
	var p2 = a.src[start+count:]

	a.src = a.src[0:0]

	a.src = append(a.src, p1...)
	a.src = append(a.src, p2...)
}

func (a Slice[T, E]) Push(elem ...E) {
	a.src = append(a.src, elem...)
}

func (a Slice[T, E]) Pop() E {
	var elem = a.src[len(a.src)-1]
	a.src = a.src[:len(a.src)-1]
	return elem
}

func (a Slice[T, E]) Shift() E {
	var elem = a.src[0]
	a.src = a.src[1:]
	return elem
}

func (a Slice[T, E]) UnShift(elem ...E) {
	a.src = append(elem, a.src...)
}

func (a Slice[T, E]) Concat(src ...T) T {
	var res = a.src[:]
	for i := 0; i < len(src); i++ {
		res = append(res, src[i]...)
	}
	return res
}

func (a Slice[T, E]) Reverse() {
	for i := 0; i < len(a.src)/2; i++ {
		a.src[i], a.src[len(a.src)-1-i] = a.src[len(a.src)-1-i], a.src[i]
	}
}

func (a Slice[T, E]) ForEach(fn func(elem E, index int)) {
	for i := 0; i < len(a.src); i++ {
		fn(a.src[i], i)
	}
}

func (a Slice[T, E]) Map(fn func(elem E, index int) E) T {
	var res T
	for i := 0; i < len(a.src); i++ {
		res = append(res, fn(a.src[i], i))
	}
	return res
}

func (a Slice[T, E]) Filter(fn func(elem E, index int) bool) T {
	var res T
	for i := 0; i < len(a.src); i++ {
		if fn(a.src[i], i) {
			res = append(res, a.src[i])
		}
	}
	return res
}

func (a Slice[T, E]) Reduce(fn func(prev E, curr E, index int) E, init E) E {
	var res = init
	for i := 0; i < len(a.src); i++ {
		res = fn(res, a.src[i], i)
	}
	return res
}

func (a Slice[T, E]) Sort(fn func(a E, b E) bool) {
	sort.Slice(a.src, func(i, j int) bool {
		return fn(a.src[i], a.src[j])
	})
}

func (a Slice[T, E]) Data() T {
	return a.src
}
