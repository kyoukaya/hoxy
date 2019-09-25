package equipinfo_test

import (
	"hoxy/utils/equipinfo"
	"testing"
)

func TestEquipinfo(t *testing.T) {
	equipinfo.Init()
	equip := equipinfo.Get(87)
	if equip == nil {
		t.Fatal("Unable to find equip")
	}
	if equip.Name != "BM 3-12X40" {
		t.Errorf("Expected equip name \"BM 3-12X40\", got %s", equip.Name)
	}
}
