package atomicx

type Uint128 [2]uint64

// CompareAndSwapUint128(CMPXCHG16B) requires that the destination (memory) operand be 16-byte aligned.
// See https://www.felixcloutier.com/x86/cmpxchg8b:cmpxchg16b.
func CompareAndSwapUint128(addr *Uint128, old, new Uint128) (swapped bool)

func LoadUint128(addr *Uint128) (val Uint128)

func BitTestAndSetUint32(addr *uint32, offset uint32)

func BitTestAndSetUint64(addr *uint64, offset uint64)
