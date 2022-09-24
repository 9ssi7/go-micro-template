## Mapper Folder

Here are the functions that convert DTO's to entities, entities to events and responses.

It is useful to use mapper for data security. It also prevents unnecessary data from being included in objects. That's why using mapper is highly recommended.

### Naming Conventions

The naming convention for mapper files is as follows:

`Map` + `INPUT_NAME` + `To` + `OUTPUT_NAME`

For example, if we want to create a mapper file for User, we will use the following file name:

- `MapSomeEntityToCreatedEvent`
- `MapSomeEntityToFoundDto`
- `MapSomeCreateDtoToEntity`


### Example

```go
func (m *Mapper) MapSomeEntityToCreatedEvent(s *entity.Some) *event.SomeCreated {
	return &event.SomeCreated{
		ID:   s.ID,
		Name: s.Name,
	}
}
```
