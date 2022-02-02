package endpoints

import (
	"net/http"

	api "github.com/breadinator/lolapi/lolapi"
)

// Gets the challenger league for a given queue.
//
// Argument 1 is the queue e.g. api.QueueLolRankedSolo.
//
// https://developer.riotgames.com/apis#league-v4/GET_getChallengerLeague
var GetLeagueChallenger api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/league/v4/challengerleagues/by-queue/$1",
	UseRoute: false,
}

//
//
// Argument 1 is the encrypted summoner ID.
//
// https://developer.riotgames.com/apis#league-v4/GET_getLeagueEntriesForSummoner
var GetLeagueBySummoner api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/league/v4/entries/by-summoner/$1",
	UseRoute: false,
}

// Get a rank divison for a given queue.
//
// Args: (1) queue, (2) tier, (3) division, (4) result page
//
// https://developer.riotgames.com/apis#league-v4/GET_getLeagueEntries
var GetLeagueAll api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/league/v4/entries/$1/$2/$3?page=$4",
	UseRoute: false,
}

// Gets the Grandmaster league of a given queue.
//
// Arg: queue e.g. QueueLolRankedSolo
//
// https://developer.riotgames.com/apis#league-v4/GET_getGrandmasterLeague
var GetLeagueGrandmaster api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/league/v4/grandmasterleagues/by-queue/$1",
	UseRoute: false,
}

// Get a league from a given league ID (including inactive entries).
//
// Arg: league ID
//
// https://developer.riotgames.com/apis#league-v4/GET_getLeagueById
var GetLeagueFromID api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/league/v4/leagues/$1",
	UseRoute: false,
}

// Gets the Master league of a given queue.
//
// Arg: queue e.g. QueueLolRankedSolo
//
// https://developer.riotgames.com/apis#league-v4/GET_getMasterLeague
var GetLeagueMaster api.Endpoint = api.Endpoint{
	Method:   http.MethodGet,
	Path:     "/lol/league/v4/masterleagues/by-queue/{queue}",
	UseRoute: false,
}
