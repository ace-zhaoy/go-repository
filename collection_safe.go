package repository

import (
	"github.com/ace-zhaoy/go-repository/contract"
	"sync"
)

type CollectionSafe[ID comparable, T contract.ENTITY[ID]] struct {
	c  contract.Collection[ID, T]
	mu sync.RWMutex
}

var _ contract.Collection[int64, contract.ENTITY[int64]] = (*CollectionSafe[int64, contract.ENTITY[int64]])(nil)

func NewCollectionSafe[ID comparable, T contract.ENTITY[ID]](c contract.Collection[ID, T]) *CollectionSafe[ID, T] {
	return &CollectionSafe[ID, T]{
		c: c,
	}
}

func (c *CollectionSafe[ID, T]) Clone() contract.Collection[ID, T] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return NewCollectionSafe(c.c.Clone())
}

func (c *CollectionSafe[ID, T]) All() []T {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.All()
}

func (c *CollectionSafe[ID, T]) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.Count()
}

func (c *CollectionSafe[ID, T]) CountBy(fn func(T) bool) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.CountBy(fn)
}

func (c *CollectionSafe[ID, T]) Has(id ID) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.Has(id)
}

func (c *CollectionSafe[ID, T]) HasBy(fn func(T) bool) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.HasBy(fn)
}

func (c *CollectionSafe[ID, T]) Filter(fn func(T) bool) contract.Collection[ID, T] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return NewCollectionSafe(c.c.Filter(fn))
}

func (c *CollectionSafe[ID, T]) IsEmpty() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.IsEmpty()
}

func (c *CollectionSafe[ID, T]) IDs() []ID {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.IDs()
}

func (c *CollectionSafe[ID, T]) IDsBy(fn func(T) bool) []ID {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.IDsBy(fn)
}

func (c *CollectionSafe[ID, T]) Chunk(size int) []contract.Collection[ID, T] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	res := c.c.Chunk(size)
	cs := make([]contract.Collection[ID, T], 0, len(res))
	for i := range res {
		cs = append(cs, NewCollectionSafe(res[i]))
	}
	return cs
}

func (c *CollectionSafe[ID, T]) Sort(fn func(T, T) bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.c.Sort(fn)
}

func (c *CollectionSafe[ID, T]) ForEach(fn func(T)) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.c.ForEach(fn)
}

func (c *CollectionSafe[ID, T]) ForEachWithIndex(fn func(index int, entity T)) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.c.ForEachWithIndex(fn)
}

func (c *CollectionSafe[ID, T]) Range(fn func(T) (finished bool)) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.c.Range(fn)
}

func (c *CollectionSafe[ID, T]) RangeWithIndex(fn func(index int, entity T) (finished bool)) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.c.RangeWithIndex(fn)
}

func (c *CollectionSafe[ID, T]) Intersect(collection contract.Collection[ID, T]) contract.Collection[ID, T] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return NewCollectionSafe(c.c.Intersect(collection))
}

func (c *CollectionSafe[ID, T]) Difference(collection contract.Collection[ID, T]) contract.Collection[ID, T] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return NewCollectionSafe(c.c.Difference(collection))
}

func (c *CollectionSafe[ID, T]) Union(collection contract.Collection[ID, T]) contract.Collection[ID, T] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return NewCollectionSafe(c.c.Union(collection))
}

func (c *CollectionSafe[ID, T]) Unique() contract.Collection[ID, T] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return NewCollectionSafe(c.c.Unique())
}

func (c *CollectionSafe[ID, T]) Reverse() contract.Collection[ID, T] {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c = c.c.Reverse()
	return c
}

func (c *CollectionSafe[ID, T]) Append(elem T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c.Append(elem)
}

func (c *CollectionSafe[ID, T]) LPush(elem T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c.LPush(elem)
}

func (c *CollectionSafe[ID, T]) LPop() T {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c.LPop()
}

func (c *CollectionSafe[ID, T]) RPush(elem T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c.RPush(elem)
}

func (c *CollectionSafe[ID, T]) RPop() T {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c.RPop()
}

func (c *CollectionSafe[ID, T]) ToDict() contract.Dict[ID, T] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return NewDictSafe(c.ToDict())
}

func (c *CollectionSafe[ID, T]) Safe(safe bool) contract.Collection[ID, T] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if safe {
		return c
	}

	return c.c
}
