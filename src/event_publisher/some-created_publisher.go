package event_publisher

import "github.com/ssibrahimbas/micro-template/src/event"

func (ep *Publisher) PublishSomeCreated(e *event.SomeCreated) {
	ep.parseAndPublish("Some.Created", e)
}
