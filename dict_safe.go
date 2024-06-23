package repository

import (
	"github.com/ace-zhaoy/go-repository/contract"
	"sync"
)

type DictSafe[K comparable, V any] struct {
	d  contract.Dict[K, V]
	mu sync.RWMutex
}

var _ contract.Dict[struct{}, any] = (*DictSafe[struct{}, any])(nil)

func NewDictSafe[K comparable, V any](d contract.Dict[K, V]) *DictSafe[K, V] {
	return &DictSafe[K, V]{
		d: d,
	}
}

func (d *DictSafe[K, V]) KeyExists(key K) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.d.KeyExists(key)
}

func (d *DictSafe[K, V]) Keys() []K {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.d.Keys()
}

func (d *DictSafe[K, V]) Value(key K) V {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.d.Value(key)
}

func (d *DictSafe[K, V]) Get(key K) (v V, ok bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.d.Get(key)
}

func (d *DictSafe[K, V]) Values() []V {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.d.Values()
}

func (d *DictSafe[K, V]) ForEach(fn func(key K, value V)) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	d.d.ForEach(fn)
}

func (d *DictSafe[K, V]) ForEachKey(fn func(key K)) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	d.d.ForEachKey(fn)
}

func (d *DictSafe[K, V]) ForEachValue(fn func(value V)) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	d.d.ForEachValue(fn)
}

func (d *DictSafe[K, V]) ValueExistsBy(fn func(value V) bool) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.d.ValueExistsBy(fn)
}

func (d *DictSafe[K, V]) ValueExists(value V) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.d.ValueExists(value)
}

func (d *DictSafe[K, V]) Set(key K, value V) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.d.Set(key, value)
}

func (d *DictSafe[K, V]) Delete(key K) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.d.Delete(key)
}

func (d *DictSafe[K, V]) Len() int {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.d.Len()
}

func (d *DictSafe[K, V]) IsEmpty() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.d.IsEmpty()
}

func (d *DictSafe[K, V]) Clear() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.d.Clear()
}

func (d *DictSafe[K, V]) All() map[K]V {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.d.All()
}

func (d *DictSafe[K, V]) Clone() contract.Dict[K, V] {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return NewDictSafe(d.d.Clone())
}

func (d *DictSafe[K, V]) Safe(safe bool) contract.Dict[K, V] {
	d.mu.Lock()
	defer d.mu.Unlock()
	if safe {
		return d
	}
	return d.d
}
