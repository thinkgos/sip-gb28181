package sip_gb28181_extra

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Encode_Keepalive(t *testing.T) {
	want := Keepalive{
		CmdType:  "Keepalive",
		Sn:       1,
		DeviceId: "deviceId",
		Status:   "大家好",
	}
	getData, err := MarshalXML(want)
	require.NoError(t, err)
	t.Logf("got: %v\r\n", string(getData))

	want.XMLName = xml.Name{Local: CmdNotify}
	got := Keepalive{}
	err = UnmarshalXML(getData, &got)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func Test_Decode_Keepalive(t *testing.T) {
	data := `<?xml version="1.0" encoding="GB2312"?>
<Notify>
<CmdType>Keepalive</CmdType>
<SN>9</SN>
<DeviceID>34020000001320000001</DeviceID>
<Status>OK</Status>
</Notify>`
	want := Keepalive{
		XMLName:  xml.Name{Local: CmdNotify},
		CmdType:  "Keepalive",
		Sn:       9,
		DeviceId: "34020000001320000001",
		Status:   "OK",
	}
	got := Keepalive{}
	err := UnmarshalXML([]byte(data), &got)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
