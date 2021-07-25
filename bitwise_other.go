// +build !amd64

package atomicx

import "sync/atomic"

func OrUint32(addr *uint32, val uint32) {
	for {
		old := atomic.LoadUint32(addr)
		if old&val != val {
			if atomic.CompareAndSwapUint32(addr, old, old|val) {
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
			if atomic.CompareAndSwapUint64(addr, old, old|val) {
				return
			}
			continue
		}
		return
	}
}

func XorUint32(addr *uint32, val uint32) {
	if val == 0 {
		return
	}
	for {
		old := atomic.LoadUint32(addr)
		if atomic.CompareAndSwapUint32(addr, old, old^val) {
			return
		}
	}
}

func XorUint64(addr *uint64, val uint64) {
	if val == 0 {
		return
	}
	for {
		old := atomic.LoadUint64(addr)
		if atomic.CompareAndSwapUint64(addr, old, old^val) {
			return
		}
	}
}

func AndUint32(addr *uint32, val uint32) {
	for {
		old := atomic.LoadUint32(addr)
		if old != old&val {
			if atomic.CompareAndSwapUint32(addr, old, old&val) {
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
			if atomic.CompareAndSwapUint64(addr, old, old&val) {
				return
			}
			continue
		}
		return
	}
}

func OrUint32IfNeeded(addr *uint32, val uint32) { OrUint32(addr, val) }

func OrUint64IfNeeded(addr *uint64, val uint64) { OrUint64(addr, val) }

func AndUint32IfNeeded(addr *uint32, val uint32) { AndUint32(addr, val) }

func AndUint64IfNeeded(addr *uint64, val uint64) { AndUint64(addr, val) }
