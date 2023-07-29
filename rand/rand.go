/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2019-11-04 16:15
**/

package rand

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	ra "math/rand"
	"strconv"
	"time"
)

// RandomIntn [begin,end)
func RandomIntn(start int, end int) int {
	if start == end {
		return start
	}
	if start > end {
		panic("start can not bigger than end")
	}
	return start + ra.New(ra.NewSource(time.Now().UnixNano())).Intn(end-start)
}

// RandomFloat64n [begin,end)
func RandomFloat64n(start float64, end float64) float64 {
	if start == end {
		return start
	}
	if start > end {
		panic("start can not bigger than end")
	}
	return start + (end-start)*ra.New(ra.NewSource(time.Now().UnixNano())).Float64()
}

func UUID() string {
	var bytes = make([]byte, 16)
	var _, err = io.ReadFull(rand.Reader, bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func OrderID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.Itoa(RandomIntn(10000, 100000))
}
