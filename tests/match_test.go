package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/breadinator/lolapi/lolapi"
	"github.com/breadinator/lolapi/lolapi/data"
	"github.com/breadinator/lolapi/lolapi/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestGetMatchesByPUUID(t *testing.T) {
	conf, err := lolapi.GetConfig()
	assert.Nil(t, err)
	client := lolapi.ClientFromConfig(conf, lolapi.RegionEuropeWest)
	client.Limit = 0
	client.Region = lolapi.RegionEuropeWest
	resp, _, err := client.Call(endpoints.GetSummonerByName, "good bye iceland")
	assert.Nil(t, err)

	var summoner data.SummonerDTO
	json.Unmarshal(resp, &summoner)
	assert.NotEqual(t, "", summoner.PUUID)

	client.AwaitNoLimit()

	resp, _, err = client.Call(endpoints.GetMatchesFromPUUID, summoner.PUUID)
	fmt.Println(string(resp))
	assert.Nil(t, err)
	var matches []string
	json.Unmarshal(resp, &matches)
	fmt.Println(matches)
}

func TestGetMatchByID(t *testing.T) {
	conf, err := lolapi.GetConfig()
	assert.Nil(t, err)
	client := lolapi.ClientFromConfig(conf, lolapi.RegionEuropeWest)
	client.Limit = 0
	client.Region = lolapi.RegionEuropeWest
	resp, _, err := client.Call(endpoints.GetMatchByID, "EUW1_5283033444")
	assert.Nil(t, err)

	matchData := data.GetMatchDTO(resp)
	fmt.Printf("%#v\n", matchData)

	participants := []string{
		"5FqT0QB2yWgokyLCUGFXgBJm7zdFYdtmlnDfQmgLAO9kF8VICwn60vr8Bi7WnJVI97lVGPYG5MV7uw",
		"dZbDi-6G0Atmc8CM85RqGl8Rv2t-sG8AGNsKf422bFuba3GFcMngYdMfl4DXAVvC_cp9TBB0utDxpg",
		"nJXiqvwVW8b_fih02Cuh34WI4a8zFeTi6vtL0x711PoIZeQHhrUuHSTiY9BPBvkjBuj_ikhdTX-H-w",
		"Ermk-6p142Bdn3RNDvDURk4qofzxWMaMmXy9lJDy_96MN4quBN5nQgFq9wH2Z9mv3YufmbfzcUoN5Q",
		"njWb15QMDew9YM1XOVeEPBRlW5B3t4VFlzvkoe489WqRm6OY4YQeoej3mLqOafPpRh_DkTrYy1-UkA",
		"ofL1e4L-vDIYKotJgHq3_OPXjLk7N3rfxW5K5Zt6ydqR84kYb-0Y_Fkldhw-Z1JO8PKcwojftf36jg",
		"hYUGG51BJMMfVfK5edHQ7YB0sGVQ3eWq7vHhAkReMhETk_ypycFxhePjr5TufoO7tX3tNVtnEusKZg",
		"fMUQz6jU76sIlmS7v83S9n_bKLGqxqlzl2kEx0TB80s3Jf15ZCif5sM3Ct_d7SZrwD0P8ko3OlLa7g",
		"FxM_xsxCyWxpE9uGgRnuXlZprW2STOnpdF8R3-4qzOqvrPxRVDSZ_n9qchure0djWUOEpkRnRAwH3g",
		"GKQwIoqT0r-PKOAuK4nr4TK71P4qOfRscP4I1_FyD3QIRqagNgnPPtH0q2TVlBEZgx-9kPTNtOjqwg",
	}

	for _, participant := range participants {
		assert.Contains(t, matchData.Metadata.Participants, participant)
	}

	assert.EqualValues(t, 1621727385000, matchData.Info.GameCreation)
	assert.EqualValues(t, 1492261, matchData.Info.GameDuration)
	assert.EqualValues(t, 5283033444, matchData.Info.GameID)
	assert.Equal(t, "CLASSIC", matchData.Info.GameMode)
	assert.Equal(t, "teambuilder-match-5283033444", matchData.Info.GameName)
	assert.EqualValues(t, 1621727422082, matchData.Info.GameStartTimestamp)
	assert.Equal(t, "MATCHED_GAME", matchData.Info.GameType)
	assert.Equal(t, "11.10.376.4811", matchData.Info.GameVersion)
	assert.EqualValues(t, 11, matchData.Info.MapID)

	for _, participant := range matchData.Info.Participants {
		assert.Contains(t, participants, participant.PUUID)
	}

	assert.Equal(t, "EUW1", matchData.Info.PlatformID)
	assert.EqualValues(t, 420, matchData.Info.QueueID)
}
