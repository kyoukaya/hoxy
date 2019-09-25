package constructioninfo

import (
	"fmt"
	"github.com/kyoukaya/hoxy/defs"
	"github.com/kyoukaya/hoxy/utils/dollinfo"
	"strings"
)

func handleSingleGun(sdata *defs.SGunDevelopGun, cdata *defs.CGunDevelopGun) (string, string, error) {
	gunID := sdata.GunID
	rscUsed := fmt.Sprintf("[%d/%d/%d/%d]", cdata.Mp, cdata.Ammo, cdata.Mre, cdata.Part)
	gun := dollinfo.Get(gunID)
	return rscUsed, fmt.Sprintf("%d* %s %s", gun.Rarity, gun.DollType(), gun.Name), nil
}

func handleMultiGun(sdata *defs.SGunDevelopMultiGun, cdata *defs.CGunDevelopMultiGun) (string, string, error) {
	var gunNames []string
	for _, v := range sdata.GunIds {
		gunID := v.ID
		gun := dollinfo.Get(gunID)
		gunNames = append(gunNames, fmt.Sprintf("%d* %s %s", gun.Rarity, gun.DollType(), gun.Name))
	}
	rscUsed := fmt.Sprintf("[%d/%d/%d/%d]", cdata.Mp, cdata.Ammo, cdata.Mre, cdata.Part)
	return rscUsed, strings.Join(gunNames, ","), nil
}
