package main

import (
	"testing"
)

func TestDefaultPort(t *testing.T) {
	servStartOrig := serverStart
	defer func() {
		serverStart = servStartOrig
	}()

	called := 0
	serverStart = func(port int) {
		called = port
	}

	main()

	if called != defaultPort {
		t.Errorf("Got %d but wanted %d", called, defaultPort)
	}
}

// we are now not testing the flag parsing logic at all, how can we add a test for that?
func TestPassedInPort(t *testing.T) {
	servStartOrig := serverStart
	parseFlagsOrig := parseFlags
	defer func() {
		serverStart = servStartOrig
		parseFlags = parseFlagsOrig
	}()

	wantedPort := 8082
	called := 0

	serverStart = func(port int) {
		called = port
	}
	parseFlags = func() *opts {
		return &opts{port: wantedPort}
	}

	main()

	if called != wantedPort {
		t.Errorf("Got %d but wanted %d", called, wantedPort)
	}
}
