package sip_gb28181

import "encoding/xml"

type DeviceInfoQuery struct {
	XMLName  xml.Name `xml:"Query"`
	CmdType  string   `xml:"CmdType"`  // M, 命令类型
	Sn       int      `xml:"SN"`       // M, sn
	DeviceId string   `xml:"DeviceID"` // M, 目标设备的编码
}
type DeviceInfoResponse struct {
	XMLName      xml.Name `xml:"Response"`
	CmdType      string   `xml:"CmdType"`      // M, 命令类型
	Sn           int      `xml:"SN"`           // M, sn
	DeviceId     string   `xml:"DeviceID"`     // M, 目标设备的编码
	DeviceType   string   `xml:"DeviceType"`   // O, 目标设备的名称
	Manufacturer string   `xml:"Manufacturer"` // O, 设备生产商
	Model        string   `xml:"Model"`        // O, 设备型号
	Firmware     string   `xml:"Firmware"`     // O, 设备固件版本
	MaxCamera    int      `xml:"MaxCamera"`    //
	MaxAlarm     int      `xml:"MaxAlarm"`     //
	Result       string   `xml:"Result"`       // M, 査询结果
}
