package fairyinfo_test

import (
	"testing"

	"github.com/kyoukaya/hoxy/utils/fairyinfo"
)

func TestFairyInfo(t *testing.T) {
	fairyinfo.Init()
	fairy := fairyinfo.Get(8)
	if fairy == nil {
		t.Fatal("Unable to find fairy")
	}
	if fairy.Name != "Artillery Fairy" {
		t.Errorf("Expected equip name \"Artillery Fairy\", got \"%s\"", fairy.Name)
	}
}
