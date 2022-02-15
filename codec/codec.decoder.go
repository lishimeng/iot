package codec

import (
	"github.com/lishimeng/iot/module"
)

type Decoder interface {
	Decode(func(data module.NetworkData)) // 下发
}

type DecoderHandler interface {
	OnDataUpstream(Decoder)
}
