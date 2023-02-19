package models

type EventProcessorResponse struct {
	Response string `json:"response" validate:"required"`
}
