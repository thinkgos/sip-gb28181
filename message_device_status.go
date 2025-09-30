package sip_gb28181

import (
	"encoding/xml"
	"time"
)

type DeviceStatusQuery struct {
	XMLName  xml.Name `xml:"Query"`
	CmdType  string   `xml:"CmdType"`  // M, 命令类型
	Sn       int64    `xml:"SN"`       // M, sn
	DeviceId string   `xml:"DeviceID"` // M, 目标设备的编码
}
type DeviceStatusResponse struct {
	XMLName     xml.Name                  `xml:"Response"`
	CmdType     string                    `xml:"CmdType"`    // M, 命令类型
	Sn          int64                     `xml:"SN"`         // M, sn
	DeviceId    string                    `xml:"DeviceID"`   // M, 目标设备的编码
	Result      string                    `xml:"Result"`     // M, 査询结果
	Online      string                    `xml:"Online"`     // M, 设备在线
	Status      string                    `xml:"Status"`     // M, 设备状态
	DeviceTime  time.Time                 `xml:"DeviceTime"` // M, 设备时间
	Encode      string                    `xml:"Encode"`     // M, 设备编码
	Record      string                    `xml:"Record"`     // M, 设备是否有录像
	AlarmStatus []DeviceStatusAlarmStatus `xml:"Alarmstatus"`
}

type DeviceStatusAlarmStatus struct {
	Num  int                         `xml:"Num,attr"`
	Item DeviceStatusAlarmStatusItem `xml:"Item"`
}

type DeviceStatusAlarmStatusItem struct{}
