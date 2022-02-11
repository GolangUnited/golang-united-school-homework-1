package main

import (
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestPrintMain(t *testing.T) {
	ac_txt := PrintMain()
	ex_tx := emoji.Sprint("Hello :world_map:!")
	if ac_txt != ex_tx {
		t.Errorf(`Error text ` + ac_txt)
	}
}
