package ssrc

import (
	"context"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/bits-and-blooms/bitset"
)

const DOMAIN_SSRC_SEQ_MAX = 9999

// DomainSsrcGen 单个域的SSRC生成器
// 每个域内最多生成9999个SSRC([1,9999]).
type DomainSsrcGen struct {
	mu       sync.Mutex
	domainId uint32
	seq      uint32
	seqUsed  *bitset.BitSet
}

// NewSsrcDomainGen
// SIP监控域ID的4-8位作为域标识.
// 每个域内最多生成9999个SSRC([1,9999]).
func NewDomainSsrcGen(domainId uint32) *DomainSsrcGen {
	return &DomainSsrcGen{
		mu:       sync.Mutex{},
		domainId: domainId,
		seq:      rand.Uint32N(DOMAIN_SSRC_SEQ_MAX + 1),
		seqUsed:  bitset.New(DOMAIN_SSRC_SEQ_MAX + 1),
	}
}

// SetInUsed 设置已使用, 一般用于初始化时设置.
func (s *DomainSsrcGen) SetInUsed(seqs ...uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, seq := range seqs {
		s.seq = max(s.seq, seq)
		s.seqUsed.Set(uint(seq))
	}
}

// Get 获取域ssrc, 成功立马占用, 返回false, 表示获取失败, 一般为seq用尽.
// NOTE: 使用完成, 请及时归还
func (s *DomainSsrcGen) Get(live bool) (DomainSsrc, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()
	for {
		s.mu.Lock()
		if s.seq = s.seq + 1; s.seq > DOMAIN_SSRC_SEQ_MAX {
			s.seq = 1
		}
		seq := s.seq
		if !s.seqUsed.Test(uint(seq)) {
			s.seqUsed.Set(uint(seq))
			domainId := s.domainId
			s.mu.Unlock()
			return DomainSsrc{live: live, seq: seq, domainId: domainId}, true
		}
		s.mu.Unlock()
		select {
		case <-ctx.Done():
			return DomainSsrc{}, false
		default:
		}
	}
}

// Put 归还域ssrc
func (s *DomainSsrcGen) Put(seq uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.seqUsed.Clear(uint(seq))
}
