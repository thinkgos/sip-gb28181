package sip_gb28181

import "encoding/xml"

type DevicePtzControl struct {
	XMLName  xml.Name             `xml:"Control"`
	CmdType  string               `xml:"CmdType"`  // M, 命令类型
	Sn       int                  `xml:"SN"`       // M, sn
	DeviceId string               `xml:"DeviceID"` // M, 目标设备的编码
	PtzCmd   string               `xml:"PTZCmd"`   // M, PTZ控制命令
	Info     DevicePtzControlInfo `xml:"Info"`     // M, 控制信息
}
type DevicePtzControlInfo struct {
	ControlPriority int `xml:"ControlPriority"` // M, 优先级
}
