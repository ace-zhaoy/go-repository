package repository

import (
	"github.com/ace-zhaoy/go-repository/contract"
	"sort"
)

type Collection[ID comparable, T contract.ENTITY[ID]] struct {
	s []T
}

var _ contract.Collection[int64, contract.ENTITY[int64]] = (*Collection[int64, contract.ENTITY[int64]])(nil)

func NewCollection[ID comparable, T contract.ENTITY[ID]](s []T) *Collection[ID, T] {
	return &Collection[ID, T]{
		s: s,
	}
}

func NewCollectionWithCapacity[ID comparable, T contract.ENTITY[ID]](capacity int) *Collection[ID, T] {
	return &Collection[ID, T]{
		s: make([]T, 0, capacity),
	}
}

func (c *Collection[ID, T]) Clone() contract.Collection[ID, T] {
	s := make([]T, len(c.s))
	copy(s, c.s)
	return NewCollection[ID, T](s)
}

func (c *Collection[ID, T]) All() []T {
	return c.s
}

func (c *Collection[ID, T]) Count() int {
	return len(c.s)
}

func (c *Collection[ID, T]) CountBy(fn func(T) bool) int {
	n := 0
	for i := range c.s {
		if fn(c.s[i]) {
			n++
		}
	}
	return n
}

func (c *Collection[ID, T]) Has(id ID) bool {
	for i := range c.s {
		if c.s[i].GetID() == id {
			return true
		}
	}
	return false
}

func (c *Collection[ID, T]) HasBy(fn func(T) bool) bool {
	for i := range c.s {
		if fn(c.s[i]) {
			return true
		}
	}
	return false
}

func (c *Collection[ID, T]) Filter(fn func(T) bool) contract.Collection[ID, T] {
	s := make([]T, 0)
	for i := range c.s {
		if fn(c.s[i]) {
			s = append(s, c.s[i])
		}
	}
	return NewCollection[ID, T](s)
}

func (c *Collection[ID, T]) IsEmpty() bool {
	return len(c.s) == 0
}

func (c *Collection[ID, T]) IDs() []ID {
	ids := make([]ID, 0, len(c.s))
	for i := range c.s {
		ids = append(ids, c.s[i].GetID())
	}
	return ids
}

func (c *Collection[ID, T]) IDsBy(fn func(T) bool) []ID {
	ids := make([]ID, 0)
	for i := range c.s {
		if fn(c.s[i]) {
			ids = append(ids, c.s[i].GetID())
		}
	}
	return ids
}

func (c *Collection[ID, T]) Chunk(size int) []contract.Collection[ID, T] {
	if size <= 0 {
		return nil
	}
	s := make([]contract.Collection[ID, T], 0, len(c.s)/size+1)
	for i := 0; i < len(c.s); i += size {
		end := i + size
		if end > len(c.s) {
			end = len(c.s)
		}
		s = append(s, NewCollection[ID, T](c.s[i:end]))
	}
	return s
}

func (c *Collection[ID, T]) Sort(fn func(T, T) bool) {
	sort.SliceStable(c.s, func(i, j int) bool {
		return fn(c.s[i], c.s[j])
	})
}

func (c *Collection[ID, T]) ForEach(fn func(T)) {
	for i := range c.s {
		fn(c.s[i])
	}
}

func (c *Collection[ID, T]) ForEachWithIndex(fn func(index int, entity T)) {
	for i := range c.s {
		fn(i, c.s[i])
	}
}

func (c *Collection[ID, T]) Range(fn func(T) (finished bool)) {
	for i := range c.s {
		if fn(c.s[i]) {
			return
		}
	}
}

func (c *Collection[ID, T]) RangeWithIndex(fn func(index int, entity T) (finished bool)) {
	for i := range c.s {
		if fn(i, c.s[i]) {
			return
		}
	}
}

func (c *Collection[ID, T]) Intersect(collection contract.Collection[ID, T]) contract.Collection[ID, T] {
	m, s := make(map[ID]struct{}, len(c.s)), make([]T, 0, len(c.s))
	collection.ForEach(func(entity T) {
		m[entity.GetID()] = struct{}{}
	})
	for i := range c.s {
		if _, ok := m[c.s[i].GetID()]; ok {
			s = append(s, c.s[i])
		}
	}
	return NewCollection[ID, T](s)
}

func (c *Collection[ID, T]) Difference(collection contract.Collection[ID, T]) contract.Collection[ID, T] {
	m, s := make(map[ID]struct{}, len(c.s)), make([]T, 0, len(c.s))
	collection.ForEach(func(entity T) {
		m[entity.GetID()] = struct{}{}
	})
	for i := range c.s {
		if _, ok := m[c.s[i].GetID()]; !ok {
			s = append(s, c.s[i])
		}
	}
	return NewCollection[ID, T](s)
}

func (c *Collection[ID, T]) Union(collection contract.Collection[ID, T]) contract.Collection[ID, T] {
	m, s := make(map[ID]struct{}, len(c.s)), make([]T, 0, len(c.s)+collection.Count())
	for i := range c.s {
		if _, ok := m[c.s[i].GetID()]; !ok {
			s = append(s, c.s[i])
			m[c.s[i].GetID()] = struct{}{}
		}
	}
	collection.ForEach(func(entity T) {
		if _, ok := m[entity.GetID()]; !ok {
			s = append(s, entity)
			m[entity.GetID()] = struct{}{}
		}
	})
	return NewCollection[ID, T](s)
}

func (c *Collection[ID, T]) Unique() contract.Collection[ID, T] {
	m, s := make(map[ID]struct{}, len(c.s)), make([]T, 0, len(c.s))
	for i := range c.s {
		if _, ok := m[c.s[i].GetID()]; !ok {
			m[c.s[i].GetID()] = struct{}{}
			s = append(s, c.s[i])
		}
	}
	return NewCollection[ID, T](s)
}

func (c *Collection[ID, T]) Reverse() contract.Collection[ID, T] {
	for i, j := 0, len(c.s)-1; i < j; i, j = i+1, j-1 {
		c.s[i], c.s[j] = c.s[j], c.s[i]
	}
	return c
}

func (c *Collection[ID, T]) Append(elem T) {
	c.s = append(c.s, elem)
}

func (c *Collection[ID, T]) LPush(elem T) {
	c.s = append([]T{elem}, c.s...)
}

func (c *Collection[ID, T]) LPop() T {
	elem := c.s[0]
	c.s = c.s[1:]
	return elem
}

func (c *Collection[ID, T]) RPush(elem T) {
	c.s = append(c.s, elem)
}

func (c *Collection[ID, T]) RPop() T {
	elem := c.s[len(c.s)-1]
	c.s = c.s[:len(c.s)-1]
	return elem
}

func (c *Collection[ID, T]) ToDict() contract.Dict[ID, T] {
	m := make(map[ID]T, len(c.s))
	for i := range c.s {
		m[c.s[i].GetID()] = c.s[i]
	}
	return NewDict(m)
}

func (c *Collection[ID, T]) Safe(safe bool) contract.Collection[ID, T] {
	if safe {
		return NewCollectionSafe(contract.Collection[ID, T](c))
	}
	return c
}
