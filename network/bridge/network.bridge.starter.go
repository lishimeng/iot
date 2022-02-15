package bridge

import "fmt"

var bridgeList map[string]NetworkBridge

func init() {
	bridgeList = make(map[string]NetworkBridge)
}

func Register(name string, bridge NetworkBridge) {
	bridgeList[name] = bridge
}

func Get(name string) (b NetworkBridge, err error) {
	b, ok := bridgeList[name]
	if !ok {
		err = fmt.Errorf("unknown bridge name: %s", name)
		return
	}
	return
}