package utils

import (
	"math/rand"
	"sync"
	"sync/atomic"
)

type (
	Strategy uint
)

const (
	Random Strategy = iota
	InOrder
)

func NewBalancer(strategy Strategy) *Balancer {
	return &Balancer{strategy: strategy}
}

type Balancer struct {
	s []string

	strategy Strategy

	mu  sync.RWMutex
	idx atomic.Int32
}

func (b *Balancer) Set(s []string) {
	b.mu.Lock()
	b.s = s
	b.mu.Unlock()

	b.idx.Store(0)
}

func (b *Balancer) Get() string {
	b.mu.RLock()
	val := b.s[b.idx.Load()]
	b.mu.RUnlock()

	switch b.strategy {
	case InOrder:
		b.idx.Add(1)
	default:
		idx := rand.Int31n(int32(len(b.s)))
		b.idx.Store(idx)
	}

	return val
}
