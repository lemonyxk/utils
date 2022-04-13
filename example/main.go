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

	"github.com/lemonyxk/utils/v3"
)

type A struct {
	Name string `json:"name"`
}

type B struct {
	Name int `json:"name"`
}

func main() {

	// log.Println(utils.Sort(1, 2, -1).Asc(1))
	// log.Println(utils.Sort(1, 2, -1).Desc())
	//
	// var a = utils.Ternary(true, 2, 1)
	//
	// log.Println(a)
	//
	// var b = 2
	//
	// b = a
	//
	// log.Println(b)

	var a = A{Name: "hello"}
	var b = A{Name: "world"}
	var c = []A{a, b}
	var d = []any{a, b, c}
	var e = map[string]any{"Name": 111}
	var res = utils.Extract.Src(d).Field("Name").String()
	log.Println(res)

	log.Println(utils.Structure.GetTags(A{}))

	utils.Assign.Dest(&a).Src(&e).AllowWeak().Do()

	log.Println(a)

	var aa = A{Name: "50"}
	var bb = B{Name: 11111}
	utils.Assign.Dest(&bb).Src(&aa).AllowWeak().Do()
	log.Println(bb)
}
