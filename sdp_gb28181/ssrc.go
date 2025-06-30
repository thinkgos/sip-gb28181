package sdp_gb28181

// y字段：为10位十进制整数字符串, 表示SSRC值.
// 格式:
// 第1位: 历史或实时体流的标识位, 0: 实时, 1: 历史
// 第2位到第6位取20位SIP监控域ID的4-8位作为域标识.
// 第7位到第10位作为域内媒体流标识, 是一个在当前域内产生的媒体流ssrc值的后4位不重复的4位十进制整数
type Ssrc string

func (s Ssrc) String() string {
	return stringFromMarshal(s.marshalInto, s.marshalSize)
}

func (s Ssrc) marshalInto(b []byte) []byte {
	return append(b, s...)
}

func (s Ssrc) marshalSize() (size int) {
	return len(s)
}
