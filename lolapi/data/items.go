package data

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/breadinator/lolapi/lolapi/endpoints"
	jsoniter "github.com/json-iterator/go"
)

type Items struct {
	Type    string                                 // "items"
	Version string                                 // Game version
	Basic   Item                                   // Empty item template
	Data    map[int]Item                           // Actual items
	Groups  []struct{ ID, MaxGroupOwnable string } // Defines various item groups
	Tree    []struct {                             // Categorises tags
		Header string
		Tags   []string
	}
}

type Item struct {
	Name string
	Rune struct {
		IsRune bool
		Tier   int
		Type   string
	}
	Gold struct {
		Base, Total, Sell int
		Purchasable       bool
	}
	Group            string
	Description      string
	Colloq           string
	PlainText        string
	Consumed         bool
	Stacks           int
	Depth            int
	ConsumeOnFull    bool
	From             []string // TODO: verify type
	Into             []string // TODO: verify type
	SpecialRecipe    int
	InStore          bool
	HideFromAll      bool
	RequiredChampion string
	RequiredAlly     string
	Stats            map[string]float64
	Tags             []string // TODO: verify type
	Maps             map[int]bool
	Image            ChampionImage
}

func GetItemsParsed(version, localisation string) (*Items, error) {
	items := new(Items)

	url := endpoints.GetItems.Path
	url = strings.ReplaceAll(url, "$1", version)
	url = strings.ReplaceAll(url, "$2", localisation)

	resp, err := http.Get(url)
	if err != nil {
		return items, err
	}
	if resp.StatusCode != 200 {
		return items, fmt.Errorf("couldn't get items, %s", resp.Status)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return items, err
	}

	return items, jsoniter.ConfigFastest.Unmarshal(data, items)
}
