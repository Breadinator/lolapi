package endpoints

import (
	"net/http"

	api "github.com/breadinator/lolapi/lolapi"
)

// Gets an account by its PUUID.
//
// Argument 1 is the PUUID.
//
// https://developer.riotgames.com/apis#account-v1/GET_getByPuuid
var GetAccountByPUUID api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/riot/account/v1/accounts/by-puuid/$1",
	UseRoute: true,
}

// Gets an account by its Riot game name and tag line
//
// Argument 1 is the name (before the #) and argument 2 is the tag line (after the #)
//
// https://developer.riotgames.com/apis#account-v1/GET_getByRiotId
var GetAccountByRiotID api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/riot/account/v1/accounts/by-riot-id/$1/$2",
	UseRoute: true,
}

//
//
// Argument 1 is the game and argument 2 is the PUUID
//
// https://developer.riotgames.com/apis#account-v1/GET_getActiveShard
var GetActiveShards api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/riot/account/v1/active-shards/by-game/$1/by-puuid/$2",
	UseRoute: true,
}

// Gets an account based on the token making the call.
//
// Takes no arguments.
//
// https://developer.riotgames.com/apis#account-v1/GET_getByAccessToken
var GetAccountByAccessToken api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/riot/account/v1/accounts/me",
	UseRoute: true,
}
