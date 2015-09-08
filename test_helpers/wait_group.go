package test_helpers

import (
	"sync"
)

type WaitGroup interface {
	Add(int)
	Done()
	Wait()
}

// testWg is a replica of a sync.WaitGroup that can be introspected.
type testWg struct {
	sync.Mutex
	count        int
	WaitCalled   chan int
	CountChanged chan int
}

func NewWaitGroup() *testWg {
	return &testWg{
		WaitCalled:   make(chan int, 1),
		CountChanged: make(chan int, 1024),
	}
}

func (wg *testWg) Add(delta int) {
	wg.Lock()
	wg.count++
	wg.CountChanged <- wg.count
	wg.Unlock()
}

func (wg *testWg) Done() {
	wg.Lock()
	wg.count--
	wg.CountChanged <- wg.count
	wg.Unlock()
}

func (wg *testWg) Wait() {
	wg.Lock()
	wg.WaitCalled <- wg.count
	wg.Unlock()
}
