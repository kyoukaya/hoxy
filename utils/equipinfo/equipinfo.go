// package equipinfo provides an in memory database for equipment information.

package equipinfo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/kyoukaya/hoxy/log"
	"github.com/kyoukaya/hoxy/utils"
)

// EquipInfo contains information about a specific equipment type.
type EquipInfo struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	KrName    string `json:"krName"`
	Name      string `json:"enName"`
	JpName    string `json:"jpName"`
	Rarity    int    `json:"rank"`
	Type      int    `json:"type"`
	Category  string `json:"category"`
	BuildTime int    `json:"buildTime"`
	Maxlevel  int    `json:"maxlevel"`
	Company   string `json:"company"`
	Fitgun    []int  `json:"fitgun,omitempty"`
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

// initEquipInfo loads equipment information from the equips.json file
// equips.json from https://raw.githubusercontent.com/jinsung0907/GirlsFrontline-DB/master/data/json/equip.json
// Used to use the information from brainlets.moe but the IDs they use seem to be arbitrary.
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
