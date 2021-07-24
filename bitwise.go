package atomicx

import "sync/atomic"

func OrUint32IfNeeded(addr *uint32, val uint32) {
	old := atomic.LoadUint32(addr)
	if old&val != val {
		OrUint32(addr, val)
	}
}

func OrUint64IfNeeded(addr *uint64, val uint64) {
	old := atomic.LoadUint64(addr)
	if old&val != val {
		OrUint64(addr, val)
	}
}

func AndUint32IfNeeded(addr *uint32, val uint32) {
	old := atomic.LoadUint32(addr)
	if old != old&val {
		AndUint32(addr, val)
	}
}

func AndUint64IfNeeded(addr *uint64, val uint64) {
	old := atomic.LoadUint64(addr)
	if old != old&val {
		AndUint64(addr, val)
	}
}
