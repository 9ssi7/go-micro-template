package event

type SomeCreated struct {
	ID   string `json:"uuid,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}
