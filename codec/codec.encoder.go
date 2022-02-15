package codec

import (
	"github.com/lishimeng/iot/module"
)

// Encoder 编码器
type Encoder interface {
	Encode(data module.DeviceData) // 上报
}

type EncoderHandler interface {
	OnDataDownstream(Encoder)
}
