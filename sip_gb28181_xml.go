package sip_gb28181

import (
	"bytes"
	"encoding/xml"
	"io"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func XmlGB18030CharsetReader(charset string, r io.Reader) (io.Reader, error) {
	return simplifiedchinese.GB18030.NewDecoder().Reader(r), nil
}

func MarshalXML(v any) ([]byte, error) {
	data, err := xml.Marshal(v)
	if err != nil {
		return nil, err
	}
	return Utf8ToGB18030(append([]byte(XMLHeader_GB2312), data...))
}

func UnmarshalXML(data []byte, v any) error {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(charset string, r io.Reader) (io.Reader, error) {
		if utf8.Valid(data) {
			return r, nil
		}
		return simplifiedchinese.GB18030.NewDecoder().Reader(r), nil
	}
	err := decoder.Decode(v)
	if err == nil {
		return nil
	}
	// 没带encoding, 且不是UTF8, 则转成GB18030
	return xml.NewDecoder(transform.NewReader(bytes.NewReader(data), simplifiedchinese.GB18030.NewDecoder())).
		Decode(v)
}
