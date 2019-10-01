// package dollinfo provides an in memory database for doll information.

package dollinfo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/kyoukaya/hoxy/log"
	"github.com/kyoukaya/hoxy/utils"

	"github.com/mitchellh/mapstructure"
)

// DollInfo contains information about a specific doll type.
type DollInfo struct {
	ID             int      `json:"id"`
	IDIndex        int      `json:"id_index"`
	Aliases        []string `json:"aliases"`
	APIName        string   `json:"api_name"`
	SpritesheetRow int      `json:"spritesheet_row"`
	SpritesheetCol int      `json:"spritesheet_col"`
	Name           string   `json:"name"`
	Rarity         int      `json:"rarity"`
	Type           int      `json:"type"`
	Mod            bool     `json:"mod"`
	Hp             int      `json:"hp"`
	Fp             int      `json:"fp"`
	Acc            int      `json:"acc"`
	Eva            int      `json:"eva"`
	Rof            int      `json:"rof"`
	Crit           int      `json:"crit"`
	Critdmg        int      `json:"critdmg"`
	Ap             int      `json:"ap"`
	Rounds         int      `json:"rounds"`
	Armor          int      `json:"armor"`
	GrowthRating   int      `json:"growth_rating"`
	ConstructTime  int      `json:"construct_time"`
	EnCraftable    bool     `json:"en_craftable"`
	EnReleased     bool     `json:"en_released"`
	Artist         string   `json:"artist"`
	Voice          string   `json:"voice"`
	NameSkill1     string   `json:"name_skill1"`
	IconNameSkill1 string   `json:"icon_name_skill1"`
	TooltipSkill1  string   `json:"tooltip_skill1"`
	Skill          struct {
		Icd     int       `json:"icd"`
		Cd      []float64 `json:"cd"`
		Effects []struct {
			Type   string `json:"type"`
			Target string `json:"target"`
			Stat   struct {
				Fp []int `json:"fp"`
			} `json:"stat"`
			Duration []float64 `json:"duration"`
		} `json:"effects"`
	} `json:"skill"`
	TooltipSkill2 string `json:"tooltip_skill2"`
	TooltipTiles  string `json:"tooltip_tiles"`
	Tiles         struct {
		Self       int    `json:"self"`
		Target     string `json:"target"`
		TargetType []int  `json:"target_type"`
		Effect     struct {
			Fp      []int `json:"fp"`
			Acc     []int `json:"acc"`
			Eva     []int `json:"eva"`
			Rof     []int `json:"rof"`
			Crit    []int `json:"crit"`
			Skillcd []int `json:"skillcd"`
			Armor   []int `json:"armor"`
		} `json:"effect"`
	} `json:"tiles"`
}

var (
	infoMap     map[int]*DollInfo
	infoMapLock sync.Mutex
	initOnce    sync.Once
)

// Init initializes the in memory doll information. This should be called by your module
// when initializing for a user, if such information is required.
func Init() {
	initOnce.Do(initDollInfo)
}

// Get DollInfo based on Index ID.
func Get(id int) *DollInfo {
	infoMapLock.Lock()
	defer infoMapLock.Unlock()
	return infoMap[id]
}

// DollType returns a string representation of an doll's type.
func (info *DollInfo) DollType() string {
	switch info.Type {
	case 1:
		return "HG"
	case 2:
		return "SMG"
	case 3:
		return "RF"
	case 4:
		return "AR"
	case 5:
		return "MG"
	case 6:
		return "SG"
	default:
		return "error"
	}
}

// initDollInfo loads doll information from the dolls.json file
// dolls.json from https://raw.githubusercontent.com/umang-p/brainlets/master/brainlets/girlsfrontline/static/girlsfrontline/equips.json
func initDollInfo() {
	log.Infof("Initializing doll information")
	f, err := os.Open(utils.PackageRoot + "/data/dolls.json")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var m []interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
	}

	// The JSON data in tiles.effect are represented as single values or as arrays.
	// Using a mapstructure here to properly handle that looseness.
	var infoSlice []DollInfo
	conf := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &infoSlice,
		TagName:          "json"}
	dec, err := mapstructure.NewDecoder(conf)
	if err != nil {
		log.Fatal(err)
	}
	err = dec.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	infoMap = make(map[int]*DollInfo)
	for k, v := range infoSlice {
		// Hack: Don't overwrite existing keys to prioritize showing non-modded guns.
		if _, ok := infoMap[v.IDIndex]; !ok {
			infoMap[v.IDIndex] = &infoSlice[k]
		}
	}
}
