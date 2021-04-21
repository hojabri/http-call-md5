package waitgroup

import "sync"

// WaitGroup implements a limit-control goroutine pool.
type WaitGroup struct {
	parallelLimit int
	pool          chan byte
	waitGroup     sync.WaitGroup
}

// NewWaitGroup creates a waitgroup with a specific parallel limit (the maximum number of
// goroutines to run at the same time). If limit=0 or less, it will act as a normal sync.WaitGroup
func NewWaitGroup(limit int) *WaitGroup {
	wg := &WaitGroup{
		parallelLimit: limit,
	}
	if limit > 0 {
		wg.pool = make(chan byte, limit)
	}
	return wg
}

// Add pushes 1 into the pool channel and blocks running extra go routines if the pool channel is full.
func (wg *WaitGroup) Add() {
	if wg.parallelLimit > 0 {
		wg.pool <- 1
	}
	wg.waitGroup.Add(1)
}

// Done pops 1 out the pool channel and makes the inner waitgroup done.
func (wg *WaitGroup) Done() {
	if wg.parallelLimit > 0 {
		<-wg.pool
	}
	wg.waitGroup.Done()
}

// Wait waits until the pool is empty
func (wg *WaitGroup) Wait() {
	wg.waitGroup.Wait()
}

