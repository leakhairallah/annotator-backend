package models

import (
	"encoding/json"
)

type Annotation struct {
	Id       int             `json:"id"`
	Text     string          `json:"text"`
	Metadata json.RawMessage `json:"metadata"`
}
