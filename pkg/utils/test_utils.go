package utils

import "encoding/json"

type DummyMetadata struct {
	Comment string `json:"comment"`
}

func CreateDummyMetadata() json.RawMessage {
	dummyMetadata := &DummyMetadata{Comment: "comment"}
	jsonData, _ := json.Marshal(dummyMetadata)
	return jsonData
}
