## EventPublisher Folder

This folder contains the code for the event publisher. We publish an event to share a transaction within the microservice with our other microservices via `EventPublisher`.

`EventPublisher` must be used inside the service.

### Naming Conventions

The naming convention for event publishers is as follows:

`ENTITY_NAME` + `ACTION` + `PAST TENSE VERB`

For example, if we want to create a new user, we will use the following event publisher:

```go
// user_created.go
package event_publisher

import (
    "encoding/json"

    "github.com/ssibrahimbas/micro-template/src/event"
    "github.com/ssibrahimbas/ssi-core/pkg/helper"
)

func (ep *Publisher) UserCreated(e *event.UserCreated) {
    p, err := json.Marshal(e)
    helper.CheckErr(err)
    _, err := ep.n.Publish("user.created", p)
    	if err != nil {
		// TODO: log error
		fmt.Println(err)
	}
}
```