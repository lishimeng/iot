package nwconnector

type NetworkStruct struct {
	NetworkId         string // 编号  a
	NetworkInstanceId string // 连接实例编号
	NetWorkType
	ConnectInfo map[string]interface{}
}
type EventType int

const (
	EventCreate  EventType = 1 // 创建连接
	EventDestroy EventType = 2 // 销毁连接
)

type NetWorkType int

type NetworkEvent struct {
	EventType EventType
	NetworkStruct
}
