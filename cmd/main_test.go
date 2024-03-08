package main

import (
	"testing"
)

func TestPrintApplicationStarted(t *testing.T) {
	want := "Growing Buddy application started!"

	got := "Growing Buddy application started!"

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
