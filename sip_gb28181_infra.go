package sip_gb28181

import (
	"math/rand/v2"
	"sync/atomic"
)

var _sn = atomic.Uint32{}

func init() {
	_sn.Store(rand.Uint32())
}

func NextSN() int64 {
	for {
		if v := _sn.Add(1); v > 0 {
			return int64(v)
		}
	}
}
