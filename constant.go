package sip_gb28181

// 隶属关系
const (
	Parental_Direct  = 0 // 无父设备, 表示设备直属于当前平台
	Parental_Cascade = 1 // 有父设备, 表示设备是从上级平台级联接入的
)

// 安全模式
const (
	SafetyWay_Disabled   = 0 // 不启用安全认证
	SafetyWay_DualAuth   = 1 // 基于双向数字证书认证
	SafetyWay_SingleAuth = 2 // 基于单向数字证书认证
)

// 注册方式
const (
	RegisterWay_UnRegister  = 0 // 0:未注册
	RegisterWay_StandardSip = 1 // 1:标准GB28181的sip注册
	RegisterWay_Passive     = 2 // 2:被动注册
	RegisterWay_IpDirect    = 3 // 3:ip直连
)

// 保密等级
const (
	Secrecy_Public       = 0 // 公开
	Secrecy_Sensitive    = 1 // 敏感
	Secrecy_Confidential = 2 // 机密
)
