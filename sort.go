package utils

import (
	"github.com/lemoyxk/utils/v3/constraints"
)

type Sortable[T constraints.Ordered] struct {
	data []T
}

func Sort[T constraints.Ordered](arr ...T) Sortable[T] {
	return Sortable[T]{data: arr}
}

func (s Sortable[T]) Asc(a T) []T {
	var data = s.data
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}

func (s Sortable[T]) Desc() []T {
	var data = s.data
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
