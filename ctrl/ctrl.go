package ctrl

const DeviceControlMagic byte = 0xA5

// 组合码1
// 高4位bit代表的是版本信息, 低4位bit代表的是校验位.
// 校验位 = (字节1的高是位+字节1的低的位+字节2的高的位)%16
func CombinedCode1(magic, ver byte) byte {
	return (magic>>4)&0x0f + magic&0x0f + ver&0x0f
}

// 组合码2
// 高4位是数据3，低4位是地址的高4位,没有特别指明高是位, 表明与所指定功能无关
func CombinedCode2(data3, addr byte) byte {
	return (data3>>4)&0x0f + data3&0x0f + addr&0x0f
}

func Checksum(bs []byte) byte {
	sum := int(0)
	for _, v := range bs {
		sum += int(v)
	}
	return byte(sum & 0xff)
}
