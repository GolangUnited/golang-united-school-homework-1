package main

import (
	"testing"
	"github.com/kyokomi/emoji/v2"
)

func TestHello(t *testing.T) {
	got := GetMessage()()
	want := ("Hello :world_map" !")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
