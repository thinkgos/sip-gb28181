package sip_gb28181

import "encoding/xml"

type CatalogQuery struct {
	XMLName  xml.Name `xml:"Query"`    // O, 不填
	CmdType  string   `xml:"CmdType"`  // M, 命令类型
	Sn       int      `xml:"SN"`       // M, sn
	DeviceId string   `xml:"DeviceID"` // M, 目标设备的编码
}
type CatalogResponse struct {
	XMLName    xml.Name             `xml:"Response"`
	CmdType    string               `xml:"CmdType"`    // M, 命令类型
	Sn         int                  `xml:"SN"`         // M, sn
	DeviceId   string               `xml:"DeviceID"`   // M, 目标设备的编码
	SumNum     int                  `xml:"SumNum"`     // M, 总数
	DeviceList []CatalogDeviceEntry `xml:"DeviceList"` // O, 设备列表
	Result     string               `xml:"Result"`     // M, 査询结果
}

type CatalogDeviceEntry struct {
	Num  int               `xml:"Num,attr"`
	Item CatalogDeviceItem `xml:"Item"`
}

type CatalogDeviceItem struct {
	DeviceId     string `xml:"DeviceID"`     // 通道id
	Name         string `xml:"Name"`         // 名称
	Manufacturer string `xml:"Manufacturer"` // 制造商
	Model        string `xml:"Model"`        // 型号
	Owner        string `xml:"Owner"`        // 拥有者
	CivilCode    string `xml:"CivilCode"`    // 民用编码
	Address      string `xml:"Address"`      // 地址
	Parental     int    `xml:"Parental"`     // 隶属关系
	ParentId     string `xml:"ParentID"`     // 父级设备id
	RegisterWay  int    `xml:"RegisterWay"`  // 注册方式
	Secrecy      int    `xml:"Secrecy"`      // 保密等级
	Status       string `xml:"Status"`       // 状态
}
