package sdp_extend

import "github.com/pion/sdp/v3"

type SessionDescription struct {
	sdp.SessionDescription
	Ssrc Ssrc
}

// `$type=` and CRLF size.
const lineBaseSize = 4

func (s *SessionDescription) Marshal() ([]byte, error) {
	marsh := make(marshaller, 0, s.MarshalSize())
	ss, err := s.SessionDescription.Marshal()
	if err != nil {
		return nil, err
	}
	marsh = append(marsh, ss...)
	marsh.addKeyValue("y=", s.Ssrc.marshalInto)
	return marsh, nil
}

// MarshalSize returns the size of the SessionDescription once marshaled.
func (s *SessionDescription) MarshalSize() (marshalSize int) { //nolint:cyclop
	marshalSize = s.SessionDescription.MarshalSize()
	marshalSize += lineBaseSize + s.Ssrc.marshalSize()
	return
}

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
