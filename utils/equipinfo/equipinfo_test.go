package equipinfo_test

import (
	"testing"

	"github.com/kyoukaya/hoxy/utils/equipinfo"
)

func TestEquipinfo(t *testing.T) {
	equipinfo.Init()
	equip := equipinfo.Get(87)
	if equip == nil {
		t.Fatal("Unable to find equip")
	}
	if equip.Name != "Ragged Cape" {
		t.Errorf("Expected equip name \"BM 3-12X40\", got \"%s\"", equip.Name)
	}
}
