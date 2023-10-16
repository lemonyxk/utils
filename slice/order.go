/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-10-16 14:43
**/

package slice

import (
	"github.com/lemonyxk/utils/constraints"
	"sort"
)

func Order[T ~[]E, E constraints.Ordered](src T) Ordered[T, E] {
	return Ordered[T, E]{Comparable: Comparable[T, E]{Slice[T, E]{src: src}}}
}

type Ordered[T ~[]E, E constraints.Ordered] struct {
	Comparable[T, E]
}

func (a Ordered[T, E]) Max() E {

	if len(a.src) == 0 {
		panic("array is empty")
	}

	var item = a.src[0]

	for i := 0; i < len(a.src); i++ {
		if a.src[i] > item {
			item = a.src[i]
		}
	}

	return item
}

func (a Ordered[T, E]) Min() E {

	if len(a.src) == 0 {
		panic("array is empty")
	}

	var item = a.src[0]

	for i := 0; i < len(a.src); i++ {
		if a.src[i] < item {
			item = a.src[i]
		}
	}

	return item
}

func (a Ordered[T, E]) Asc() {
	sort.Slice(a.src, func(i, j int) bool {
		return a.src[i] < a.src[j]
	})
}

func (a Ordered[T, E]) Desc() {
	sort.Slice(a.src, func(i, j int) bool {
		return a.src[i] > a.src[j]
	})
}
