package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const mockFormat = "%s : %s"

type mockhandler struct{}

func (m *mockhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf(mockFormat, r.Method, r.URL.Path)))
}

func TestNewFrontController(t *testing.T) {
	s := "front"
	m := map[string]http.Handler{
		s: &mockhandler{},
	}
	fc := NewFrontController(m)
	if fc == nil {
		t.Error("Got nil")
	}
	if _, ok := fc.handlerMap[s]; !ok {
		t.Errorf("%s not mapped correctly", s)
	}
}

func TestFrontController(t *testing.T) {
	s := "front"
	m := map[string]http.Handler{
		s: &mockhandler{},
	}
	fc := FrontController{handlerMap: m}

	t.Run("when path registered", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "http://example.com/"+s, nil)
		w := httptest.NewRecorder()
		fc.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		if string(body) != fmt.Sprintf(mockFormat, http.MethodGet, "/") {
			t.Errorf("Wanted %s but got %s", fmt.Sprintf(mockFormat, http.MethodGet, "/"), string(body))
		}
	})

	t.Run("when path not registered", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "http://example.com/notregistered", nil)
		w := httptest.NewRecorder()
		fc.ServeHTTP(w, r)
		resp := w.Result()
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Wanted %d but got %d", http.StatusNotFound, resp.StatusCode)
		}
	})
}

func TestNewServer(t *testing.T) {
	t.Run("when port is empty", func(t *testing.T) {
		x, err := NewServer(0, nil)
		if x != nil {
			t.Errorf("Wanted %v but got %v", nil, x)
		}
		if err != errNoPort {
			t.Errorf("Wanted %v but got %v", errNoPort, err)
		}
	})

	t.Run("happy path", func(t *testing.T) {
		x, _ := NewServer(8080, &FrontController{nil})
		if x == nil {
			t.Error("server not returned")
		}
	})
}
