/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2021-05-20 15:18
**/

package utils

import (
	"net/http"
	"net/textproto"
	url2 "net/url"
	"time"
)

type httpInfo struct {
	handler         *httpClient
	headerKey       []string
	headerValue     []string
	cookies         []*http.Cookie
	body            interface{}
	progress        *progress
	userName        string
	passWord        string
	clientTimeout   time.Duration
	proxy           func(*http.Request) (*url2.URL, error)
	dialerKeepAlive time.Duration
}

func (h *httpInfo) Progress(progress *progress) *httpInfo {
	h.progress = progress
	return h
}

func (h *httpInfo) Timeout(timeout time.Duration) *httpInfo {
	h.clientTimeout = timeout
	return h
}

func (h *httpInfo) Proxy(url string) *httpInfo {
	var fixUrl, _ = url2.Parse(url)
	h.proxy = http.ProxyURL(fixUrl)
	return h
}

func (h *httpInfo) KeepAlive(keepalive time.Duration) *httpInfo {
	h.dialerKeepAlive = keepalive
	return h
}

func (h *httpInfo) SetBasicAuth(userName, passWord string) *httpInfo {
	h.userName = userName
	h.passWord = passWord
	return h
}

func (h *httpInfo) SetHeaders(headers map[string]string) *httpInfo {
	h.headerKey = nil
	h.headerValue = nil
	for key, value := range headers {
		h.headerKey = append(h.headerKey, textproto.CanonicalMIMEHeaderKey(key))
		h.headerValue = append(h.headerValue, value)
	}
	return h
}

func (h *httpInfo) AddHeader(key string, value string) *httpInfo {
	h.headerKey = append(h.headerKey, textproto.CanonicalMIMEHeaderKey(key))
	h.headerValue = append(h.headerValue, value)
	return h
}

func (h *httpInfo) SetHeader(key string, value string) *httpInfo {
	for i := 0; i < len(h.headerKey); i++ {
		if textproto.CanonicalMIMEHeaderKey(h.headerKey[i]) == textproto.CanonicalMIMEHeaderKey(key) {
			h.headerValue[i] = value
			return h
		}
	}

	h.headerKey = append(h.headerKey, key)
	h.headerValue = append(h.headerValue, value)
	return h
}

func (h *httpInfo) SetCookies(cookies []*http.Cookie) *httpInfo {
	h.cookies = cookies
	return h
}

func (h *httpInfo) AddCookie(cookie *http.Cookie) *httpInfo {
	for i := 0; i < len(h.cookies); i++ {
		if h.cookies[i].String() == cookie.String() {
			h.cookies[i] = cookie
			return h
		}
	}
	h.cookies = append(h.cookies, cookie)
	return h
}

func (h *httpInfo) Json(body ...interface{}) *params {
	h.SetHeader(contentType, applicationJson)
	h.body = body
	request, cancel, err := getRequest(h.handler.method, h.handler.url, h)
	if err != nil {
		return &params{err: err}
	}
	return &params{info: h, request: request, cancel: cancel}
}

func (h *httpInfo) Query(body ...map[string]interface{}) *params {
	h.body = body
	request, cancel, err := getRequest(h.handler.method, h.handler.url, h)
	if err != nil {
		return &params{err: err}
	}
	return &params{info: h, request: request, cancel: cancel}
}

func (h *httpInfo) Form(body ...map[string]interface{}) *params {
	h.SetHeader(contentType, applicationFormUrlencoded)
	h.body = body
	request, cancel, err := getRequest(h.handler.method, h.handler.url, h)
	if err != nil {
		return &params{err: err}
	}
	return &params{info: h, request: request, cancel: cancel}
}

func (h *httpInfo) Multipart(body ...map[string]interface{}) *params {
	h.SetHeader(contentType, multipartFormData)
	h.body = body
	request, cancel, err := getRequest(h.handler.method, h.handler.url, h)
	if err != nil {
		return &params{err: err}
	}
	return &params{info: h, request: request, cancel: cancel}
}
