package sip_gb28181

import "encoding/xml"

type ReceiveMessage struct {
	XMLName xml.Name
	CmdType string `xml:"CmdType"`
	Sn      int64  `xml:"SN"`
}
