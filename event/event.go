package event

import (
	"fmt"
)

type Event struct {
	Id           string  `json:"id"`
	SomeIntVal   int     `json:"someIntVal"`
	SomeFloatVal float64 `json:"someFloatVal"`
}

func HandleEvent(e Event) bool {
	if wasAlreadyHandled(e) {
		return false
	}

	handleOneEvent(e)

	return true
}

var m = make(map[string]bool)

func wasAlreadyHandled(e Event) bool {
	_, exists := m[e.Id]
	if exists {
		return true
	}

	m[e.Id] = true

	return false
}

func handleOneEvent(e Event) {
	fmt.Println("Handling event:", e.Id, e.SomeIntVal, e.SomeFloatVal)
}
