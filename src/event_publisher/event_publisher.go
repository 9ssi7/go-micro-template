package event_publisher

import (
	"encoding/json"
	"fmt"

	"github.com/ssibrahimbas/ssi-core/pkg/helper"
	"github.com/ssibrahimbas/ssi-core/pkg/nats"
)

type Publisher struct {
	n *nats.Conn
}

func New(n *nats.Conn) *Publisher {
	return &Publisher{
		n: n,
	}
}

func (ep *Publisher) parseAndPublish(s string, d interface{}) {
	p, err := json.Marshal(&d)
	helper.CheckErr(err)
	_, err = ep.n.Publish(s, p)
	if err != nil {
		// TODO: log error
		fmt.Println(err)
	}
}
