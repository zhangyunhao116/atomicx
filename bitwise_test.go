package atomicx

import (
	"math"
	"math/rand"
	"sync"
	"testing"
)

func TestOr(t *testing.T) {
	var x uint32
	OrUint32(&x, 24)
	if x != 24 {
		t.Fatal("invalid")
	}

	OrUint32(&x, 4)
	if x != 28 {
		t.Fatal("invalid")
	}
	OrUint32(&x, 28)
	if x != 28 {
		t.Fatal("invalid")
	}

	var y uint64
	OrUint64(&y, 24<<50)
	if y != 24<<50 {
		t.Fatal("invalid")
	}

	OrUint64(&y, 4<<50)
	if y != 28<<50 {
		t.Fatal("invalid")
	}
	OrUint64(&y, 28<<50)
	if y != 28<<50 {
		t.Fatal("invalid")
	}
}

func TestAnd(t *testing.T) {
	x := uint32(math.MaxUint32)
	AndUint32(&x, 28)
	if x != 28 {
		t.Fatal("invalid")
	}

	AndUint32(&x, 24)
	if x != 24 {
		t.Fatal("invalid")
	}
	AndUint32(&x, 24)
	if x != 24 {
		t.Fatal("invalid")
	}

	y := uint64(math.MaxUint64)
	AndUint64(&y, 28<<50)
	if y != 28<<50 {
		t.Fatal("invalid")
	}

	AndUint64(&y, 24<<50)
	if y != 24<<50 {
		t.Fatal("invalid")
	}
	AndUint64(&y, 24<<50)
	if y != 24<<50 {
		t.Fatal("invalid")
	}
}

func TestXor(t *testing.T) {
	var x uint32
	XorUint32(&x, 0)
	if x != 0 {
		t.Fatal("invalid")
	}

	XorUint32(&x, math.MaxUint32)
	if x != math.MaxUint32 {
		t.Fatal("invalid")
	}

	XorUint32(&x, 28)
	if x != math.MaxUint32-28 {
		t.Fatal("invalid")
	}

	XorUint32(&x, x)
	if x != math.MaxUint32-28^(math.MaxUint32-28) {
		t.Fatal("invalid")
	}

	var y uint64
	XorUint64(&y, 0)
	if y != 0 {
		t.Fatal("invalid")
	}

	XorUint64(&y, math.MaxUint64)
	if y != math.MaxUint64 {
		t.Fatal("invalid")
	}

	XorUint64(&y, 28<<50)
	if y != math.MaxUint64-28<<50 {
		t.Fatal("invalid")
	}

	XorUint64(&y, y)
	if y != math.MaxUint64-28<<50^(math.MaxUint64-28<<50) {
		t.Fatal("invalid")
	}
}

func TestBitWise(t *testing.T) {
	var x1, x2 uint32
	for i := 0; i < 10000; i++ {
		switch rand.Int31n(3) {
		case 0:
			r := rand.Uint32()
			OrUint32(&x1, r)
			x2 = x2 | r
		case 1:
			r := rand.Uint32()
			AndUint32(&x1, r)
			x2 = x2 & r
		case 2:
			r := rand.Uint32()
			XorUint32(&x1, r)
			x2 = x2 ^ r
		}
	}
	if x1 != x2 {
		t.Fatal("invalid")
	}

	var y1, y2 uint64
	for i := 0; i < 10000; i++ {
		switch rand.Int31n(3) {
		case 0:
			r := rand.Uint64()
			OrUint64(&y1, r)
			y2 = y2 | r
		case 1:
			r := rand.Uint64()
			AndUint64(&y1, r)
			y2 = y2 & r
		case 2:
			r := rand.Uint64()
			XorUint64(&y1, r)
			y2 = y2 ^ r
		}
	}
	if y1 != y2 {
		t.Fatal("invalid")
	}
}

func TestBitWiseIfNeeded(t *testing.T) {
	var x1, x2 uint32
	for i := 0; i < 10000; i++ {
		switch rand.Int31n(4) {
		case 0:
			r := rand.Uint32()
			OrUint32IfNeeded(&x1, r)
			x2 = x2 | r
		case 1:
			r := rand.Uint32()
			AndUint32IfNeeded(&x1, r)
			x2 = x2 & r
		}
	}
	if x1 != x2 {
		t.Fatal("invalid")
	}

	var y1, y2 uint64
	for i := 0; i < 10000; i++ {
		switch rand.Int31n(2) {
		case 0:
			r := rand.Uint64()
			OrUint64IfNeeded(&y1, r)
			y2 = y2 | r
		case 1:
			r := rand.Uint64()
			AndUint64IfNeeded(&y1, r)
			y2 = y2 & r
		}
	}
	if y1 != y2 {
		t.Fatal("invalid")
	}
}

func TestRace(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		var (
			x uint32
			y uint64
		)
		i := i
		go func() {
			r := i % 10
			vx := uint32(i)
			vy := uint64(i)
			switch r {
			case 0:
				OrUint32(&x, vx)
			case 1:
				OrUint64(&y, vy)
			case 2:
				AndUint32(&x, vx)
			case 3:
				AndUint64(&y, vy)
			case 4:
				XorUint32(&x, vx)
			case 5:
				XorUint64(&y, vy)
			case 6:
				OrUint32IfNeeded(&x, vx)
			case 7:
				OrUint64IfNeeded(&y, vy)
			case 8:
				AndUint32IfNeeded(&x, vx)
			case 9:
				AndUint64IfNeeded(&y, vy)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
