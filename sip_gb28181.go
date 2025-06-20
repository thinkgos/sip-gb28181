package sip_gb28181

// ContentType define
const (
	// ContentTypeSDP SDP ContentType
	ContentTypeSDP = "application/sdp"
	// ContentTypeXML XML ContentType
	ContentTypeXML = "Application/MANSCDP+xml"
)

const (
	CmdControl  = "Control"
	CmdQuery    = "Query"
	CmdNotify   = "Notify"
	CmdResponse = "Response"
)

const (
	CmdType_Keepalive     = "Keepalive"
	CmdType_Catalog       = "Catalog"
	CmdType_DeviceInfo    = "DeviceInfo"
	CmdType_RecordInfo    = "RecordInfo"
	CmdType_DeviceControl = "DeviceControl"
)
