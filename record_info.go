package sip_gb28181

import (
	"encoding/xml"
	"time"
)

type RecordInfoQuery struct {
	XMLName   xml.Name  `xml:"Query"`
	CmdType   string    `xml:"CmdType"`  // M, 命令类型
	Sn        int       `xml:"SN"`       // M, sn
	DeviceId  string    `xml:"DeviceID"` // M, 目标设备的编码
	StartTime time.Time `xml:"StartTime"`
	EndTime   time.Time `xml:"EndTime"`
	Secrecy   int       `xml:"Secrecy"`
	Type      string    `xml:"Type"`
}
type RecordInfoResponse struct {
	XMLName    xml.Name          `xml:"Response"`
	CmdType    string            `xml:"CmdType"`    // M, 命令类型
	Sn         int               `xml:"SN"`         // M, sn
	DeviceId   string            `xml:"DeviceID"`   // M, 目标设备的编码
	SumNum     int               `xml:"SumNum"`     // M, 总数
	RecordList []RecordInfoEntry `xml:"RecordList"` // O, 设备列表
	Result     string            `xml:"Result"`     // M, 査询结果
}
type RecordInfoEntry struct {
	Num  int            `xml:"Num,attr"`
	Item RecordInfoItem `xml:"Item"`
}

type RecordInfoItem struct {
	DeviceId  string    `xml:"DeviceID"`
	Name      string    `xml:"Name"`
	FilePath  string    `xml:"FilePath"`
	Address   string    `xml:"Address"`
	StartTime time.Time `xml:"StartTime"`
	EndTime   time.Time `xml:"EndTime"`
	Secrecy   string    `xml:"Secrecy"`
	Type      string    `xml:"Type"`
}
