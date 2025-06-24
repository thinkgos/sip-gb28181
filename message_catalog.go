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
	DeviceId     string `xml:"DeviceID"`
	Name         string `xml:"Name"`
	Manufacturer string `xml:"Manufacturer"`
	Model        string `xml:"Model"`
	Owner        string `xml:"Owner"`
	CivilCode    string `xml:"CivilCode"`
	Address      string `xml:"Address"`
	Parental     int    `xml:"Parental"`
	ParentId     string `xml:"ParentID"`
	RegisterWay  int    `xml:"RegisterWay"`
	Secrecy      int    `xml:"Secrecy"`
	Status       string `xml:"Status"`
}
