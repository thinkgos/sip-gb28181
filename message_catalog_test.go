package sip_gb28181

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Decode_CatalogResponse(t *testing.T) {
	data := `<?xml version="1.0" encoding="GB2312" ?>
<Response>
  <CmdType>Catalog</CmdType>
  <SN>508692</SN>
  <DeviceID>34020000001118000001</DeviceID>
  <SumNum>1</SumNum>
  <DeviceList Num="1">
	<Item>
		<DeviceID>34020000001318000001</DeviceID>
		<Name>IPC111</Name>
		<Manufacturer>Dahua</Manufacturer>
		<Model>IPC-HFW1220M-I1</Model>
		<Owner>0</Owner>
		<CivilCode>6532</CivilCode>
		<Address>axy</Address>
		<Parental>0</Parental>
		<ParentID>34020000001118000001</ParentID>
		<RegisterWay>1</RegisterWay>
		<Secrecy>0</Secrecy>
		<Status>ON</Status>
	</Item>
  </DeviceList>
</Response>`
	want := CatalogResponse{
		XMLName:  xml.Name{Local: "Response"},
		CmdType:  "Catalog",
		Sn:       508692,
		DeviceId: "34020000001118000001",
		SumNum:   1,
		DeviceList: []CatalogDeviceEntry{
			{
				Num: 1,
				Item: CatalogDeviceItem{
					DeviceId:     "34020000001318000001",
					Name:         "IPC111",
					Manufacturer: "Dahua",
					Model:        "IPC-HFW1220M-I1",
					Owner:        "0",
					CivilCode:    "6532",
					Address:      "axy",
					Parental:     0,
					ParentId:     "34020000001118000001",
					RegisterWay:  1,
					Secrecy:      0,
					Status:       "ON",
				},
			},
		},
		Result: "",
	}
	got := CatalogResponse{}
	err := UnmarshalXML([]byte(data), &got)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
