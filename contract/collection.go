package contract

type Collection[ID comparable, T ENTITY[ID]] interface {
	Clone() Collection[ID, ENTITY[ID]]
	All() []T
	Count() int
	CountBy(fn func(T) bool) int
	Has(id ID) bool
	HasBy(fn func(T) bool) bool
	Filter(fn func(T) bool) Collection[ID, T]
	IsEmpty() bool
	IDs() []ID
	IDsBy(fn func(T) bool) []ID
	Chunk(size int) []Collection[ID, T]
	Sort(fn func(T, T) bool)
	ForEach(fn func(T))
	ForEachWithIndex(fn func(index int, entity T))
	Range(fn func(T) (finished bool))
	RangeWithIndex(fn func(index int, entity T) (finished bool))
	Intersect(Collection[ID, T]) Collection[ID, T]
	Difference(Collection[ID, T]) Collection[ID, T]
	Union(Collection[ID, T]) Collection[ID, T]
	Unique(Collection[ID, T]) Collection[ID, T]
	Reverse() Collection[ID, T]
	Push(T)
	Pop() T
	ToDict() Dict[ID, T]
}
