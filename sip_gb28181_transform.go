package sip_gb28181_extra

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const XMLHeader_GB2312 = `<?xml version="1.0" encoding="GB2312"?>` + "\n"
const XmlHeader_GBK = `<?xml version="1.0" encoding="GBK"?>` + "\n"

// GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(reader)
}

// UTF-8 转 GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	return io.ReadAll(reader)
}

// GB18030 转 UTF-8
func GbkToGB18030(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder())
	return io.ReadAll(reader)
}

// UTF-8 转 GB18030
func Utf8ToGB18030(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewEncoder())
	return io.ReadAll(reader)
}
