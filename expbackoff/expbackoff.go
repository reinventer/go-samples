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
}

func (b *backoff) run(slotTime time.Duration, maxTime time.Duration) {
	for {
		b.RLock()
		var attempts = b.attempts
		b.RUnlock()

		if attempts == 0 {
			b.tick()
			continue
		}

		wait := time.Duration(rand.Int63n(1<<attempts) * slotTime.Nanoseconds())
		if wait > maxTime {
			wait = maxTime
		}

		select {
		case <-b.done:
			close(b.stopped)
			return
		case <-b.reset:
			b.tick()
		case <-time.After(wait):
			b.tick()
		}
	}
}

func (b *backoff) tick() {
	select {
	case b.c <- struct{}{}:
	case <-b.done:
		close(b.stopped)
		return
	}
}
