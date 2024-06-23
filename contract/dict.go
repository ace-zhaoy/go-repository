package contract

type Dict[K comparable, V any] interface {
	KeyExists(key K) bool
	Keys() []K
	Value(key K) V
	Get(key K) (v V, ok bool)
	Values() []V
	ForEach(fn func(key K, value V))
	ForEachKey(fn func(key K))
	ForEachValue(fn func(value V))
	ValueExistsBy(fn func(value V) bool) bool
	ValueExists(value V) bool
	Set(key K, value V)
	Delete(key K)
	Len() int
	IsEmpty() bool
	Clear()
	All() map[K]V
	Clone() Dict[K, V]
	Safe(safe bool) Dict[K, V]
}
