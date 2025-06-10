package zcache

type Cache[K comparable, V any] interface {
	Get(key K) (value V, found bool)
	Set(key K, value V) error
}

func NewCache[K comparable, V any]() Cache[K, V] {
	// Call private function in the implementation file.
	return newZCache[K, V]()
}
