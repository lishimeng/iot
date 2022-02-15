package bridge

import (
	"github.com/lishimeng/iot/module"
)

// NetworkBridge 网络层桥接器
type NetworkBridge interface {
	UpLink(func(module.DeviceData)) // 设置上报回调
	DownLink(module.DeviceData)     // 下发
}
