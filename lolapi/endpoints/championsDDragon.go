package endpoints

import (
	"net/http"

	"github.com/breadinator/lolapi/lolapi"
)

// Get a basic list of all champions. Doesn't require any keys.
//
// Args: (1) game version, (2) locale e.g. "en_US"
//
// https://ddragon.leagueoflegends.com/cdn/12.3.1/data/en_US/champion.json
var GetChampionsAll lolapi.Endpoint = lolapi.Endpoint{
	Method:  http.MethodGet,
	Path:    "http://ddragon.leagueoflegends.com/cdn/$1/data/$2/champion.json",
	RawPath: true,
}

// Get a specific champion. Doesn't require any keys.
//
// Arg: (1) game version, (2) locale e.g. "en_US", (3) champion name, e.g. "Aatrox"
//
// E.g. https://ddragon.leagueoflegends.com/cdn/12.3.1/data/en_US/champion/Aatrox.json
var GetChampionsSpecific lolapi.Endpoint = lolapi.Endpoint{
	Method:  http.MethodGet,
	Path:    "http://ddragon.leagueoflegends.com/cdn/$1/data/$2/champion/$3.json",
	RawPath: true,
}
