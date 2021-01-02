package main

import "testing"

func TestPageCreation(t *testing.T) {
	t.Run("it is created with a page name", func(t *testing.T) {
		page := Page{Name: "index"}
		if page.Name != "index" {
			t.Errorf("got page name %q wanted %q", page.Name, "index")
		}
	})
	t.Run("created with visit count of 0", func(t *testing.T) {
		page := Page{Name: "index"}
		if page.Visits != 0 {
			t.Errorf("number of vists got %d wanted %d", page.Visits, 0)
		}
	})
}
