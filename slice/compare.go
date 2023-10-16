/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-10-16 14:44
**/

package slice

func Compare[T ~[]E, E comparable](src T) Comparable[T, E] {
	return Comparable[T, E]{Slice[T, E]{src: src}}
}

type Comparable[T ~[]E, E comparable] struct {
	Slice[T, E]
}

func (a Comparable[T, E]) Has(s E) bool {
	for i := 0; i < len(a.src); i++ {
		if a.src[i] == s {
			return true
		}
	}
	return false
}

func (a Comparable[T, E]) Index(s E) int {
	for i := 0; i < len(a.src); i++ {
		if a.src[i] == s {
			return i
		}
	}
	return -1
}

func (a Comparable[T, E]) Count(s E) int {
	var count int
	for i := 0; i < len(a.src); i++ {
		if a.src[i] == s {
			count++
		}
	}
	return count
}

func (a Comparable[T, E]) Unique() T {
	var res T
	var mapSet = make(map[E]bool)
	for i := 0; i < len(a.src); i++ {
		if _, ok := mapSet[a.src[i]]; !ok {
			mapSet[a.src[i]] = true
			res = append(res, a.src[i])
		}
	}
	return res
}

func (a Comparable[T, E]) Diff(s T) T {
	var res T
	var mapSet = make(map[E]bool)
	for i := 0; i < len(s); i++ {
		mapSet[s[i]] = true
	}
	for i := 0; i < len(a.src); i++ {
		if _, ok := mapSet[a.src[i]]; !ok {
			res = append(res, a.src[i])
		}
	}
	return res
}

func (a Comparable[T, E]) Intersect(s T) T {
	var res T
	var mapSet = make(map[E]bool)
	for i := 0; i < len(s); i++ {
		mapSet[s[i]] = true
	}
	for i := 0; i < len(a.src); i++ {
		if _, ok := mapSet[a.src[i]]; ok {
			res = append(res, a.src[i])
		}
	}
	return res
}

func (a Comparable[T, E]) Union(s T) T {
	var res T
	var mapSet = make(map[E]bool)
	for i := 0; i < len(a.src); i++ {
		mapSet[a.src[i]] = true
	}
	for i := 0; i < len(s); i++ {
		mapSet[s[i]] = true
	}
	for k := range mapSet {
		res = append(res, k)
	}
	return res
}
