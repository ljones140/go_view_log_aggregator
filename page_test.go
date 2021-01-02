package main

import (
	"net"
	"testing"
)

func TestPageCreation(t *testing.T) {
	t.Run("it is created with a page name", func(t *testing.T) {
		page := Page{Name: "index"}
		if page.Name != "index" {
			t.Errorf("got page name %q wanted %q", page.Name, "index")
		}
	})
	t.Run("created with visit count of 0", func(t *testing.T) {
		page := Page{Name: "index"}
		assertPageVisits(t, page.Visits, 0)
	})
}

func TestAddingPageVisits(t *testing.T) {
	t.Run("visit count is incremented when visit is added", func(t *testing.T) {
		page := Page{Name: "index"}
		address := net.ParseIP("192.0.2.1")
		page.AddVisit(address)

		assertPageVisits(t, page.Visits, 1)
	})
}

func assertPageVisits(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("number of visits got %d wanted %d", got, want)
	}
}
