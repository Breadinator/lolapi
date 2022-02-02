package lolapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Endpoint struct {
	Method   string // e.g http.MethodGet or just "GET"
	Path     string // e.g. /riot/account/v1/accounts/me
	UseRoute bool   // If you use a Route instead of a Region
	RawPath  bool   // If true, instead of using api.riotgames.com it uses the path directly from `Path`, but still adds variables.
}

func (e *Endpoint) Call(client *Client, args ...string) ([]byte, int, error) {
	httpClient := &http.Client{}

	var url string
	if e.UseRoute && !e.RawPath {
		url = string(RouteFromRegion(client.Region))
	} else if !e.RawPath {
		url = fmt.Sprintf("https://%s.api.riotgames.com", client.Region)
	}
	url += e.Path

	for i, arg := range args {
		url = strings.ReplaceAll(url, "$"+strconv.Itoa(i+1), arg)
	}
	re, err := regexp.Compile(`\$\d`)
	if err != nil {
		return make([]byte, 0), -1, err
	}
	url = re.ReplaceAllString(url, "")

	req, err := http.NewRequest(e.Method, url, nil)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Add("X-Riot-Token", client.Key)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	return b, resp.StatusCode, err
}
