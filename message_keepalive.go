package sip_gb28181

import "encoding/xml"

type Keepalive struct {
	XMLName  xml.Name `xml:"Notify"`
	CmdType  string   `xml:"CmdType"`
	Sn       int      `xml:"SN"`
	DeviceId string   `xml:"DeviceID"`
	Status   string   `xml:"Status"`
}
