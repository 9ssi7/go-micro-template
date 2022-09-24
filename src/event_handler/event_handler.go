package event_handler

import (
	"fmt"
	"os"

	natsGo "github.com/nats-io/nats.go"
	"github.com/ssibrahimbas/micro-template/src/internal"
	"github.com/ssibrahimbas/ssi-core/pkg/nats"
)

type Handler struct {
	n    *nats.Conn
	s    *internal.Service
	subs []*natsGo.Subscription
}

func New(n *nats.Conn, s *internal.Service) *Handler {
	return &Handler{
		n: n,
		s: s,
	}
}

func (eh *Handler) ListenAll() {
	// eh.n.AddStream("NameSpace")
	// eh.subscribeSub("NameSpace.Event", eh.OnSomeFeatureCreated)
}

func (eh *Handler) subscribeSub(s string, h natsGo.MsgHandler) {
	sub, err := eh.n.Subscribe(s, h)
	if err != nil {
		fmt.Println(err)
		return
	}
	eh.subs = append(eh.subs, sub)
}

func (eh *Handler) OnClose(s os.Signal) {
	for _, sub := range eh.subs {
		_ = sub.Unsubscribe()
	}
}
