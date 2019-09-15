package dollinfo

import (
	"encoding/json"
	"hoxy/log"
	"hoxy/utils"
	"io/ioutil"
	"os"
	"sync"
)

// Dolls.json from http://brainlets.moe/static/girlsfrontline/dolls.json
// TODO: Pull data at runtime. Maybe pull it from a repo if possible so we don't
// have to put load on brainlets.moe
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
	// Artist         string   `json:"artist"`
	// Voice          string   `json:"voice"`
	// NameSkill1     string   `json:"name_skill1"`
	// IconNameSkill1 string   `json:"icon_name_skill1"`
	// TooltipSkill1  string   `json:"tooltip_skill1"`
	// Skill          struct {
	// 	Icd     int       `json:"icd"`
	// 	Cd      []float64 `json:"cd"`
	// 	Effects []struct {
	// 		Type   string `json:"type"`
	// 		Target string `json:"target"`
	// 		Stat   struct {
	// 			Fp []int `json:"fp"`
	// 		} `json:"stat"`
	// 		Duration []float64 `json:"duration"`
	// 	} `json:"effects"`
	// } `json:"skill"`
	// TooltipSkill2 string `json:"tooltip_skill2"`
	// TooltipTiles  string `json:"tooltip_tiles"`
	// Tiles         struct {
	// 	Self       int    `json:"self"`
	// 	Target     string `json:"target"`
	// 	TargetType int    `json:"target_type"`
	// 	Effect     struct {
	// 		Fp      []int `json:"fp"`
	// 		Acc     []int `json:"acc"`
	// 		Eva     []int `json:"eva"`
	// 		Rof     []int `json:"rof"`
	// 		Crit    []int `json:"crit"`
	// 		Skillcd []int `json:"skillcd"`
	// 		Armor   []int `json:"armor"`
	// 	} `json:"effect"`
	// } `json:"tiles"`
}

var (
	infoMap     map[int]*DollInfo
	infoMapLock *sync.Mutex
)

// InitDollInfo loads doll information from the dolls.json file
func InitDollInfo() {
	log.Infof("Initializing doll information")
	f, err := os.Open(utils.PackageRoot + "/utils/dollinfo/dolls.json")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var infoSlice []DollInfo
	err = json.Unmarshal(data, &infoSlice)
	if err != nil {
		log.Fatal(err)
	}

	infoMap = make(map[int]*DollInfo)
	infoMapLock = &sync.Mutex{}
	for k, v := range infoSlice {
		// Hack: Don't overwrite existing keys to prioritize showing non-modded guns.
		if _, ok := infoMap[v.IDIndex]; !ok {
			infoMap[v.IDIndex] = &infoSlice[k]
		}
	}
}

// Get DollInfo based on Index ID.
func Get(id int) *DollInfo {
	infoMapLock.Lock()
	defer infoMapLock.Unlock()
	return infoMap[id]
}
