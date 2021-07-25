package atomicx

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestCAS2(t *testing.T) {
	if runtime.GOARCH != "amd64" && runtime.GOARCH != "arm64" {
		t.Skip("unsupprted arch")
	}
	x := new(Uint128)
	x[0], x[1] = 123, 456
	var y Uint128
	y[0], y[1] = 123, 456

	if LoadUint128(x) != y {
		t.Fatal("invalid")
	}
	// old != real old, fail
	if CompareAndSwapUint128(x, Uint128{}, y) {
		t.Fatal("invalid")
	}
	// old == real old, success
	y[0], y[1] = 111, 222
	if !CompareAndSwapUint128(x, *x, y) {
		t.Fatal("invalid")
	}
	if LoadUint128(x) != y {
		t.Fatal("invalid")
	}
	// Reset x.
	if !CompareAndSwapUint128(x, LoadUint128(x), Uint128{}) {
		t.Fatal("invalid")
	}

	// Concurrent CAS2.
	var (
		wg     sync.WaitGroup
		succ   int64
		p1, p2 uint64 = 123, 456
	)
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		wg.Add(1)
		go func() {
			for {
				old := LoadUint128(x)
				new := old
				if new[0] < p1 {
					new[0]++
				} else {
					new[1]++
					if new[1] > p2 {
						break
					}
				}
				if CompareAndSwapUint128(x, old, new) {
					atomic.AddInt64(&succ, 1)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if x[0] != p1 || x[1] != p2 || succ != int64(p1+p2) {
		t.Fatal("invalid")
	}
}
