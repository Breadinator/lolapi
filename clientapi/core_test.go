package clientapi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEndpoints(t *testing.T) {
	endpoints, err := GetEndPoints()
	assert.Nil(t, err)
	assert.False(t, len(endpoints) == 0)
	fmt.Println(endpoints)
}

func TestGetPem(t *testing.T) {
	assert.NotNil(t, certs)
}
