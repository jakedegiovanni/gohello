package urlutil

import "testing"

func TestShiftPath(t *testing.T) {
	t.Run("empty path", func(t *testing.T) {
		head, tail := ShiftPath("")
		if head != "" {
			t.Errorf("Wanted %s but got %s", "", head)
		}
		if tail != "/" {
			t.Errorf("wanted %s but got %s", "/", tail)
		}
	})

	t.Run("root path", func(t *testing.T) {
		head, tail := ShiftPath("/")
		if head != "" {
			t.Errorf("Wanted %s but got %s", "", head)
		}
		if tail != "/" {
			t.Errorf("wanted %s but got %s", "/", tail)
		}
	})

	t.Run("single level path", func(t *testing.T) {
		head, tail := ShiftPath("/hello")
		if head != "hello" {
			t.Errorf("Wanted %s but got %s", "hello", head)
		}
		if tail != "/" {
			t.Errorf("wanted %s but got %s", "/", tail)
		}
	})

	t.Run("multi level path", func(t *testing.T) {
		head, tail := ShiftPath("/hello/world")
		if head != "hello" {
			t.Errorf("Wanted %s but got %s", "hello", head)
		}
		if tail != "/world" {
			t.Errorf("wanted %s but got %s", "/world", tail)
		}
	})

	t.Run("multi level path trailing slash", func(t *testing.T) {
		head, tail := ShiftPath("/hello/world/")
		if head != "hello" {
			t.Errorf("Wanted %s but got %s", "hello", head)
		}
		if tail != "/world" {
			t.Errorf("wanted %s but got %s", "/world", tail)
		}
	})
}
