## DTO Folder

This folder contains all the DTOs used in the application. DTOs are used to transfer data (two way) between the client and the server.
### Naming Conventions

The naming convention for DTO is divided into two. The reason is that we have two different DTO types.

These are the two types of DTOs:

#### Request DTOs

Request DTOs are used to transfer data from the client to the server. The naming convention for these DTOs is as follows:

`ENTITY_NAME` + `ACTION`

For example, if we want to create a new user, we will use the following DTO:

```go
// user_create.go
package dto

type UserCreate struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
}
```

#### Response DTOs

Response DTOs are used to transfer data from the server to the client. The naming convention for these DTOs is as follows:

`ENTITY_NAME` + `ACTION` + `PAST TENSE VERB`

For example, if we want to create a new user, we will use the following DTO:

```go
// user_created.go
package dto

type UserCreated struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
```
