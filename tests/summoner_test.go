package tests

import (
	"fmt"
	"testing"

	"github.com/breadinator/lolapi/lolapi"
	"github.com/breadinator/lolapi/lolapi/data"
	"github.com/breadinator/lolapi/lolapi/endpoints"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestGetSummoner(t *testing.T) {
	conf, err := lolapi.GetConfig()
	assert.Nil(t, err)
	client := lolapi.ClientFromConfig(conf, lolapi.RegionEuropeWest)
	client.Region = lolapi.RegionEuropeWest

	resp, _, err := client.Call(endpoints.GetSummonerByName, "good bye iceland")
	fmt.Println(string(resp))
	assert.Nil(t, err)

	json := jsoniter.ConfigFastest
	var summoner data.SummonerDTO
	json.Unmarshal(resp, &summoner)
	fmt.Printf("%+#v\n", summoner)

	assert.Equal(t, "good bye iceland", summoner.Name)
	assert.Equal(t, "2UVIWB-VKOlq06KZNuQYBO1cmFvNNfl5Brh4VoX9eeZXIAHRsoGPqRBL", summoner.AccountID)
	assert.Equal(t, "6q1SAoDhU67-YJkUgrbtjzzRK_h3JMmPi00oStvW-YEKqDFgyzyslj1QEA", summoner.ID)
	assert.Equal(t, "5FqT0QB2yWgokyLCUGFXgBJm7zdFYdtmlnDfQmgLAO9kF8VICwn60vr8Bi7WnJVI97lVGPYG5MV7uw", summoner.PUUID)
}
