package tlock

import (
	"runtime"
	"sync/atomic"
)

type TLock int64

func (l *TLock) Lock() {
	for {
		if l.TryLock() {
			return
		}
		runtime.Gosched()
	}
}

func (l *TLock) TryLock() bool {
	return atomic.CompareAndSwapInt64((*int64)(l), 0, 1)
}

func (l *TLock) Unlock() {
	atomic.StoreInt64((*int64)(l), 0)
}
