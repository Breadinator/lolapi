package endpoints

import (
	"net/http"

	"github.com/breadinator/lolapi/lolapi"
)

// Returns data on in-game items.
//
// Args: (1) game version e.g. "12.3.1", (2) localisation e.g. "en_US"
var GetItems lolapi.Endpoint = lolapi.Endpoint{
	Method:  http.MethodGet,
	Path:    "http://ddragon.leagueoflegends.com/cdn/$1/data/$2/item.json",
	RawPath: true,
}
