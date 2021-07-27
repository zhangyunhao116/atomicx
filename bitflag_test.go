package atomicx

import "testing"

type embedBitFlag struct {
	flag BitFlag
}

func TestFlag(t *testing.T) {
	// Correctness.
	const (
		f0 = 1 << iota
		f1
		f2
		f3
		f4
		f5
		f6
		f7
	)
	x := new(embedBitFlag)

	x.flag.Set(f1 | f3)
	if !x.flag.Get(f0|f1|f2|f3, f1|f3) {
		t.Fatal("invalid")
	}
	x.flag.Set(f1)
	x.flag.Set(f1 | f3)
	if x.flag != f1+f3 {
		t.Fatal("invalid")
	}

	x.flag.Unset(f1 | f2)
	if !x.flag.Get(f0|f1|f2|f3, f3) {
		t.Fatal("invalid")
	}
	x.flag.Unset(f1 | f2)
	if x.flag != f3 {
		t.Fatal("invalid")
	}
}
