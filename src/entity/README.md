## Entity Folder

This folder contains all the entities used in the application. Entities are used to store data in the database.

The types here should not be used directly without transforming them.

### Naming Conventions

The naming convention for entities is as follows:

`ENTITY_NAME`

For example, if we have a user, we will use the following entity:

```go

// user.go
package entity

type User struct {
    ID uint `json:"uuid" bson:"_id,omitempty"`
    Name string `json:"name" bson:"name"`
    Email string `json:"email" bson:"email"`
    Password string `json:"password,omitempty" bson:"password"`
}
```