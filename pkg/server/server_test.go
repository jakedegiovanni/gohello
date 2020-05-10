package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWorldHandler(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, helloWorldEndpoint, nil)
	response := httptest.NewRecorder()

	originWorldData := worldData
	worldMsg := "world"
	worldData = func() string {
		return worldMsg
	}

	worldHandler(response, request)

	gotBody := response.Body.String()
	gotCT := response.HeaderMap.Get("Content-Type")
	gotCode := response.Code

	if gotCode != http.StatusOK {
		t.Errorf("Got %d but wanted %d", gotCode, http.StatusOK)
	}

	if gotCT != "application/json" {
		t.Errorf("Got %s but wanted %s", gotCT, "application/json")
	}

	resp := genericResponse{}
	json.Unmarshal([]byte(gotBody), &resp)
	if resp.Message != worldMsg {
		t.Errorf("Got %s but wanted %s", resp.Message, worldMsg)
	}

	worldData = originWorldData
}

func TestWorldHandlerUnsupportedMethod(t *testing.T) {
	request, _ := http.NewRequest(http.MethodOptions, helloWorldEndpoint, nil)
	response := httptest.NewRecorder()

	worldHandler(response, request)

	gotCode := response.Code
	gotBody := response.Body.String()

	if gotCode != http.StatusNotImplemented {
		t.Errorf("Got %d but wanted %d", gotCode, http.StatusNotImplemented)
	}

	want := fmt.Sprintf(methodNotSupportMessage, http.MethodOptions)
	if gotBody != want {
		t.Errorf("Got %s but wanted %s", gotBody, want)
	}
}
