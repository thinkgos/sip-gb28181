package ssrc

import "fmt"

// DomainSsrc 域ssrc
// 10位十进制整数字符串
// 格式:
// 第1位: 历史/实时流的标识位, 0: 实时, 1: 历史
// 第2位到第6位取20位SIP监控域ID的4-8位作为域标识.
// 第7位到第10位作为域内媒体流标识, 是一个在当前域内产生的媒体流ssrc值的不重复的4位十进制整数
type DomainSsrc struct {
	live     bool   // 直播/回放
	domainId uint32 // 域id
	seq      uint32 // 序号, 域内媒体流标识
}

// Live 是否直播/回放
func (s *DomainSsrc) Live() bool { return s.live }

// Seq 序号, 域内媒体流标识
func (s *DomainSsrc) Seq() uint32 { return s.seq }

// DomainId 域id
func (s *DomainSsrc) DomainId() uint32 { return s.domainId }

// Value 10位十进制整数字符串, 不足前面补0
func (s *DomainSsrc) Value() string {
	return fmt.Sprintf("%010d", s.value())
}

// HexValue 8位十六进制字符串, 不足前面补0
func (s *DomainSsrc) HexValue() string {
	return fmt.Sprintf("%08X", s.value())
}

func (s *DomainSsrc) value() uint32 {
	v := s.domainId*10000 + s.seq
	if !s.live {
		v += 1000000000
	}
	return v
}
