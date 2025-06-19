package sip_gb28181

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Decode_DeviceInfoResponse(t *testing.T) {

	data := `<?xml version="1.0" encoding="GB2312" ?>
<Response>
	<CmdType>DeviceInfo</CmdType>
	<SN>508983</SN>
	<DeviceID>34020000001118000001</DeviceID>
	<Result>OK</Result>
	<DeviceType>IPC</DeviceType>
	<Manufacturer>Dahua</Manufacturer>
	<Model>IPC-HFW1220M-I1</Model>
	<Firmware>V2.420.14.R.2016-06-18</Firmware>
	<MaxCamera>1</MaxCamera>
	<MaxAlarm>0</MaxAlarm>
</Response>`

	want := DeviceInfoResponse{
		XMLName:      xml.Name{Local: CmdResponse},
		CmdType:      CmdType_DeviceInfo,
		Sn:           508983,
		DeviceId:     "34020000001118000001",
		DeviceType:   "IPC",
		Manufacturer: "Dahua",
		Model:        "IPC-HFW1220M-I1",
		Firmware:     "V2.420.14.R.2016-06-18",
		Result:       "OK",
	}
	got := DeviceInfoResponse{}
	err := UnmarshalXML([]byte(data), &got)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
