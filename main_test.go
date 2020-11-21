package main

import "testing"

func TestFunctions(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Having go bind9 zone format", func(t *testing.T) {
		got := bind9ZonesFormat()
		want := "zone \"%s\" {type master; file \"/etc/bind/blacklisted.db\";};\n"

		assertCorrectMessage(t, got, want)
	})
}