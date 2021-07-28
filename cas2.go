// +build amd64,!gccgo,!appengine arm64,!gccgo,!appengine

package atomicx

import "unsafe"

type Uint128 [2]uint64

// CompareAndSwapUint128(CMPXCHG16B) requires that the destination (memory) operand be 16-byte aligned.
// See https://www.felixcloutier.com/x86/cmpxchg8b:cmpxchg16b.
func CompareAndSwapUint128(addr *Uint128, old, new Uint128) (swapped bool)

func LoadUint128(addr *Uint128) (val Uint128)

type UP struct {
	v1 uint64
	v2 unsafe.Pointer
}

func CompareAndSwapUint64Pointer(addr unsafe.Pointer, old, new UP) (swapped bool)

func LoadUint64Pointer(addr unsafe.Pointer) (val UP)

type PU struct {
	v1 unsafe.Pointer
	v2 uint64
}

func CompareAndSwapPointerUint64(addr unsafe.Pointer, old, new PU) (swapped bool)

func LoadPointerUint64(addr unsafe.Pointer) (val PU)

type TP [2]unsafe.Pointer

func CompareAndSwapTwoPointer(addr unsafe.Pointer, old, new TP) (swapped bool)

func LoadTwoPointer(addr unsafe.Pointer) (val TP)
