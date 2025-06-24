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
	if !f(e) {
		return false
	}

	handleOneEvent(e)

	return true
}

var x []string

func f(e Event) bool {
	for i := 0; i < len(x); i++ {
		if x[i] == e.Id {
			return false
		}
	}

	x = append(x, e.Id) // Add the event ID

	return true
}

func handleOneEvent(e Event) {
	fmt.Println("Handling event:", e.Id, e.SomeIntVal, e.SomeFloatVal)
}
