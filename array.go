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

import "github.com/lemonyxk/utils/v3/constraints"

func Ordered[T constraints.Ordered](src *[]T) ordered[T] {
	return ordered[T]{array: array[T]{src: src}}
}

type ordered[T constraints.Ordered] struct {
	// src *[]T
	array[T]
}

func (a ordered[T]) Sum() T {
	var sum T
	var src = *a.src
	for _, v := range src {
		sum += v
	}
	return sum
}

func (a ordered[T]) Max() T {

	var src = *a.src

	if len(src) == 0 {
		panic("array is empty")
	}

	var max = src[0]

	for i := 0; i < len(src); i++ {
		if src[i] > max {
			max = src[i]
		}
	}

	return max
}

func (a ordered[T]) Min() T {

	var src = *a.src

	if len(src) == 0 {
		panic("array is empty")
	}

	var min = src[0]

	for i := 0; i < len(src); i++ {
		if src[i] < min {
			min = src[i]
		}
	}

	return min
}

func (a ordered[T]) Asc() {
	var src = *a.src
	var data = src
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	*a.src = data
}

func (a ordered[T]) Desc() {
	var src = *a.src
	var data = src
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	*a.src = data
}

func Array[T comparable](src *[]T) array[T] {
	return array[T]{src}
}

type array[T comparable] struct {
	src *[]T
}

func (a array[T]) Has(s T) bool {
	var src = *a.src
	for i := 0; i < len(src); i++ {
		if src[i] == s {
			return true
		}
	}
	return false
}

func (a array[T]) Index(s T) int {
	var src = *a.src
	for i := 0; i < len(src); i++ {
		if src[i] == s {
			return i
		}
	}
	return -1
}

func (a array[T]) Slice(start, end int) []T {

	var src = *a.src

	var res []T

	if start < 0 {
		start = len(src) + start
	}

	if end < 0 {
		end = len(src) + end
	}

	if start > end {
		panic("start must less than end")
	}

	for i := start; i < end; i++ {
		res = append(res, src[i])
	}

	return res
}

func (a array[T]) Splice(start int, count int, elem ...T) []T {

	var src = *a.src

	if start < 0 {
		panic("start must be greater than 0")
	}

	if start > len(src) {
		panic("start must be less than length of array")
	}

	if start+count > len(src) {
		count = len(src) - start
	}

	var p1 = src[:start]
	var p2 = src[start+count:]

	var p3 = src[start : start+count]

	src = src[0:0]

	src = append(src, p1...)
	src = append(src, elem...)
	src = append(src, p2...)

	*a.src = src

	return p3
}

func (a array[T]) Insert(start int, elem ...T) {

	var src = *a.src

	if start < 0 {
		panic("start must be greater than 0")
	}

	if start > len(src) {
		panic("start must be less than length of array")
	}

	var p1 = src[:start]
	var p2 = src[start:]

	src = src[0:0]

	src = append(src, p1...)
	src = append(src, elem...)
	src = append(src, p2...)

	*a.src = src
}

func (a array[T]) Delete(start int, count int) {

	var src = *a.src

	if start < 0 {
		panic("start must be greater than 0")
	}

	if start > len(src) {
		panic("start must be less than length of array")
	}

	if start+count > len(src) {
		count = len(src) - start
	}

	var p1 = src[:start]
	var p2 = src[start+count:]

	src = src[0:0]

	src = append(src, p1...)
	src = append(src, p2...)

	*a.src = src
}
