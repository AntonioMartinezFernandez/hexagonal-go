package models

type Healthcheck struct {
	Status string `json:"status" validate:"required"`
}
