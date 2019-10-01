package constructioninfo

import (
	"fmt"
	"strings"

	"github.com/kyoukaya/hoxy/defs"
	"github.com/kyoukaya/hoxy/utils/equipinfo"
	"github.com/kyoukaya/hoxy/utils/fairyinfo"
)

func handleSingleEquip(sdata *defs.SEquipDevelop, cdata *defs.CEquipDevelop) (string, string, error) {
	rscUsed := fmt.Sprintf("[%d/%d/%d/%d]", cdata.Mp, cdata.Ammo, cdata.Mre, cdata.Part)
	equip := equipinfo.Get(sdata.EquipID)
	if equip == nil {
		fairy := fairyinfo.Get(sdata.FairyID)
		// TODO: Document PassiveSkill
		return rscUsed, fmt.Sprintf("%s", fairy.Name), nil
	}
	return rscUsed, fmt.Sprintf("%d* %s", equip.Rarity, equip.EquipType()), nil
}

func handleMultiEquip(sdata *defs.SEquipDevelopMulti, cdata *defs.CEquipDevelopMulti) (string, string, error) {
	var equipNames []string
	for _, v := range sdata.EquipIds {
		equip := equipinfo.Get(v.Info.EquipID)
		var s string
		if equip == nil {
			fairy := fairyinfo.Get(v.Info.FairyID)
			s = fmt.Sprintf("%s(%d)", fairy.Name, v.Info.PassiveSkill)
		} else {
			s = fmt.Sprintf("%d* %s", equip.Rarity, equip.EquipType())
		}
		equipNames = append(equipNames, s)
	}
	rscUsed := fmt.Sprintf("[%d/%d/%d/%d]", cdata.Mp, cdata.Ammo, cdata.Mre, cdata.Part)
	return rscUsed, strings.Join(equipNames, ","), nil
}
