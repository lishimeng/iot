package mq

import (
	"context"
	"github.com/lishimeng/go-log"
)
var MqService Service
type Ds struct {}

func(ds *Ds) Subscribe(topic string, v interface{}, upstream Upstream) {
	log.Info("ds receive:" + topic)
	log.Info("data:")
	s := v.([]byte)
	log.Info(string(s))
}

func(ds *Ds) Topic() string {
	return ""
}

func AmqpStart(ctx context.Context)(err error) {
	ds := Ds{}
	MqService, err = Start(ctx, Connector{Conn: "amqp://office:thingple@192.168.10.76:5672/"}, &ds)
	if err != nil {
		log.Info("Amqp start err ")
		log.Info(err)
		return
	}
	return
}
