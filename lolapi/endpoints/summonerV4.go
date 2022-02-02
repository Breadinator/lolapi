package endpoints

import (
	"net/http"

	api "github.com/breadinator/lolapi/lolapi"
)

// Gets a summoner from their encrypted account ID (region-specific).
//
// Arg: encrypted account ID.
//
// https://developer.riotgames.com/apis#summoner-v4/GET_getByAccountId
var GetSummonerByAccountID api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/summoner/v4/summoners/by-account/$1",
	UseRoute: false,
}

// Gets a summoner by their summoner name (region-specific).
//
// Arg: summoner name.
//
// https://developer.riotgames.com/apis#summoner-v4/GET_getBySummonerName
var GetSummonerByName api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/summoner/v4/summoners/by-name/$1",
	UseRoute: false,
}

// Gets a summoner by their encrypted PUUID.
//
// Arg: encrypted PUUID.
//
// https://developer.riotgames.com/apis#summoner-v4/GET_getByPUUID
var GetSummonerByPUUID api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/summoner/v4/summoners/by-puuid/$1",
	UseRoute: false,
}

// Gets a summoner by their encrypted summoner ID.
//
// Arg: encrypted summoner ID.
//
// https://developer.riotgames.com/apis#summoner-v4/GET_getBySummonerId
var GetSummonerBySummonerId api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/summoner/v4/summoners/$1",
	UseRoute: false,
}

// Gets the summoner associated with the API key.
//
// https://developer.riotgames.com/apis#summoner-v4/GET_getByAccessToken
var GetSummonerMe api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/summoner/v4/summoners/me",
	UseRoute: false,
}
