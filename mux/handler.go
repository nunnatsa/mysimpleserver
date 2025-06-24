package mux

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mysimpleserver/event"
)

func eventHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	e := event.Event{}
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if event.HandleEvent(e) {
		fmt.Fprintf(w, "Event %s handled successfully\n", e.Id)
		return
	}

	fmt.Fprintf(w, "Event %s already processed\n", e.Id)
}
