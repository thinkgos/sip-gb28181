package sdp

// `$type=` and CRLF size.
const lineBaseSize = 4

// marshaller contains state during marshaling.
type marshaller []byte

func (m *marshaller) addKeyValue(key string, value func([]byte) []byte) {
	*m = append(*m, key...)
	*m = value(*m)
	*m = append(*m, "\r\n"...)
}

func lenUint(i uint64) (count int) {
	if i == 0 {
		return 1
	}
	for i != 0 {
		i /= 10
		count++
	}
	return
}

func lenInt(i int64) (count int) {
	if i < 0 {
		return lenUint(uint64(-i)) + 1
	}

	return lenUint(uint64(i))
}
func stringFromMarshal(marshalFunc func([]byte) []byte, sizeFunc func() int) string {
	return string(marshalFunc(make([]byte, 0, sizeFunc())))
}
