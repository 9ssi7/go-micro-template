## EventHandler Folder

This folder contains the code for the event handler. With the `ListenAll` function, events that occur in other microservices are listened. The event is handled by the handlers here as soon as it is publish.

As timing, our objects in internal should be created after they are created so that functions within the ``internal/service`` can be called from here.


### Naming Conventions

The naming convention for event handlers is as follows:

`On` + `ENTITY_NAME` + `ACTION` + `PAST TENSE VERB`

For example, if we want to create a new user, we will use the following event handler:

```go
// user_created.go
package event_handler

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/ssibrahimbas/micro-template/src/event"
)

func (eh *Handler) OnUserCreated(msg *nats.Msg) {
    e := event.UserCreated{}
	err := json.Unmarshal(msg.Data, &e)
	if err != nil {
		return
	}
    eh.s.UserService.Create(&e)
}
```

Listen to the event in the `ListenAll` function:

```go
// event_handler.go
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
	eh.n.AddStream("User")
	eh.subscribeSub("User.Created", eh.OnUserCreated)
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
```