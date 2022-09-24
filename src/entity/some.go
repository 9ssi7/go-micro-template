package entity

import "time"

type Some struct {
	ID        string    `json:"uuid,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	IsActive  bool      `json:"isActive,omitempty" bson:"is_active"`
	CreatedAt time.Time `json:"dateOfCreate,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"dateOfUpdate,omitempty" bson:"updated_at,omitempty"`
}
