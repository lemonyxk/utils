/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2021-05-20 15:16
**/

package utils

import (
	"net/http"
)

type hc int

const HttpClient hc = iota

func (h hc) NewProgress() *progress {
	return &progress{}
}

func (h hc) New() *httpClient {
	return &httpClient{}
}

func (h hc) Post(url string) *httpInfo {
	var info = &httpClient{method: http.MethodPost, url: url}
	return info.Post(url)
}

func (h hc) Get(url string) *httpInfo {
	var info = &httpClient{method: http.MethodGet, url: url}
	return info.Get(url)
}

func (h hc) Head(url string) *httpInfo {
	var info = &httpClient{method: http.MethodHead, url: url}
	return info.Head(url)
}
