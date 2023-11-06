package cache

import "time"

// A cached value that is lazily recomputed
type LazyCachedValue[T any] struct {
	// The underlying current value
	value       *T
	expireAt    time.Time
	expireAfter time.Duration
	compute     func() (T, error)
}

func (lv *LazyCachedValue[T]) Value() (T, error) {
	now := time.Now()
	if now.After(lv.expireAt) || lv.value == nil {
		computed, err := lv.Compute()
		if err != nil {
			return computed, err
		}

		lv.value = &computed
		lv.expireAt = now.Add(lv.expireAfter)
		return computed, nil
	}
	return *lv.value, nil
}

func (lv *LazyCachedValue[T]) Compute() (T, error) {
	return lv.compute()
}

func NewLazilyCachedValue[T any](expiresAfter time.Duration, compute func() (T, error)) *LazyCachedValue[T] {
	return &LazyCachedValue[T]{
		expireAfter: expiresAfter,
		expireAt:    time.Now().Add(expiresAfter),
		value:       nil,
		compute:     compute,
	}
}
