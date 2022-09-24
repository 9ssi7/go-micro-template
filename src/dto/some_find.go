package dto

type SomeFind struct {
	Name string `json:"name,omitempty" validate:"omitempty"`
}
