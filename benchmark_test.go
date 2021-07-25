package atomicx

import (
	"sync/atomic"
	"testing"

	"github.com/zhangyunhao116/fastrand"
)

func BenchmarkWARM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x uint32
		for pb.Next() {
			r1, r2 := fastrand.Uint32n(31), fastrand.Uint32n(31)
			OrUint32(&x, 1<<r1)
			if r2 == 0 {
				atomic.StoreUint32(&x, 0)
			}
		}
	})
}

func BenchmarkASM(b *testing.B) {
	var x uint32
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r1, r2 := fastrand.Uint32n(31), fastrand.Uint32n(31)
			OrUint32IfNeeded(&x, 1<<r1)
			if r2 == 0 {
				atomic.StoreUint32(&x, 0)
			}
		}
	})
}

func BenchmarkCAS(b *testing.B) {
	var x uint32
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r1, r2 := fastrand.Uint32n(31), fastrand.Uint32n(31)
			mockCAS(&x, 1<<r1)
			if r2 == 0 {
				atomic.StoreUint32(&x, 0)
			}
		}
	})
}

func mockCAS(addr *uint32, val uint32) {
	for {
		old := atomic.LoadUint32(addr)
		if old&val != val {
			if atomic.CompareAndSwapUint32(addr, old, old|val) {
				return
			}
			continue
		}
		return
	}
}
