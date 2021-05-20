/**
* @program: proxy-server
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-03 13:37
**/

package utils

import (
	"net"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const applicationFormUrlencoded = "application/x-www-form-urlencoded"
const applicationJson = "application/json"
const multipartFormData = "multipart/form-data"
const contentType = "Content-Type"
const contentLength = "Content-Length"

const dialerTimeout = 30 * time.Second
const dialerKeepAlive = 30 * time.Second
const clientTimeout = 15 * time.Second

var defaultDialer = net.Dialer{
	Timeout:   dialerTimeout,
	KeepAlive: dialerKeepAlive,
}

var defaultTransport = http.Transport{
	Proxy:                 http.ProxyFromEnvironment,
	DisableCompression:    false,
	DisableKeepAlives:     false,
	TLSHandshakeTimeout:   10 * time.Second,
	ResponseHeaderTimeout: 15 * time.Second,
	ExpectContinueTimeout: 2 * time.Second,
	MaxIdleConns:          runtime.NumCPU() * 2,
	MaxIdleConnsPerHost:   runtime.NumCPU() * 2,
	MaxConnsPerHost:       runtime.NumCPU() * 2,
	DialContext:           defaultDialer.DialContext,
}

var defaultClient = http.Client{
	Timeout:   clientTimeout,
	Transport: &defaultTransport,
}

var hMux sync.Mutex

type httpClient struct {
	method string
	url    string
}

func (h *httpClient) Post(url string) *httpInfo {
	h.method = http.MethodPost
	h.url = url
	var info = &httpInfo{handler: h}
	info.SetHeader(contentType, applicationFormUrlencoded)
	return info
}

func (h *httpClient) Get(url string) *httpInfo {
	h.method = http.MethodGet
	h.url = url
	var info = &httpInfo{handler: h}
	return info
}

func (h *httpClient) Head(url string) *httpInfo {
	h.method = http.MethodHead
	h.url = url
	var info = &httpInfo{handler: h}
	return info
}
