package data

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/breadinator/lolapi/lolapi/endpoints"
)

// Data gotten from endpoints.GetChampionsAll and endpoints.GetChampionsSpecific.
type Champions struct {
	Type    string
	Format  string
	Version string
	Data    map[string]Champion // All the champions. Maps their name (e.g. "Aatrox") to their (incomplete) champion data.
}

// Champion data gotten from endpoints.GetChampionsAll and endpoints.GetChampionsSpecific.
//
// GetChampionsSpecific will retrieve all data, whereas GetChampionsAll gets just some of the data.
//
// https://developer.riotgames.com/docs/lol#data-dragon_champions explains how to enterpret their string templates
type Champion struct {
	Version   string         // Game version, e.g. "12.3.1"
	ID        string         // Champion ID, e.g. "Aatrox"
	Key       string         // Champion key, e.g. "266"
	Name      string         // Champion name, e.g. "Aatrox"
	Title     string         // Champion title, e.g. "the Darking Blade"
	Lore      string         // Full version of blurb
	Blurb     string         // Champion blurb, abridged
	Info      ChampionInfo   // Champion info, overview data on attack, defense, magic and difficulty
	Image     ChampionImage  // Info on champion image
	Skins     []ChampionSkin // Data about champion skins
	Tags      []string       // E.g. "Fighter", "Tank"
	AllyTips  []string       // Tips for helping this champion or having good synergy with them
	EnemyTips []string       // Tips for how to play against this champion
	Partype   string         // E.g. "Blood Well"
	Stats     ChampionStats
	Spells    []ChampionSpell
	Passive   struct {
		Name, Description string
		Image             ChampionImage
	}
	Recommended []interface{} // TODO check type
}

// Image of a champion, spell, etc
type ChampionImage struct {
	Full       string // E.g. "Aatrox.png"
	Sprite     string // E.g. "champion0.png"
	Group      string // E.g. "champion", "spell", "passive"
	X, Y, W, H int
}

// Some info on a champion.
type ChampionInfo struct {
	Attack, Defense, Magic, Difficulty int
}

type ChampionSkin struct {
	ID      string
	Num     int
	Name    string
	Chromas bool
}

type ChampionStats struct {
	HP, HPPerLevel, HPRegen, HPRegenPerLevel                                                              float32
	MP, MPPerLevel, MPRegen, MPRegenPerLevel                                                              float32
	MoveSpeed                                                                                             float32
	AttackRange, Crit, CritPerLevel, AttackDamage, AttackDamagePerLevel, AttackSpeed, AttackSpeedPerLevel float32
	Armor, ArmorPerLevel, SpellBlock, SpellBlockPerLevel                                                  float32
}

type ChampionSpell struct {
	ID           string
	Name         string
	Description  string
	LevelTip     struct{ Label, Effect []string }
	ToolTip      string
	MaxRank      int
	Cooldown     []float32 // TODO: check if int or float
	CooldownBurn string
	Cost         []int // TODO: check if int or float
	CostBurn     string
	CostType     string // E.g. " {{ abilityresourcename }}"
	MaxAmmo      string
	Range        []int // ? could be a float?
	RangeBurn    string
	Image        ChampionImage
	Resource     string
	Vars         []interface{} // TODO: check type
	Effect       [][]float32
	EffectBurn   []string
	DataValues   map[interface{}]interface{} // TODO: check type
}

// Downloads an image.
//
// Leaving the version param empty (i.e. "") will get the latest.
//
// Returns the URL, the file's bytes and an error.
func (i *ChampionImage) Get(version string) (string, *[]byte, error) {
	v := version
	if version == "" {
		v, _ = endpoints.GetLatestVersion()
	}

	var url string

	switch i.Group {
	case "champion":
		url = fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/champion/%s", v, i.Full)
	case "spell":
		url = fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/spell/%s", v, i.Full)
	case "passive":
		url = fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/passive/%s", v, i.Full)
	case "item":
		url = fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/item/%s", v, i.Full)
	default:
		return url, nil, fmt.Errorf("invalid image group %s", i.Group)
	}

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return url, nil, err
	}
	resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	return url, &data, err
}
