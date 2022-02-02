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

func TestGetChampionsAll(t *testing.T) {
	conf, err := lolapi.GetConfig()
	assert.Nil(t, err)
	client := lolapi.ClientFromConfig(conf, lolapi.RegionEuropeWest)
	client.Limit = 0

	resp, stat, err := client.Call(endpoints.GetChampionsAll, "12.3.1", "en_US")
	assert.Nil(t, err)
	assert.Equal(t, stat, 200)

	champs := new(data.Champions)
	err = jsoniter.Unmarshal(resp, champs)
	assert.Nil(t, err)

	fmt.Printf("%+#v\n", champs)
}

func TestGetChampionSpecific(t *testing.T) {
	conf, err := lolapi.GetConfig()
	assert.Nil(t, err)
	client := lolapi.ClientFromConfig(conf, lolapi.RegionEuropeWest)
	client.Limit = 0

	resp, stat, err := client.Call(endpoints.GetChampionsSpecific, "12.3.1", "en_US", "Janna")
	assert.Nil(t, err)
	assert.Equal(t, stat, 200)

	c := new(data.Champions)
	err = jsoniter.Unmarshal(resp, c)
	assert.Nil(t, err)

	// Overall response data
	assert.Equal(t, "champion", c.Type)
	assert.Equal(t, "standAloneComplex", c.Format)
	assert.Equal(t, "12.3.1", c.Version)

	// Janna data
	assert.Equal(t, "Janna", c.Data["Janna"].ID)
	assert.Equal(t, "40", c.Data["Janna"].Key)
	assert.Equal(t, "Janna", c.Data["Janna"].Name)
	assert.Equal(t, "Janna.png", c.Data["Janna"].Image.Full)
	assert.Equal(t, 13, len(c.Data["Janna"].Skins))
	assert.Equal(t, "Armed with the power of Runeterra's gales, Janna is a mysterious, elemental wind spirit who protects the dispossessed of Zaun. Some believe she was brought into existence by the pleas of Runeterra's sailors who prayed for fair winds as they navigated treacherous waters and braved rough tempests. Her favor and protection has since been called into the depths of Zaun, where Janna has become a beacon of hope to those in need. No one knows where or when she will appear, but more often than not, she's come to help.", c.Data["Janna"].Lore)
	assert.Equal(t, "Armed with the power of Runeterra's gales, Janna is a mysterious, elemental wind spirit who protects the dispossessed of Zaun. Some believe she was brought into existence by the pleas of Runeterra's sailors who prayed for fair winds as they navigated...", c.Data["Janna"].Blurb)
	assert.Equal(t, 3, len(c.Data["Janna"].AllyTips))
	assert.Equal(t, 3, len(c.Data["Janna"].EnemyTips))
	assert.Equal(t, "Mage", c.Data["Janna"].Tags[1])
	assert.Equal(t, "Mana", c.Data["Janna"].Partype)
	assert.Equal(t, 7, c.Data["Janna"].Info.Magic)
	assert.EqualValues(t, 0.55, c.Data["Janna"].Stats.HPRegenPerLevel)
	assert.Equal(t, "Cooldown", c.Data["Janna"].Spells[2].LevelTip.Label[2])
	assert.Equal(t, "passive", c.Data["Janna"].Passive.Image.Group)
}
