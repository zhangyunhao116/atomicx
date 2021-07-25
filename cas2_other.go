// +build !amd64,!arm64

package atomicx

type Uint128 [2]uint64

func CompareAndSwapUint128(addr *Uint128, old, new Uint128) (swapped bool) {
	panic("not supported")
}

func LoadUint128(addr *Uint128) (val Uint128) {
	panic("not supported")
}
