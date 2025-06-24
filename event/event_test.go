package event

import "testing"

func TestHandleEvent(t *testing.T) {
	e1 := Event{Id: "1", SomeIntVal: 10, SomeFloatVal: 20.5}
	e2 := Event{Id: "2", SomeIntVal: 15, SomeFloatVal: 25.5}

	if !HandleEvent(e1) {
		t.Errorf("Expected to handle event 1, but did not")
	}

	if HandleEvent(e1) {
		t.Errorf("Expected to not handle event 1 again, but did handled it")
	}

	if !HandleEvent(e2) {
		t.Errorf("Expected to handle event 2, but did not")
	}

	if HandleEvent(e1) {
		t.Errorf("Expected to not handle event 1 again, after handled event 2, but did handled it")
	}
}
