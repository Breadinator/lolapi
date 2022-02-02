package clientapi

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

const (
	EndpointActivePlayer          = "/liveclientdata/activeplayer"
	EndpointActivePlayerAbilities = "/liveclientdata/activeplayerabilities"
	EndpointActivePlayerName      = "/liveclientdata/activeplayername"
	EndpointActivePlayerRunes     = "/liveclientdata/activeplayerrunes"
	EndpointAllGameData           = "/liveclientdata/allgamedata"
	EndpointEventData             = "/liveclientdata/eventdata"
	EndpointGameStats             = "/liveclientdata/gamestats"
	EndpointPlayerItems           = "/liveclientdata/playeritems"
	EndpointPlayerList            = "/liveclientdata/playerlist"
	EndpointPlayerMainRunes       = "/liveclientdata/playermainrunes"
	EndpointPlayerScores          = "/liveclientdata/playerscores"
	EndpointPlayerSummonerSpells  = "/liveclientdata/playersummonerspells"
)

func GetActivePlayer() (*ActivePlayer, error) {
	body, err := GetRead(EndpointActivePlayer)
	if err != nil {
		return nil, err
	}

	var activePlayer ActivePlayer
	return &activePlayer, jsoniter.ConfigFastest.Unmarshal(body, &activePlayer)
}

func GetActivePlayerAbilities() (*Abilities, error) {
	body, err := GetRead(EndpointActivePlayerAbilities)
	if err != nil {
		return nil, err
	}

	var abilities Abilities
	return &abilities, jsoniter.ConfigFastest.Unmarshal(body, &abilities)
}

func GetActivePlayerName() (*string, error) {
	body, err := GetRead(EndpointActivePlayerName)
	if err != nil {
		return nil, err
	}

	name := string(body)
	if len(name) < 3 {
		return &name, fmt.Errorf("invalid name %s", name)
	}
	name = name[1 : len(name)-1]

	return &name, nil
}

func GetActivePlayerRunes() (*Runes, error) {
	body, err := GetRead(EndpointActivePlayerRunes)
	if err != nil {
		return nil, err
	}

	var runes Runes
	return &runes, jsoniter.ConfigFastest.Unmarshal(body, &runes)
}

func GetAllGameData() (*AllGameData, error) {
	body, err := GetRead(EndpointAllGameData)
	if err != nil {
		return nil, err
	}

	var gameData AllGameData
	return &gameData, jsoniter.ConfigFastest.Unmarshal(body, &gameData)
}

func GetEventData() (*Events, error) {
	body, err := GetRead(EndpointEventData)
	if err != nil {
		return nil, err
	}

	var eventData Events
	return &eventData, jsoniter.ConfigFastest.Unmarshal(body, &eventData)
}

func GetGameStats() (*GameStats, error) {
	body, err := GetRead(EndpointGameStats)
	if err != nil {
		return nil, err
	}

	var gameStats GameStats
	return &gameStats, jsoniter.ConfigFastest.Unmarshal(body, &gameStats)
}

func GetPlayerItems(playerName string) (*[]Item, error) {
	body, err := GetRead(EndpointPlayerItems + "?summonerName=" + playerName)
	if err != nil {
		return nil, err
	}

	var items []Item
	return &items, jsoniter.ConfigFastest.Unmarshal(body, &items)
}

func GetPlayerList() (*[]Player, error) {
	body, err := GetRead(EndpointPlayerList)
	if err != nil {
		return nil, err
	}

	var players []Player
	return &players, jsoniter.ConfigFastest.Unmarshal(body, &players)
}

func GetPlayerMainRunes(playerName string) (*Runes, error) {
	body, err := GetRead(EndpointPlayerMainRunes + "?summonerName=" + playerName)
	if err != nil {
		return nil, err
	}

	var runes Runes
	return &runes, jsoniter.ConfigFastest.Unmarshal(body, &runes)
}

func GetPlayerScores(playerName string) (*Scores, error) {
	body, err := GetRead(EndpointPlayerScores + "?summonerName=" + playerName)
	if err != nil {
		return nil, err
	}

	var scores Scores
	return &scores, jsoniter.ConfigFastest.Unmarshal(body, &scores)
}

func GetPlayerSummonerSpells(playerName string) (*struct{ SummonerSpellOne, SummonerSpellTwo DataEntry }, error) {
	body, err := GetRead(EndpointPlayerSummonerSpells + "?summonerName=" + playerName)
	if err != nil {
		return nil, err
	}

	var spells struct{ SummonerSpellOne, SummonerSpellTwo DataEntry }
	return &spells, jsoniter.ConfigFastest.Unmarshal(body, &spells)
}
