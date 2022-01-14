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

type B struct {
	Name int `json:"name"`
}

func main() {
	var a = A{Name: "hello"}
	var b = A{Name: "world"}
	var c = []A{a, b}
	var d = []interface{}{a, b, c}
	var e = map[string]interface{}{"Name": 111}
	var res = utils.Extract.Src(d).Field("Name").String()
	log.Println(res)

	log.Println(utils.Structure.GetTags(A{}))

	utils.Assign.Dest(&a).Src(&e).AllowWeak().Do()

	log.Println(a)

	var aa = A{Name: "hello"}
	var bb = B{Name: 11111}
	utils.Assign.Dest(&bb).Src(&aa).AllowWeak().Do()
	log.Println(bb)
}
