package clientapi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
These won't succeed unless in a game
*/

func TestGetActivePlayer(t *testing.T) {
	res, err := GetActivePlayer()
	assert.Nil(t, err)
	assert.False(t, res.SummonerName == "")
	assert.False(t, res.FullRunes.PrimaryRuneTree.DisplayName == "")
	fmt.Printf("%+#v\n", res)
}

func TestGetActivePlayerAbilities(t *testing.T) {
	res, err := GetActivePlayerAbilities()
	assert.Nil(t, err)
	assert.False(t, res.Passive.DisplayName == "")
	fmt.Printf("%+#v\n", res)
}

func TestGetActivePlayerName(t *testing.T) {
	res, err := GetActivePlayerName()
	assert.Nil(t, err)
	assert.False(t, *res == "")
	fmt.Println(*res)
}

func TestGetActivePlayerRunes(t *testing.T) {
	res, err := GetActivePlayerRunes()
	assert.Nil(t, err)
	assert.False(t, res.Keystone.DisplayName == "")
	fmt.Printf("%+#v\n", res)
}

func TestGetAllGameData(t *testing.T) {
	res, err := GetAllGameData()
	assert.Nil(t, err)
	fmt.Printf("%+#v\n", *res)
}

func TestGetEventData(t *testing.T) {
	res, err := GetEventData()
	assert.Nil(t, err)
	assert.True(t, res.Events[0].EventName == "GameStart")
	fmt.Printf("%+#v\n", *res)
}

func TestGetGameStats(t *testing.T) {
	res, err := GetGameStats()
	assert.Nil(t, err)
	assert.False(t, res.GameMode == "")
	fmt.Printf("%+#v\n", *res)
}

func TestGetPlayerItems(t *testing.T) {
	players, _ := GetPlayerList()
	res, err := GetPlayerItems((*players)[0].SummonerName)
	assert.Nil(t, err)
	fmt.Printf("%+#v\n", *res)
}

func TestGetPlayerList(t *testing.T) {
	res, err := GetPlayerList()
	assert.Nil(t, err)
	assert.False(t, len(*res) == 0)
	assert.False(t, (*res)[0].ChampionName == "")
	fmt.Printf("%+#v\n", *res)
}

func TestGetPlayerMainRunes(t *testing.T) {
	players, _ := GetPlayerList()
	res, err := GetPlayerMainRunes((*players)[0].SummonerName)
	assert.Nil(t, err)
	assert.False(t, res.Keystone.DisplayName == "")
	assert.False(t, res.PrimaryRuneTree.DisplayName == "")
	assert.False(t, res.SecondaryRuneTree.DisplayName == "")
	fmt.Printf("%+#v\n", *res)
}

func TestGetPlayerScores(t *testing.T) {
	players, _ := GetPlayerList()
	res, err := GetPlayerScores((*players)[0].SummonerName)
	assert.Nil(t, err)
	fmt.Printf("%+#v\n", *res)
}

func TestGetPlayerSummonerSpells(t *testing.T) {
	players, _ := GetPlayerList()
	res, err := GetPlayerSummonerSpells((*players)[0].SummonerName)
	assert.Nil(t, err)
	assert.False(t, res.SummonerSpellOne.DisplayName == "")
	assert.False(t, res.SummonerSpellOne.RawDescription == "")
	assert.False(t, res.SummonerSpellOne.RawDisplayName == "")
	assert.False(t, res.SummonerSpellTwo.DisplayName == "")
	assert.False(t, res.SummonerSpellTwo.RawDescription == "")
	assert.False(t, res.SummonerSpellTwo.RawDisplayName == "")
	fmt.Printf("%+#v\n", *res)
}
