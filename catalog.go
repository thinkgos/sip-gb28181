package sip_gb28181_extra

import "encoding/xml"

type CatalogQuery struct {
	XMLName  xml.Name `xml:"Query"`
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
	Parental     string `xml:"Parental"`
	ParentId     string `xml:"ParentID"`
	RegisterWay  string `xml:"RegisterWay"`
	Secrecy      string `xml:"Secrecy"`
	Status       string `xml:"Status"`
}

// <?xml version="1.0" encoding="GB2312"?>
// <Query>
// <CmdType>Catalog</CmdType>
// <SN>588056</SN>
// <DeviceID>34020000001320000001</DeviceID>
// </Query>

// <?xml version="1.0" encoding="GB2312"?>
// <Response>
// <CmdType>Catalog</CmdType>
// <SN>588056</SN>
// <DeviceID>34020000001320000001</DeviceID>
// <SumNum>1</SumNum>
// <DeviceList Num="1">
// <Item>
// <DeviceID>34020000001320000001</DeviceID>
// <Name>TL-IPC45AW-COLOR 5.1</Name>
// <Manufacturer>TP-LINK</Manufacturer>
// <Model>IPCamera 01</Model>
// <Owner>Owner</Owner>
// <CivilCode>3402000000</CivilCode>
// <Address>Address</Address>
// <Parental>0</Parental>
// <ParentID>3402000000200000001</ParentID>
// <RegisterWay>1</RegisterWay>
// <Secrecy>0</Secrecy>
// <Status>ON</Status>
// </Item>
// </DeviceList>
// <Result>OK</Result>
// </Response>
