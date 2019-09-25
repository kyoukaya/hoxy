package constructioninfo

import (
	"fmt"
	"github.com/kyoukaya/hoxy/defs"
	"github.com/kyoukaya/hoxy/utils/equipinfo"
	"strings"
)

func handleSingleEquip(sdata *defs.SEquipDevelop, cdata *defs.CEquipDevelop) (string, string, error) {
	rscUsed := fmt.Sprintf("[%d/%d/%d/%d]", cdata.Mp, cdata.Ammo, cdata.Mre, cdata.Part)
	equipID := sdata.EquipID
	equip := equipinfo.Get(equipID)
	return rscUsed, fmt.Sprintf("%d* %s", equip.Rarity, equip.EquipType()), nil
}

func handleMultiEquip(sdata *defs.SEquipDevelopMulti, cdata *defs.CEquipDevelopMulti) (string, string, error) {
	var equipNames []string
	for _, v := range sdata.EquipIds {
		equipID := v.Info.EquipID
		equip := equipinfo.Get(equipID)
		equipNames = append(equipNames, fmt.Sprintf("%d* %s", equip.Rarity, equip.EquipType()))
	}
	rscUsed := fmt.Sprintf("[%d/%d/%d/%d]", cdata.Mp, cdata.Ammo, cdata.Mre, cdata.Part)
	return rscUsed, strings.Join(equipNames, ","), nil
}
