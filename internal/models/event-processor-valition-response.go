package models

type EventProcessorValidationResponse struct {
	Validated bool `json:"validated" validate:"required"`
}
