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
	Name string `json:"name"`
}

func main() {
	var a = A{Name: "hello"}
	var b = A{Name: "world"}
	var c = []A{a, b}
	var d = []interface{}{a, b, c}
	var res = utils.Extract.Src(d).Field("Name").String()
	log.Println(res)

	log.Println(utils.Structure.GetTags(A{}))
}
