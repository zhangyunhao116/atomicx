package bit

// Warning: offset must smaller or equal to 31.
// If offset larger than 31, undefined behavior will occurs!
func BitTestAndSetUint32(addr *uint32, offset uint32)

func BitTestAndSetUint(addr *uint32, offset uint32)

func BitTestAndSetUint64(addr *uint64, offset uint32)

func BitTestAndResetUint32(addr *uint32, offset uint32)

func BitTestAndResetUint64(addr *uint64, offset uint32)
