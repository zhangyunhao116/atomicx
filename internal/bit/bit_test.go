package bit

import (
	"sync"
	"testing"
)

func TestBTS(t *testing.T) {
	// Single goroutine.
	var x1 uint32
	BitTestAndSetUint32(&x1, 0)
	if x1 != 1<<0 {
		t.Fatal("invalid")
	}
	BitTestAndSetUint32(&x1, 1)
	if x1 != 1<<0+1<<1 {
		t.Fatal("invalid")
	}
	BitTestAndSetUint32(&x1, 4)
	if x1 != 1<<0+1<<1+1<<4 {
		t.Fatal("invalid", x1)
	}
	BitTestAndSetUint32(&x1, 31)
	if x1 != 1<<0+1<<1+1<<4+1<<31 {
		t.Fatal("invalid")
	}
	BitTestAndSetUint32(&x1, 31)
	if x1 != 1<<0+1<<1+1<<4+1<<31 {
		t.Fatal("invalid")
	}

	// Multiple goroutines.
	x := uint32(0)
	y := uint32(0)
	wg := new(sync.WaitGroup)
	for i := 0; i < 32; i++ {
		wg.Add(1)
		i := i
		go func() {
			BitTestAndSetUint32(&x, uint32(i))
			BitTestAndSetUint32(&x, uint32(i))
			wg.Done()
		}()
		y += 1 << i
	}
	wg.Wait()
	if x != y {
		t.Fatal("invalid", x, y)
	}

	a := uint64(0)
	b := uint64(0)
	for i := 0; i < 64; i++ {
		wg.Add(1)
		i := i
		go func() {
			BitTestAndSetUint64(&a, uint32(i))
			BitTestAndSetUint64(&a, uint32(i))
			wg.Done()
		}()
		b += 1 << i
	}
	wg.Wait()
	if a != b {
		t.Fatal("invalid", a, b)
	}
}
