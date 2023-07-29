/**
* @program: utils
*
* @description:
*
* @author: lemon
*
* @create: 2023-04-19 17:33
**/

package args

import "os"

func Get(flag ...string) string {
	var args = os.Args[1:]
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(flag); j++ {
			if args[i] == flag[j] {
				if i+1 < len(args) {
					return args[i+1]
				}
			}
		}
	}
	return ""
}

func Has(flag ...string) bool {
	var args = os.Args[1:]
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(flag); j++ {
			if args[i] == flag[j] {
				return true
			}
		}
	}
	return false
}
