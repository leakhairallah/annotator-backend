package models

import (
	"encoding/json"
)

type Annotation struct {
	Id       int64           `json:"id" validate:"required,gte=0"`
	Text     string          `json:"text" validate:"required,min=1"`
	Metadata json.RawMessage `json:"metadata" validate:"required"`
}
