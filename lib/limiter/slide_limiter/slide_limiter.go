package slide_limiter

import (
	"figment/lib/limiter"
	"figment/lib/lock"
	"figment/lib/lock/tlock"
	"time"
)

type SlideCount struct {
	lock       lock.Lock
	unit       int64
	maxConnPer float64
	prevTime   int64
	prevCount  int64
	curCount   int64
	nextTime   int64
}

func New(maxConnPer float64, unit time.Duration) limiter.Limiter {
	now := int64(time.Now().UnixNano())
	return &SlideCount{
		lock:       new(tlock.TLock),
		unit:       int64(unit),
		maxConnPer: maxConnPer,
		prevTime:   now - int64(unit),
		prevCount:  0,
		curCount:   0,
		nextTime:   now + int64(unit),
	}
}

func (s *SlideCount) TryTake(_ []byte) (bool, int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	now := int64(time.Now().UnixNano())

	if now < s.prevTime {
		return false, 0
	}

	for now > s.nextTime {
		s.prevTime = s.nextTime - s.unit
		s.prevCount = s.curCount
		s.curCount = 0
		s.nextTime = s.nextTime + s.unit
	}

	req := float64(s.prevCount)*float64(-now+s.prevTime+2*s.unit)/float64(s.unit) + float64(s.curCount+1)
	if req > s.maxConnPer {
		return false, 0
	}

	s.curCount++

	return true, 0
}
