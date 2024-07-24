package internal

import "encoding/json"

type HTTPEvent struct {
	Path    string            `json:"path"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

func HttpEventFromString(input string) *HTTPEvent {
	var event HTTPEvent
	err := json.Unmarshal([]byte(input), &event)
	if err != nil {
		panic(err)
	}
	return &event
}
