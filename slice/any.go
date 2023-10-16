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

func Any[T ~[]E, E any](src T) Slice[T, E] {
	return Slice[T, E]{src}
}
