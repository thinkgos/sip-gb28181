package sdp

import "github.com/pion/sdp/v3"

type SessionDescription struct {
	sdp.SessionDescription
	Ssrc *Ssrc
}

func (s *SessionDescription) Marshal() ([]byte, error) {
	marsh := make(marshaller, 0, s.MarshalSize())
	ss, err := s.SessionDescription.Marshal()
	if err != nil {
		return nil, err
	}
	marsh = append(marsh, ss...)
	if s.Ssrc != nil {
		marsh.addKeyValue("y=", s.Ssrc.marshalInto)
	}
	return marsh, nil
}

// MarshalSize returns the size of the SessionDescription once marshaled.
func (s *SessionDescription) MarshalSize() (marshalSize int) { //nolint:cyclop
	marshalSize = s.SessionDescription.MarshalSize()
	if s.Ssrc != nil {
		marshalSize += lineBaseSize + s.Ssrc.marshalSize()
	}
	return
}
