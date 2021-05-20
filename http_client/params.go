/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2021-05-20 17:06
**/

package utils

import (
	"context"
	"net/http"
)

type params struct {
	info    *httpInfo
	err     error
	request *http.Request

	cancel context.CancelFunc
}

func (p *params) Send() *Request {
	if p.err != nil {
		return &Request{err: p.err}
	}
	return send(p.info, p.request, p.cancel)
}

func (p *params) Abort() {
	p.cancel()
}
