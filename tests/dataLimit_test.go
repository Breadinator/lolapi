package tests

import (
	"testing"

	"github.com/breadinator/lolapi/errors"
	"github.com/breadinator/lolapi/lolapi"
	"github.com/breadinator/lolapi/lolapi/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestExceedDataLimit(t *testing.T) {
	conf, err := lolapi.GetConfig()
	assert.Nil(t, err)

	client := lolapi.ClientFromConfig(conf, lolapi.RegionEuropeWest)
	client.Limit = 1 // 1 request per minute

	client.Call(endpoints.GetLeagueChallenger, string(lolapi.QueueLolRankedSolo))
	_, _, err = client.Call(endpoints.GetLeagueMaster, string(lolapi.QueueLolRankedSolo))

	_, ok := err.(errors.ErrDataLimitExceeded)
	assert.True(t, err != nil)
	assert.True(t, ok)
}
