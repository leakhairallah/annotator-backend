package dtos

import (
	"encoding/json"
)

type AnnotationRequest struct {
	Text     string          `json:"text" validate:"required"`
	Metadata json.RawMessage `json:"metadata" validate:"required"`
}
