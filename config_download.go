package sip_gb28181_extra

import "encoding/xml"

type ConfigDownloadQuery struct {
	XMLName    xml.Name `xml:"Query"`
	CmdType    string   `xml:"CmdType"`    // M, 命令类型
	Sn         int      `xml:"SN"`         // M, sn
	DeviceId   string   `xml:"DeviceID"`   // M, 目标设备的编码
	ConfigType string   `xml:"ConfigType"` // M, 查询配置参数类型
}
type ConfigDownloadResponse struct {
	XMLName    xml.Name    `xml:"Response"`
	CmdType    string      `xml:"CmdType"`
	Sn         string      `xml:"SN"`
	DeviceId   string      `xml:"DeviceID"`
	BasicParam *BasicParam `xml:"BasicParam"`
	Result     string      `xml:"Result"`
}
type BasicParam struct {
	Name              string `xml:"Name"`              // 设备名称
	Expiration        string `xml:"Expiration"`        // 注册过期时间
	HeartBeatInterval string `xml:"HeartBeatInterval"` // 心跳间隔时间
	HeartBeatCount    string `xml:"HeartBeatCount"`    // 心跳超时次数
}

// <?xml version="1.0" encoding="GB2312"?>
// <Query><CmdType>ConfigDownload</CmdType><SN>1</SN><DeviceID>34020000001320000001</DeviceID><ConfigType>BasicParam</ConfigType></Query>
