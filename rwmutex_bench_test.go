package spinlock_test

import (
	"github.com/WinnerSoftLab/spinlock"
	"sync"
	"sync/atomic"
	"testing"
)

/*
	This benchmark was written in order to compare original spinlock.RWMutex with forked version.
	Forked version showed no significant benefit, but I believe this should give some impact on
	production because of rare concurrent access and relatively fast operations under lock.
	This test saved here for further research.
*/

var value1 int64 = 0
var value2 int64 = 0
var value3 int64 = 0

func Benchmark_RWMutex_Atomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&value1, 1)
		}
	})
}

func Benchmark_RWMutex_Spinlock(b *testing.B) {
	lock := spinlock.RWMutex{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.RLock()
			value2++
			lock.RUnlock()
		}
	})
}

func Benchmark_RWMutex_Sync(b *testing.B) {
	lock := sync.RWMutex{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.RLock()
			value3++
			lock.RUnlock()
		}
	})
}
