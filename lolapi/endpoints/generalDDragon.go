package endpoints

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/breadinator/lolapi/lolapi"
	jsoniter "github.com/json-iterator/go"
)

// List of versions. Endpoint gives a list of strings. First in the list is most recent.
//
// https://ddragon.leagueoflegends.com/api/versions.json
var GetVersions lolapi.Endpoint = lolapi.Endpoint{
	Method:  http.MethodGet,
	Path:    "https://ddragon.leagueoflegends.com/api/versions.json",
	RawPath: true,
}

func GetLatestVersion() (string, error) {
	resp, err := http.Get(GetVersions.Path)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	versions := make([]string, 400)
	err = jsoniter.Unmarshal(data, &versions)
	if err != nil {
		return "", err
	}

	if len(versions) == 0 {
		return "", errors.New("invalid index")
	}

	return versions[0], err
}

// Gets data for a region, e.g. versions, localisation code, ddragon CDN.
//
// Arg 1: the region code, e.g. "euw", "na"
//
// https://ddragon.leagueoflegends.com/realms/euw.json
var GetRegion lolapi.Endpoint = lolapi.Endpoint{
	Method:  http.MethodGet,
	Path:    "https://ddragon.leagueoflegends.com/realms/$1.json",
	RawPath: true,
}

// Returns an array of strings of language codes, e.g. "en_US", "zh_CN", etc.
//
// https://ddragon.leagueoflegends.com/cdn/languages.json
var GetLanguages lolapi.Endpoint = lolapi.Endpoint{
	Method:  http.MethodGet,
	Path:    "https://ddragon.leagueoflegends.com/cdn/languages.json",
	RawPath: true,
}

/*
TODO
https://ddragon.leagueoflegends.com/cdn/12.3.1/data/en_US/summoner.json
http://ddragon.leagueoflegends.com/cdn/12.3.1/data/en_US/profileicon.json
minimaps https://developer.riotgames.com/docs/lol#general_map-names
icons
	http://ddragon.leagueoflegends.com/cdn/5.5.1/img/ui/champion.png
	http://ddragon.leagueoflegends.com/cdn/5.5.1/img/ui/items.png
	http://ddragon.leagueoflegends.com/cdn/5.5.1/img/ui/minion.png
	http://ddragon.leagueoflegends.com/cdn/5.5.1/img/ui/score.png
	http://ddragon.leagueoflegends.com/cdn/5.5.1/img/ui/spells.png
*/
