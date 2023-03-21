/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-11-04 16:15
**/

package utils

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	ra "math/rand"
	"strconv"
	"sync"
	"time"
)

type rd struct {
	mux sync.Mutex
}

var Rand = &rd{}

// RandomIntn [begin,end)
func (r *rd) RandomIntn(start int, end int) int {
	r.mux.Lock()
	defer r.mux.Unlock()
	if start == end {
		return start
	}
	if start > end {
		panic("start can not bigger than end")
	}
	return start + ra.New(ra.NewSource(time.Now().UnixNano())).Intn(end-start)
}

// RandomFloat64n [begin,end)
func (r *rd) RandomFloat64n(start float64, end float64) float64 {
	r.mux.Lock()
	defer r.mux.Unlock()
	if start == end {
		return start
	}
	if start > end {
		panic("start can not bigger than end")
	}
	return start + (end-start)*ra.New(ra.NewSource(time.Now().UnixNano())).Float64()
}

func (r *rd) UUID() string {
	r.mux.Lock()
	defer r.mux.Unlock()
	var bytes = make([]byte, 16)
	var _, err = io.ReadFull(rand.Reader, bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func (r *rd) OrderID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.Itoa(r.RandomIntn(10000, 100000))
}
