package event_handler

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/ssibrahimbas/micro-template/src/event"
)

func (eh *Handler) OnSomeFeatureCreated(msg *nats.Msg) {
	e := event.SomeFeatureCreated{}
	err := json.Unmarshal(msg.Data, &e)
	if err != nil {
		return
	}
	// call service eh.s.SomeFeatureService.DoSomething(e)
}
