package world

import "testing"

func TestWorldMessage(t *testing.T) {
	msg := Hello()
	if msg != worldGreeting {
		t.Errorf("Wanted %s but got %s", worldGreeting, msg)
	}
}
