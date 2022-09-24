package event_publisher

import "github.com/ssibrahimbas/micro-template/src/event"

func (ep *Publisher) PublishSomeDeleted(e *event.SomeDeleted) {
	ep.parseAndPublish("Some.Deleted", e)
}
