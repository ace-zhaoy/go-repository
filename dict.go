package repository

import (
	"github.com/ace-zhaoy/go-repository/contract"
	"reflect"
)

type Dict[K comparable, V any] struct {
	m map[K]V
}

var _ contract.Dict[struct{}, any] = (*Dict[struct{}, any])(nil)

func NewDict[K comparable, V any](m map[K]V) *Dict[K, V] {
	if m == nil {
		m = make(map[K]V)
	}
	return &Dict[K, V]{
		m: m,
	}
}

func NewDictWithSize[K comparable, V any](size int) *Dict[K, V] {
	return &Dict[K, V]{
		m: make(map[K]V, size),
	}
}

func (d *Dict[K, V]) KeyExists(key K) bool {
	_, ok := d.m[key]
	return ok
}

func (d *Dict[K, V]) Keys() []K {
	keys := make([]K, 0, len(d.m))
	for k := range d.m {
		keys = append(keys, k)
	}
	return keys
}

func (d *Dict[K, V]) Value(key K) V {
	return d.m[key]
}

func (d *Dict[K, V]) Get(key K) (v V, ok bool) {
	v, ok = d.m[key]
	return v, ok
}

func (d *Dict[K, V]) Values() []V {
	values := make([]V, 0, len(d.m))
	for _, v := range d.m {
		values = append(values, v)
	}
	return values
}

func (d *Dict[K, V]) ForEach(fn func(key K, value V)) {
	for k, v := range d.m {
		fn(k, v)
	}
}

func (d *Dict[K, V]) ForEachKey(fn func(key K)) {
	for k := range d.m {
		fn(k)
	}
}

func (d *Dict[K, V]) ForEachValue(fn func(value V)) {
	for _, v := range d.m {
		fn(v)
	}
}

func (d *Dict[K, V]) ValueExistsBy(fn func(value V) bool) bool {
	for _, v := range d.m {
		if fn(v) {
			return true
		}
	}
	return false
}

func (d *Dict[K, V]) ValueExists(value V) bool {
	for _, v := range d.m {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

func (d *Dict[K, V]) Set(key K, value V) {
	d.m[key] = value
}

func (d *Dict[K, V]) Delete(key K) {
	delete(d.m, key)
}

func (d *Dict[K, V]) Len() int {
	return len(d.m)
}

func (d *Dict[K, V]) IsEmpty() bool {
	return len(d.m) == 0
}

func (d *Dict[K, V]) Clear() {
	d.m = make(map[K]V)
}

func (d *Dict[K, V]) All() map[K]V {
	return d.m
}

func (d *Dict[K, V]) Clone() contract.Dict[K, V] {
	m := make(map[K]V)
	for k, v := range d.m {
		m[k] = v
	}
	return NewDict(m)
}

func (d *Dict[K, V]) Safe(safe bool) contract.Dict[K, V] {
	if safe {
		return NewDictSafe(contract.Dict[K, V](d))
	}
	return d
}
