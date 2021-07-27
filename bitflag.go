package atomicx

import (
	"sync/atomic"
	"unsafe"
)

// BitFlag represents a flag set base on uint32.
//
// const (
// 		F0 = 1 << iota
// 		F1
// 		F2
// )
// f := new(BitFlag)
// f.Set(F0|F1)     // set F0 and F1 to 1.
// f.Get(F0|F1,F1)  // false, check F0 and F1, expect F1 is 1 and F0 is 0.
// f.Unset(F0)      // set F0 to 0.
// f.Get(F0|F1,F1)  // true
type BitFlag uint32

// Set set the corresponding flags to 1.
//
// Set(F1|F2) will set flag F1 and F2 to 1, does not affect the rest of the flags.
func (f *BitFlag) Set(flags uint32) {
	OrUint32IfNeeded((*uint32)(unsafe.Pointer(f)), flags)
}

// Unset set the cooresponding flags to 0.
//
// Unset(F1|F2) will set flag F1 and F2 to 0, does not affect the rest of the flags.
func (f *BitFlag) Unset(flags uint32) {
	p := (*uint32)(unsafe.Pointer(f))
	for {
		old := atomic.LoadUint32(p)
		check := old & flags
		if check != 0 {
			if atomic.CompareAndSwapUint32(p, old, old^check) {
				return
			}
			continue
		}
		return
	}
}

// Get check the corresponding flags, return true if the result is expected.
// check represents the flags you want to check.
// expect represents the flags is 1 in the result you expect.
//
// f.Set(F1|F2|F4)
//
// - Check F1 F2 F3, want F1 and F2 is 1, F3 is 0.
// f.Get(F1|F2|F3,F1|F2) // true
//
// - Check F1 F2 F3, want F1 is 1, but F2 and F3 is 0.
// f.Get(F1|F2|F3,F1)    // false, because F2 is 1, does not match expects.
func (f *BitFlag) Get(check, expect uint32) bool {
	return (atomic.LoadUint32((*uint32)(unsafe.Pointer(f))) & check) == expect
}
