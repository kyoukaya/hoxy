package dollinfo_test

import (
	"github.com/kyoukaya/hoxy/utils/dollinfo"
	"testing"
)

func TestDollinfo(t *testing.T) {
	dollinfo.Init()
	saa := dollinfo.Get(1)
	if saa == nil {
		t.Fatal("Couldn't find doll")
	}
	if saa.Name != "SAA" {
		t.Errorf("Expected SAA, got %s", saa.Name)
	}
}
