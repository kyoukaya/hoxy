package fairyinfo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/kyoukaya/hoxy/log"
	"github.com/kyoukaya/hoxy/utils"
	"github.com/mitchellh/mapstructure"
)

type FairyInfo struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Type          int    `json:"type"`
	Fp            int    `json:"fp"`
	Acc           int    `json:"acc"`
	Eva           int    `json:"eva"`
	Armor         int    `json:"armor"`
	Critdmg       int    `json:"critdmg"`
	GrowthRating  int    `json:"growth_rating"`
	ConstructTime int    `json:"construct_time"`
	EnReleased    bool   `json:"en_released"`
	EnCraftable   bool   `json:"en_craftable"`
	TooltipAura   string `json:"tooltip_aura"`
	TooltipSkill  string `json:"tooltip_skill"`
	Skill         struct {
		Icd     int `json:"icd"`
		Cd      int `json:"cd"`
		Effects []struct {
			Type   string `json:"type"`
			Target string `json:"target"`
			Stat   struct {
				Fp  []int `json:"fp"`
				Rof []int `json:"rof"`
			} `json:"stat"`
			Duration int `json:"duration"`
		} `json:"effects"`
	} `json:"skill"`
	SpecialControl bool `json:"special_control,omitempty"`
}

var (
	fairyMap     map[int]*FairyInfo
	fairyMapLock sync.Mutex
	initOnce     sync.Once
)

// Init initializes the in memory equipment information. This should be called by your module
// if such information is required.
func Init() {
	initOnce.Do(initFairyInfo)
}

// Get FairyInfo based on Index ID.
func Get(id int) *FairyInfo {
	fairyMapLock.Lock()
	defer fairyMapLock.Unlock()
	return fairyMap[id]
}

// initFairyInfo loads doll information from the fairy.json file
// fairy.json from https://raw.githubusercontent.com/umang-p/brainlets/master/brainlets/girlsfrontline/static/girlsfrontline/fairies.json
func initFairyInfo() {
	log.Infof("Initializing fairy information")
	f, err := os.Open(utils.PackageRoot + "/data/fairies.json")
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

	var infoSlice []*FairyInfo
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

	fairyMap = make(map[int]*FairyInfo)
	for k, v := range infoSlice {
		fairyMap[v.ID] = infoSlice[k]
	}
}
