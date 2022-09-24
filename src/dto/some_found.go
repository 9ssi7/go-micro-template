package dto

import "time"

type SomeFound struct {
	ID        string    `json:"uuid"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"active"`
	CreatedAt time.Time `json:"dateOfCreate"`
	UpdatedAt time.Time `json:"dateOfUpdate"`
}
