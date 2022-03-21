/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-12-17 14:42
**/

package utils

func Ternary[T any](b bool, t T, f T) T {
	if b {
		return t
	}
	return f
}
