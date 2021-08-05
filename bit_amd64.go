// +build amd64,!gccgo,!appengine

package atomicx

// Warning: offset must smaller or equal to 31(for uint32) or 63(for uint64).
// If offset is larger than the limit, undefined behavior will occurs!

func BitSetUint32(addr *uint32, offset uint32)

func BitSetUint64(addr *uint64, offset uint32)

func BitResetUint32(addr *uint32, offset uint32)

func BitResetUint64(addr *uint64, offset uint32)

