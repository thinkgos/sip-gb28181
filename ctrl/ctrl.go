package ctrl

const DeviceControl_Magic byte = 0xA5

// CombinedCode1 组合码1, 高4位bit代表的是版本信息, 低4位bit代表的是校验位.
// 校验位 = (字节1的高4位+字节1的低4位+字节2的高4位)%16
func CombinedCode1(magic, ver byte) byte {
	checksum := (magic>>4)&0x0f + magic&0x0f + ver&0x0f
	version := (ver & 0x0f) << 4
	return version | checksum
}

// 组合码2
// 高4位是数据3，低4位是地址的高4位, 没有特别指明高4位, 表明与所指定功能无关.
func CombinedCode2(data3, addr byte) byte {
	return (data3&0x0f)<<4 + addr&0x0f
}

// Checksum 计算校验和, 对所有字节求和后取低8位.
func Checksum(bs []byte) byte {
	sum := int(0)
	for _, v := range bs {
		sum += int(v)
	}
	return byte(sum & 0xff)
}
