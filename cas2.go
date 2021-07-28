// +build amd64,!gccgo,!appengine arm64,!gccgo,!appengine

package atomicx

type Uint128 [2]uint64

// CompareAndSwapUint128(CMPXCHG16B) requires that the destination (memory) operand be 16-byte aligned.
// See https://www.felixcloutier.com/x86/cmpxchg8b:cmpxchg16b.
func CompareAndSwapUint128(addr *Uint128, old, new Uint128) (swapped bool)

func LoadUint128(addr *Uint128) (val Uint128)
