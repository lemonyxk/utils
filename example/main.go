/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-11 13:08
**/

package main

import (
	"log"

	"github.com/lemoyxk/utils"
)

type A struct {
	Name string
}

func main() {
	// var a = A{Name: "hello"}
	// var b = A{Name: "world"}
	var res = utils.Extract([]interface{}{}).Field("1").String()
	log.Println(res)
}
