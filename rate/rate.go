/**
* @program: utils
*
* @description:
*
* @author: lemon
*
* @create: 2021-11-07 18:10
**/

package rate

import (
	"sync"

	rate2 "golang.org/x/time/rate"
)

func New() *Limiter {
	return &Limiter{keys: make(map[string]*Allow)}
}

type Allow struct {
	limiter *rate2.Limiter
}

// func (a *Allow) SetRate(rate int) *Allow {
// 	a.limiter.SetLimit(rate2.Limit(rate))
// 	return a
// }

func (a *Allow) Allow() bool {
	return a.limiter.Allow()
}

type Limiter struct {
	keys   map[string]*Allow
	global *Allow
	mux    sync.RWMutex
}

func (l *Limiter) Global(rate int) *Allow {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.global == nil {
		l.global = &Allow{limiter: rate2.NewLimiter(rate2.Limit(rate), rate)}
	}
	return l.global
}

func (l *Limiter) Key(key string, rate int) *Allow {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.keys[key] == nil {
		l.keys[key] = &Allow{limiter: rate2.NewLimiter(rate2.Limit(rate), rate)}
	}
	return l.keys[key]
}
