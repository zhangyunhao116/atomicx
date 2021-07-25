// +build !amd64,gccgo,appengine

package atomicx

func OrUint32(addr *uint32, val uint32) {
	for {
		old := atomic.LoadUint32(addr)
		if old&val != val {
			if atomic.CompareAndSwapUint32(addr, old, old|flags) {
				return
			}
			continue
		}
		return
	}
}

func OrUint64(addr *uint64, val uint64) {
	for {
		old := atomic.LoadUint64(addr)
		if old&val != val {
			if atomic.CompareAndSwapUint64(addr, old, old|flags) {
				return
			}
			continue
		}
		return
	}
}

func AndUint32(addr *uint32, val uint32) {
	for {
		old := atomic.LoadUint32(addr)
		if old != old&val {
			if atomic.CompareAndSwapUint32(addr, old, old&flags) {
				return
			}
			continue
		}
		return
	}
}

func AndUint64(addr *uint64, val uint64) {
	for {
		old := atomic.LoadUint64(addr)
		if old != old&val {
			if atomic.CompareAndSwapUint64(addr, old, old&flags) {
				return
			}
			continue
		}
		return
	}
}

func OrUint32IfNeeded(addr *uint32, val uint32) { return OrUint32(addr, val) }

func OrUint64IfNeeded(addr *uint64, val uint64) { return OrUint64(addr, val) }

func AndUint32IfNeeded(addr *uint32, val uint32) { return AndUint32(addr, val) }

func AndUint64IfNeeded(addr *uint64, val uint64) { return AndUint64(addr, val) }
