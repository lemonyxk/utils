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

import "github.com/lemonyxk/utils/constraints"

func Number[T ~[]E, E constraints.Number](src T) Numeric[T, E] {
	return Numeric[T, E]{Ordered: Ordered[T, E]{Comparable: Comparable[T, E]{Slice[T, E]{src: src}}}}
}

type Numeric[T ~[]E, E constraints.Number] struct {
	Ordered[T, E]
}

func (a Numeric[T, E]) Sum() E {
	var sum E
	for i := 0; i < len(a.src); i++ {
		sum += a.src[i]
	}
	return sum
}
