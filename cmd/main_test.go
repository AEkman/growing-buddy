package main

import (
	"testing"
)

func TestPrintApplicationStarted(t *testing.T) {
	want := "Growing Buddy application started!"

	got := printApplicationStarted()

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
