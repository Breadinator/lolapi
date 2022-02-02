package lolapi

import (
	"fmt"
	"time"

	"github.com/breadinator/lolapi/errors"
)

type Client struct {
	Key    string
	Limit  uint16 // The number of requests per minute
	Region Region

	lastCalled time.Time // The UnixNano timestamp of the next allowed request
}

func ClientFromConfig(config *Config, region Region) (client Client) {
	client.Key = config.APIKey
	client.Limit = config.Limit
	client.Region = region
	return
}

// Makes a basic request to see if the key is valid
func (c *Client) CheckKey() bool {
	return false
}

// Checks whether another request can be made
func (c *Client) CheckLimit() bool {
	if c.Limit == 0 { // No limit
		return true
	}
	if time.Since(c.lastCalled).Seconds() >= float64(60)/float64(c.Limit) {
		return true
	}

	return false
}

func (c *Client) Call(e Endpoint, args ...string) ([]byte, int, error) {
	if !c.CheckLimit() {
		return make([]byte, 0), -1, errors.RaiseDataLimitExceeded(fmt.Sprintf("Limit of %d/min exceeded, please wait.", c.Limit))
	}
	defer func() {
		c.lastCalled = time.Now()
	}()
	return e.Call(c, args...)
}

func (c *Client) AwaitNoLimit() {
	waitTime := float64(60) / float64(c.Limit) // minimum time between calls
	time.Sleep(time.Duration(waitTime - time.Since(c.lastCalled).Seconds() + 1))
}
