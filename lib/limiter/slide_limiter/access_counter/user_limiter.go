package access_counter

import (
	"figment/lib/limiter"
	"figment/lib/limiter/slide_limiter"
	"figment/lib/lock"
	"figment/lib/lock/tlock"
	"time"
)

type AccessCounter struct {
	lock       lock.Lock
	list       map[string]limiter.Limiter
	maxConnPer float64
	unit       time.Duration
}

func New(maxConnPer float64, unit time.Duration) limiter.Limiter {
	return &AccessCounter{
		lock:       new(tlock.TLock),
		list:       nil,
		maxConnPer: maxConnPer,
		unit:       unit,
	}
}

func (acc *AccessCounter) TryTake(key []byte) (bool, int) {
	acc.lock.Lock()
	defer acc.lock.Unlock()

	if acc.list == nil {
		acc.list = make(map[string]limiter.Limiter)
	}

	slide, ok := acc.list[string(key)]
	if !ok {
		slide = slide_limiter.New(acc.maxConnPer, acc.unit)
		acc.list[string(key)] = slide
	}

	return slide.TryTake(nil)
}
