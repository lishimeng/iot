package nwconnector

import (
	"context"
	"sync"
)

// topic模板： schedule/network/{network_type}/{network_id}/{event_type}
// 示例，mqtt_bridge监听topic： schedule/network/mqtt/+/+
// 数据示例：
//    topic:     schedule/network/mqtt/gasdfg-fasd-sd-fs/1
//    payload:   {
//                    host: "xxxx",
//                    port: 1882,
//                    ssl: false,
//                    userName: "xx",
//                    password: "gdds",
//                    clientId: "fssf"
//               }
// 创建|mqtt_bridge处理| 网络ID：gasdfg-fasd-sd-fs

type Connector interface {
	// Connect 连接
	Connect() error
	// Disconnect 销毁连接
	Disconnect() error
	Ping()
}

// ConnectorPool
//   连接池
//   规则：一个网络ID维持一个连接实例. 同一网络ID如需同时维持多个实例，需要通过启动多个应用实例实现
//   负责管理连接，但不负责创建连接
type ConnectorPool interface {
	Add(id string, c Connector)
	Remove(id string)
	Subscribe(event NetworkEvent)
	NetworkType() NetWorkType
}

type connectorPoolImpl struct {
	p      map[string]Connector
	lock   sync.Mutex
	nwType NetWorkType
	ctx    context.Context
}

func New(ctx context.Context, networkType NetWorkType) (p ConnectorPool) {
	p = &connectorPoolImpl{
		ctx:    ctx,
		p:      make(map[string]Connector),
		lock:   sync.Mutex{},
		nwType: networkType,
	}
	return
}

func (cp *connectorPoolImpl) NetworkType() NetWorkType {
	return cp.nwType
}

// Add 添加一个新连接, 连接实例应是已经创建好
func (cp *connectorPoolImpl) Add(id string, c Connector) {
	cp.lock.Lock()
	defer cp.lock.Unlock()

	if _, ok := cp.p[id]; ok {
		// duplicate id
		return
	}
	cp.p[id] = c
}

// Remove 销毁一个连接
func (cp *connectorPoolImpl) Remove(id string) {
	if c, ok := cp.p[id]; ok {
		delete(cp.p, id)
		go func() {
			e := c.Disconnect()
			if e != nil {
				// TODO
			}
		}()
	}
}

// Subscribe 监听网络创建或销毁
func (cp *connectorPoolImpl) Subscribe(event NetworkEvent) {

	// TODO
	switch event.EventType {
	case EventCreate:
		// TODO
		// cp.Add(event.NetworkId, )
	case EventDestroy:
		// TODO
	}
}
