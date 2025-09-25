package ctrl

import (
	"encoding/hex"
	"strings"
)

const hextable = "0123456789ABCDEF"

// Ptz 指令
type Ptz struct {
	Zoom      *bool // 变倍控制, true: zoom in, false: zoom out
	Tilt      *bool // 倾斜控制, true: up, false: down
	Pan       *bool // 水平控制, true: left, false: right
	PlanSpeed byte  // 水平控制速度, [0,255], 0最慢,255最快.
	TiltSpeed byte  // 垂直控制速度, [0,255], 0最慢,255最快.
	ZoomSpeed byte  // 变倍控制速度, [0,16], 0最慢,16最快.
}

func (p *Ptz) Value() string {
	ctrl := byte(0) // 默认停止
	if p.Zoom != nil {
		if *p.Zoom {
			ctrl |= 0x10 // zoom in
		} else {
			ctrl |= 0x20 // zoom out
		}
	}
	if p.Tilt != nil {
		if *p.Tilt {
			ctrl |= 0x08 // up
		} else {
			ctrl |= 0x04 // down
		}
	}
	if p.Pan != nil {
		if *p.Pan {
			ctrl |= 0x02 // left
		} else {
			ctrl |= 0x01 // right
		}
	}

	var cmd [8]byte

	cmd[0] = DeviceControlMagic
	cmd[1] = byte(CombinedCode1(DeviceControlMagic, 0x00))
	cmd[2] = 0x01              // 固定地址01
	cmd[3] = ctrl              // 表示云台的镜头缩小、镜头放大、上、下、左、右，写入指令码的16进制数
	cmd[4] = p.PlanSpeed       // 表示水平控制速度，写入水平控制方向速度的十六进制数
	cmd[5] = p.TiltSpeed       // 表示垂直控制速度，写入垂直控制方向速度的十六进制数
	cmd[6] = p.ZoomSpeed << 4  // 表示变倍控制速度，写入变倍控制方向速度的十六进制数
	cmd[7] = Checksum(cmd[:7]) // 校验码

	b := strings.Builder{}
	b.Grow(hex.EncodedLen(len(cmd)))
	for _, v := range cmd {
		b.WriteByte(hextable[v>>4])
		b.WriteByte(hextable[v&0x0f])
	}
	return b.String()
}
