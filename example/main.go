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
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/lemoyxk/utils/http_client"
)

type A struct {
	Name string
}

var counter int64 = 0

func greet(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Content-Length", strconv.Itoa(1000*1000*5))
	// for i := 0; i < 50; i++ {
	// 	time.Sleep(time.Millisecond * 100)
	// 	fmt.Fprintf(w, "%s", strings.Repeat("h", 1000*100*1))
	// }
	fmt.Printf("\r%d", atomic.AddInt64(&counter, 1))
	fmt.Fprintf(w, "%s", "hello")
}

func main() {
	// var a = A{Name: "hello"}
	// var b = A{Name: "world"}
	// // var c = []A{a, b}
	// var d = []interface{}{a, b, nil}
	// var res = utils.Extract(d).Field("Name").String()
	// log.Println(res)

	go func() {
		http.HandleFunc("/", greet)
		http.ListenAndServe(":8088", nil)
	}()

	var progress = utils.HttpClient.NewProgress()
	progress.Rate(0.01).OnProgress(func(p []byte, current int64, total int64) {
		log.Printf("Downloading... %d %d B complete %.2f", current, total, float64(current)/float64(total))
	})

	// var q = utils.HttpClient.
	// 	Get("https://vdn3.vzuu.com/HD/63bd5f0e-aba6-11eb-aee0-72b968d87e7b-t1-vfdeVljgpm.mp4?auth_key=1621510196-0-0-88f27a5833db9c35ae13da56fd97dfeb").
	// 	Progress(progress).Query().Send()
	//
	// log.Println(q.LastError())

	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		utils.HttpClient.Get("http://127.0.0.1:8088").Query().Send()
	// 	}()
	// }

	var q = utils.HttpClient.Get("https://www.twle.cn/static/js/jquery.min.js").Query().Send()
	log.Println(len(q.Bytes()))
	var h = utils.HttpClient.Head("https://www.twle.cn/static/js/jquery.min.js").Query().Send()
	log.Println(h.ResponseHeader())

	select {}
}
