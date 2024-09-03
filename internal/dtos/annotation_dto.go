package dtos

import (
	"encoding/json"
)

type Annotation struct {
	Text     string          `json:"text" validate:"required,min=1"`
	Metadata json.RawMessage `json:"metadata" validate:"required"`
}
