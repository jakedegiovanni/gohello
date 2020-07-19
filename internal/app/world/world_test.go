package world

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	h := Handler()
	if h == nil {
		t.Error("Wanted a handler but got nil.")
	}
}

func TestServeHTTP(t *testing.T) {
	handler := handler{}
	t.Run("when getting default message", func(t *testing.T) {
		url := "/"
		r := httptest.NewRequest(http.MethodGet, "http://example.com/world", nil)
		r.URL.Path = url
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Wanted %d but got %d", http.StatusOK, resp.StatusCode)
		}
		if resp.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Wanted %s but got %s", "application/json", resp.Header.Get("Content-Type"))
		}
		var b response
		_ = json.Unmarshal(body, &b)
		if b.Message != worldGreeting {
			t.Errorf("Wanted %s but got %s", worldGreeting, b.Message)
		}
	})

	t.Run("when getting with passed in message", func(t *testing.T) {
		message := "thisisamessage"
		url := fmt.Sprintf("/%s", message)
		r := httptest.NewRequest(http.MethodGet, "http://example.com/world", nil)
		r.URL.Path = url
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Wanted %d but got %d", http.StatusOK, resp.StatusCode)
		}
		if resp.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Wanted %s but got %s", "application/json", resp.Header.Get("Content-Type"))
		}
		var b response
		_ = json.Unmarshal(body, &b)
		if b.Message != message {
			t.Errorf("Wanted %s but got %s", message, b.Message)
		}
	})

	t.Run("when getting with passed in message with a trailing slash", func(t *testing.T) {
		message := "thisisamessage"
		url := fmt.Sprintf("/%s/", message)
		r := httptest.NewRequest(http.MethodGet, "http://example.com/world", nil)
		r.URL.Path = url
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Wanted %d but got %d", http.StatusOK, resp.StatusCode)
		}
		if resp.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Wanted %s but got %s", "application/json", resp.Header.Get("Content-Type"))
		}
		var b response
		_ = json.Unmarshal(body, &b)
		if b.Message != message {
			t.Errorf("Wanted %s but got %s", message, b.Message)
		}
	})

	t.Run("when method isn't GET", func(t *testing.T) {
		url := "/"
		r := httptest.NewRequest(http.MethodPost, "http://example.com/world", nil)
		r.URL.Path = url
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusNotImplemented {
			t.Errorf("Wanted %d but got %d", http.StatusNotImplemented, resp.StatusCode)
		}
	})

	t.Run("when path isn't valid", func(t *testing.T) {
		url := "/something/notvalid"
		r := httptest.NewRequest(http.MethodPost, "http://example.com/world", nil)
		r.URL.Path = url
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Wanted %d but got %d", http.StatusNotFound, resp.StatusCode)
		}
	})

	t.Run("when path isn't valid with a trailing slash", func(t *testing.T) {
		url := "/something/notvalid"
		r := httptest.NewRequest(http.MethodPost, "http://example.com/world", nil)
		r.URL.Path = url
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Wanted %d but got %d", http.StatusNotFound, resp.StatusCode)
		}
	})
}
