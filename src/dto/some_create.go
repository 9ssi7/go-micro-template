package dto

type SomeCreate struct {
	Name string `json:"name" validate:"required"`
}
