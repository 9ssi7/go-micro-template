## Event Folder

This folder contains the types of all events that are published and handled within the application. It is recommended to transform when using these types.

We communicate with events between our microservices and our data here will be shared with our other microservices along with events.

The events here can be the types of events published from our other microservices or t he type of events published from our current microservice to other microservices.

### Naming Conventions

The naming convention for events is as follows:

`ENTITY_NAME` + `ACTION` + `PAST TENSE VERB`

For example, if we want to create a new user, we will use the following event:

```go
// user_created.go
package event

type UserCreated struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
```
