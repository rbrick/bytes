package cache

// Cached value is a value that is recomputed upon expiration of the value.
type CachedValue[T any] interface {
	Value() (T, error)
	Compute() (T, error)
}
