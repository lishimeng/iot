package nwconnector

import "fmt"

// 网络表结构：
//   network_type:
//      network_id:
//          network_instance_1,2,3...
//

type Schedule interface {
	// Connect 创建连接
	Connect(networkStruct NetworkStruct)
	// DisConnect 销毁连接
	DisConnect(networkStruct NetworkStruct)
	// OnPing 收到Ping(更新timeout)
	OnPing(func(networkStruct NetworkStruct))
}

type NetworkBuilder struct {
	nwt   NetWorkType
	nid   string
	event NetworkEvent
}

func (nb *NetworkBuilder) Type(nwt NetWorkType) *NetworkBuilder {
	nb.nwt = nwt
	return nb
}

func (nb *NetworkBuilder) Id(id string) *NetworkBuilder {
	nb.nid = id
	return nb
}

func (nb *NetworkBuilder) Event(e NetworkEvent) *NetworkBuilder {
	nb.event = e
	return nb
}

func (nb *NetworkBuilder) Build() string {

	return fmt.Sprintf(topicPubTpl, nb.nwt, nb.nid, nb.event.EventType)
}


