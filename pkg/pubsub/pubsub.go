package pubsub

import (
	"github.com/golang/protobuf/ptypes/any"
	"github.com/nats-io/nats"
)

type Nats struct {
	conn *nats.Conn
}

func NewNats(url string, opts ...nats.Option) (*Nats, error) {
	conn, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, err
	}
	return &Nats{
		conn: conn,
	}, nil
}

func (n *Nats) Publish(msg *any.Any) error {
	return n.conn.Publish(msg.GetTypeUrl(), msg.GetValue())
}
