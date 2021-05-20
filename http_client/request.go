/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2021-05-20 15:17
**/

package utils

import "net/http"

type Request struct {
	err            error
	code           int
	data           []byte
	responseHeader http.Header
	requestHeader  http.Header
}

func (r *Request) String() string {
	return string(r.data)
}

func (r *Request) Bytes() []byte {
	return r.data
}

func (r *Request) Code() int {
	return r.code
}

func (r *Request) LastError() error {
	return r.err
}

func (r *Request) ResponseHeader() http.Header {
	return r.responseHeader
}

func (r *Request) RequestHeader() http.Header {
	return r.requestHeader
}
