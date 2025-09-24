package ssrc

import "testing"

func Test_DomainSsrc(t *testing.T) {
	ss := DomainSsrc{
		live:     true,
		seq:      7466,
		domainId: 20000,
	}
	t.Log(ss.Seq())
	t.Log(ss.Value())
	t.Log(ss.HexValue())
}
