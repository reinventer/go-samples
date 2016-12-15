package expbackoff

import (
	"math/rand"
	"sync"
	"time"
)

type Backoffer interface {
	Fetch() <-chan struct{}
	Jam()
	Reset()
	Stop()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New(slotTime time.Duration, maxTime time.Duration) Backoffer {
	b := &backoff{
		c:       make(chan struct{}),
		done:    make(chan struct{}),
		stopped: make(chan struct{}),
		reset:   make(chan struct{}),
	}

	go b.run(slotTime, maxTime)

	return b
}

type backoff struct {
	sync.RWMutex
	c        chan struct{}
	done     chan struct{}
	stopped  chan struct{}
	reset    chan struct{}
	attempts uint32
}

func (b *backoff) Fetch() <-chan struct{} {
	return b.c
}

func (b *backoff) Jam() {
	b.Lock()
	if b.attempts < 30 { // max possible shifts for int32
		b.attempts++
	}
	b.Unlock()
}

func (b *backoff) Reset() {
	b.Lock()
	b.attempts = 0
	b.reset <- struct{}{}
	b.Unlock()
}

func (b *backoff) Stop() {
	close(b.done)
	<-b.stopped
	close(b.c)
}

func (b *backoff) run(slotTime time.Duration, maxTime time.Duration) {
	var (
		nanos int64
		wait  time.Duration
		ok    bool
	)

	for {
		b.RLock()
		var attempts = b.attempts
		b.RUnlock()

		if attempts == 0 {
			if ok = b.tick(); !ok {
				return
			}
			continue
		}

		nanos = (1<<attempts - 1) * slotTime.Nanoseconds()
		if nanos > maxTime.Nanoseconds() {
			nanos = maxTime.Nanoseconds()
		}

		wait = time.Duration(rand.Int63n(nanos))

		select {
		case <-b.done:
			close(b.stopped)
			return
		case <-b.reset:
			if ok = b.tick(); !ok {
				return
			}
		case <-time.After(wait):
			if ok = b.tick(); !ok {
				return
			}
		}
	}
}

func (b *backoff) tick() bool {
	select {
	case b.c <- struct{}{}:
		return true
	case <-b.done:
		close(b.stopped)
		return false
	}
}
