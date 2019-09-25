// package equipinfo provides an in memory database for equipment information.

package equipinfo

import (
	"encoding/json"
	"github.com/kyoukaya/hoxy/log"
	"github.com/kyoukaya/hoxy/utils"
	"io/ioutil"
	"os"
	"sync"
)

// EquipInfo contains information about a specific equipment type.
type EquipInfo struct {
	ID         int `json:"id"`
	Type       int `json:"type"`
	Rarity     int `json:"rarity"`
	Fp         int `json:"fp"`
	Acc        int `json:"acc"`
	Eva        int `json:"eva"`
	Movespeed  int `json:"movespeed"`
	Rof        int `json:"rof"`
	Critdmg    int `json:"critdmg"`
	Crit       int `json:"crit"`
	Ap         int `json:"ap"`
	Armor      int `json:"armor"`
	Nightview  int `json:"nightview"`
	Rounds     int `json:"rounds"`
	LevelBonus struct {
		Fp        int `json:"fp"`
		Acc       int `json:"acc"`
		Eva       int `json:"eva"`
		Movespeed int `json:"movespeed"`
		Rof       int `json:"rof"`
		Critdmg   int `json:"critdmg"`
		Crit      int `json:"crit"`
		Ap        int `json:"ap"`
		Armor     int `json:"armor"`
		Nightview int `json:"nightview"`
		Rounds    int `json:"rounds"`
	} `json:"level_bonus"`
	EnCraftable   bool   `json:"en_craftable"`
	EnReleased    bool   `json:"en_released"`
	ConstructTime int    `json:"construct_time"`
	Name          string `json:"name"`
	Tooltip       string `json:"tooltip"`
}

var (
	equipMap     map[int]*EquipInfo
	equipMapLock sync.Mutex
	initOnce     sync.Once
)

// Init initializes the in memory equipment information. This should be called by your module
// if such information is required.
func Init() {
	initOnce.Do(initEquipInfo)
}

// Get EquipInfo based on Index ID.
func Get(id int) *EquipInfo {
	equipMapLock.Lock()
	defer equipMapLock.Unlock()
	return equipMap[id]
}

// EquipType returns a string representation of an equipment's type.
func (info *EquipInfo) EquipType() string {
	switch info.Type {
	case 1:
		return "Scope"
	case 2:
		return "EOT"
	case 3:
		return "RDS"
	case 4:
		return "PEQ"
	case 5:
		return "AP"
	case 6:
		return "HP"
	case 7:
		return "Slug"
	case 8:
		return "HV"
	case 9:
		return "Buckshot"
	case 10:
		return "X-Exo"
	case 11:
		return "Armor"
	case 12:
		return "T-Exo"
	case 13:
		return "Suppressor"
	case 14:
		return "Ammo Box"
	case 15:
		return "Cape"
	default:
		return "Speq"
	}
}

// initEquipInfo loads doll information from the dolls.json file
// equips.json from https://raw.githubusercontent.com/umang-p/brainlets/master/brainlets/girlsfrontline/static/girlsfrontline/equips.json
func initEquipInfo() {
	log.Infof("Initializing equipment information")
	f, err := os.Open(utils.PackageRoot + "/data/equips.json")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var infoSlice []*EquipInfo
	err = json.Unmarshal(data, &infoSlice)
	if err != nil {
		log.Fatal(err)
	}

	equipMap = make(map[int]*EquipInfo)
	for k, v := range infoSlice {
		equipMap[v.ID] = infoSlice[k]
	}
}
