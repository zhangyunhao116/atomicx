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
			mockCAS(&x, 1<<r1)
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

func BenchmarkAND(b *testing.B) {
	var x uint32
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r1, r2 := fastrand.Uint32n(31), fastrand.Uint32n(31)
			mockASM(&x, 1<<r1)
			if r2 == 0 {
				atomic.StoreUint32(&x, 0)
			}
		}
	})
}

func mockCAS(x *uint32, value uint32) {
	for {
		old := atomic.LoadUint32(x)
		if old&value != value {
			// Flag is 0, need set it to 1.
			n := old | value
			if atomic.CompareAndSwapUint32(x, old, n) {
				return
			}
			continue
		}
		return
	}
}

func mockASM(x *uint32, value uint32) {
	old := atomic.LoadUint32(x)
	if old&value != value {
		// Flag is 0, need set it to 1.
		OrUint32(x, value)
	}
}
