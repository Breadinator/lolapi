package tests

import (
	"testing"

	"github.com/breadinator/lolapi/lolapi/data"
	"github.com/stretchr/testify/assert"
)

func TestGetItemsParsed(t *testing.T) {
	items, err := data.GetItemsParsed("12.3.1", "en_US")
	assert.Nil(t, err)

	// Top level keys
	assert.Equal(t, "item", items.Type)
	assert.Equal(t, "12.3.1", items.Version)
	assert.Equal(t, 206, len(items.Data))
	assert.Equal(t, 122, len(items.Groups))
	assert.Equal(t, 7, len(items.Tree))

	// Test specific item
	assert.Equal(t, "Wit's End", items.Data[3091].Name)
	assert.Equal(t, "<mainText><stats><attention>40</attention> Attack Damage<br><attention>40%</attention> Attack Speed<br><attention>40</attention> Magic Resist</stats><br><li><passive>Fray:</passive> Attacks apply <magicDamage>15 - 80 (based on level) magic damage</magicDamage> <OnHit>On-Hit</OnHit> and grant <speed>20 Move Speed</speed> for 2 seconds.</mainText><br>", items.Data[3091].Description)
	assert.Equal(t, ";", items.Data[3091].Colloq)
	assert.Equal(t, "Resist magic damage and claw your way back to life.", items.Data[3091].PlainText)
	assert.Contains(t, items.Data[3091].From, "3051")
	assert.Contains(t, items.Data[3091].From, "1033")
	assert.Contains(t, items.Data[3091].From, "1037")
	assert.Equal(t, "3091.png", items.Data[3091].Image.Full)
	assert.True(t, items.Data[3091].Gold.Purchasable)
	assert.Contains(t, items.Data[3091].Tags, "OnHit")
	assert.Equal(t, 4, len(items.Data[3091].Maps))
	assert.Equal(t, 3, len(items.Data[3091].Stats))
	assert.Equal(t, 3, items.Data[3091].Depth)
}
