package sip_gb28181

import (
	"math/rand/v2"
	"strconv"
	"sync/atomic"
)

var _sn = atomic.Uint32{}

func init() {
	_sn.Store(rand.Uint32())
}

type Sn uint32

func (s Sn) Value() int64 { return int64(s) }

func (s Sn) String() string { return strconv.FormatUint(uint64(s), 10) }

func NextSN() Sn {
	for {
		if v := _sn.Add(1); v > 0 {
			return Sn(v)
		}
	}
}
