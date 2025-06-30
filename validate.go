package sip_gb28181

import "regexp"

var reSipId = regexp.MustCompile(`^\d{20}$`)
var reSipDomainId = regexp.MustCompile(`^\d{10}$`)

// IsValidSipId 校验是否有效的sip服务器id, 20位数字字符串
func IsValidSipId(sipId string) bool {
	return reSipId.MatchString(sipId)
}

// IsValidSipDomain 校验是否有效的sip服务器域, 10位字符串, sipId的前10位
func IsValidSipDomain(sipDomainId string) bool {
	return reSipDomainId.MatchString(sipDomainId)
}
