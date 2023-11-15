/**
* @program: utils
*
* @description:
*
* @author: lemon
*
* @create: 2023-07-19 11:16
**/

package hash

import (
	"github.com/lemonyxk/utils/constraints"
	"sort"
	"sync"
)

type Hash[K comparable, V any] struct {
	mux *sync.RWMutex
	src map[K]V
}

func Any[K comparable, V any](src map[K]V) Hash[K, V] {
	return Hash[K, V]{src: src, mux: &sync.RWMutex{}}
}

func (a Hash[K, V]) Len() int {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return len(a.src)
}

func (a Hash[K, V]) Keys() []K {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var keys []K
	for k := range a.src {
		keys = append(keys, k)
	}
	return keys
}

func (a Hash[K, V]) Values() []V {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var values []V
	for _, v := range a.src {
		values = append(values, v)
	}
	return values
}

func (a Hash[K, V]) ForEach(fn func(k K, v V)) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	for k, v := range a.src {
		fn(k, v)
	}
}

func (a Hash[K, V]) Get(k K) V {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return a.src[k]
}

func (a Hash[K, V]) Set(k K, v V) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.src[k] = v
}

func (a Hash[K, V]) Delete(k K) {
	a.mux.Lock()
	defer a.mux.Unlock()
	delete(a.src, k)
}

func (a Hash[K, V]) Filter(fn func(k K, v V) bool) Hash[K, V] {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var result = make(map[K]V)
	for k, v := range a.src {
		if fn(k, v) {
			result[k] = v
		}
	}
	return Hash[K, V]{src: result}
}

func (a Hash[K, V]) Data() map[K]V {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return a.src
}

type Compare[K comparable, V comparable] struct {
	Hash[K, V]
}

func Comparable[K comparable, V comparable](src map[K]V) Compare[K, V] {
	return Compare[K, V]{Hash[K, V]{src: src, mux: &sync.RWMutex{}}}
}

func (a *Compare[K, V]) Unique(src map[K]V) Compare[K, V] {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var result = make(map[K]V)
	for k, v := range src {
		if _, ok := a.src[k]; !ok {
			result[k] = v
		}
	}
	return Compare[K, V]{Hash[K, V]{src: result}}
}

func (a *Compare[K, V]) Union(src map[K]V) Compare[K, V] {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var result = make(map[K]V)
	for k, v := range a.src {
		result[k] = v
	}
	for k, v := range src {
		result[k] = v
	}
	return Compare[K, V]{Hash[K, V]{src: result}}
}

func (a *Compare[K, V]) Intersect(src map[K]V) Compare[K, V] {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var result = make(map[K]V)
	for k, v := range a.src {
		if _, ok := src[k]; ok {
			result[k] = v
		}
	}
	return Compare[K, V]{Hash[K, V]{src: result}}
}

func (a *Compare[K, V]) Diff(src map[K]V) Compare[K, V] {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var result = make(map[K]V)
	for k, v := range a.src {
		if _, ok := src[k]; !ok {
			result[k] = v
		}
	}
	return Compare[K, V]{Hash[K, V]{src: result}}
}

type Order[K comparable, V constraints.Ordered] struct {
	Hash[K, V]
}

func Ordered[K comparable, V constraints.Ordered](src map[K]V) Order[K, V] {
	return Order[K, V]{Hash[K, V]{src: src, mux: &sync.RWMutex{}}}
}

func (a *Order[K, V]) Sort(fn func(a, b K) bool) []V {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var keys []K
	for k := range a.src {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return fn(keys[i], keys[j])
	})
	var result []V
	for _, k := range keys {
		result = append(result, a.src[k])
	}
	return result
}

func (a *Order[K, V]) Sum() V {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var sum V
	for _, v := range a.src {
		sum += v
	}
	return sum
}

func (a *Order[K, V]) Max() V {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var m V
	for _, v := range a.src {
		if v > m {
			m = v
		}
	}
	return m
}

func (a *Order[K, V]) Min() V {
	a.mux.RLock()
	defer a.mux.RUnlock()
	var m V
	for _, v := range a.src {
		if v < m {
			m = v
		}
	}
	return m
}
