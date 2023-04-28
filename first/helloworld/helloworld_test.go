package helloworld_test

import (
	"first/helloworld"
	"testing"
)

func TestPrintHello(t *testing.T) {
	want := "Hello"
	if got := helloworld.PrintHello(); got != want {
		t.Errorf("PrintHello() = %q, want %q", got, want)
	}
}

func TestPrintWorld(t *testing.T) {
	want := "World!"
	if got := helloworld.PrintWorld(); got != want {
		t.Errorf("PrintWorld() = %q, want %q", got, want)
	}
}
