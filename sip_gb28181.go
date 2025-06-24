package sip_gb28181

// ContentType define
const (
	// ContentTypeSDP SDP ContentType
	ContentTypeSDP = "application/sdp"
	// ContentTypeXML XML ContentType
	ContentTypeXML = "Application/MANSCDP+xml"
)

const (
	Cmd_Control  = "Control"
	Cmd_Query    = "Query"
	Cmd_Notify   = "Notify"
	Cmd_Response = "Response"
)

const (
	CmdType_Keepalive     = "Keepalive"
	CmdType_Catalog       = "Catalog"
	CmdType_DeviceInfo    = "DeviceInfo"
	CmdType_RecordInfo    = "RecordInfo"
	CmdType_DeviceControl = "DeviceControl"
)
