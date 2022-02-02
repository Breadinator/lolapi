package endpoints

import (
	"net/http"

	api "github.com/breadinator/lolapi/lolapi"
)

// Gets a list of matches by a player's PUUID.
//
// Required arg: PUUID.
//
// Optional args:
//
// (1) Start time int: Epoch timestamp in seconds. The matchlist started storing timestamps on June 16th, 2021. Any matches played before June 16th, 2021 won't be included in the results if the startTime filter is set.
//
// (2) End time int: Epoch timestamp in seconds.
//
// (3) Queue int: Filter the list of match ids by a specific queue id. This filter is mutually inclusive of the type filter meaning any match ids returned must match both the queue and type filters.
//
// (4) Type string: Filter the list of match ids by the type of match. This filter is mutually inclusive of the queue filter meaning any match ids returned must match both the queue and type filters.
// "ranked", "normal", "tourney" or "tutorial"
//
// (5) Start int: Defaults to 0. Start index.
//
// (6) Count int: Defaults to 20. Valid values: 0 to 100. Number of match ids to return.
//
// https://developer.riotgames.com/apis#match-v5/GET_getMatchIdsByPUUID
var GetMatchesFromPUUID api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/match/v5/matches/by-puuid/$1/ids?startTime=$2&endTime=$3&queue=$4&type=$5&start=$6&count=$7",
	UseRoute: true,
}

// Fetch a match by its ID.
//
// Arg: match ID.
//
// https://developer.riotgames.com/apis#match-v5/GET_getMatch
var GetMatchByID api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/match/v5/matches/$1",
	UseRoute: true,
}

// Fetch a match timeline by its ID.
//
// Arg: match ID.
//
// https://developer.riotgames.com/apis#match-v5/GET_getTimeline
var GetMatchTimeline api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/match/v5/matches/$1/timeline",
	UseRoute: true,
}
