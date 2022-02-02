package tests

import (
	"fmt"
	"testing"

	"github.com/breadinator/lolapi/lolapi/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestLatestVersion(t *testing.T) {
	ver, err := endpoints.GetLatestVersion()
	assert.Nil(t, err)
	assert.False(t, ver == "")
	fmt.Println(ver)
}
