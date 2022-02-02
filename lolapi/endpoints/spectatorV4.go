package endpoints

import (
	"net/http"

	api "github.com/breadinator/lolapi/lolapi"
)

// Gets the current game of a given summoner
//
// Arg: encrypted summoner ID.
//
// https://developer.riotgames.com/apis#spectator-v4/GET_getCurrentGameInfoBySummoner
var GetGamesBySummoner api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/spectator/v4/active-games/by-summoner/$1",
	UseRoute: false,
}

// Returns a list of featured games.
//
// https://developer.riotgames.com/apis#spectator-v4/GET_getFeaturedGames
var GetGamesFeatured api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/spectator/v4/featured-games",
	UseRoute: false,
}
