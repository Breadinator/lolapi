package clientapi

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

const Root = "https://127.0.0.1:2999"

var certs, _ = GetPEM()

func Get(endpoint string) (*http.Response, error) {
	if certs == nil {
		c, err := GetPEM()
		certs = c
		if err != nil {
			return nil, err
		}
	}

	client := new(http.Client)
	client.Transport = &http.Transport{
		//TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // allow insecure connections
		TLSClientConfig: &tls.Config{RootCAs: certs},
	}
	return client.Get(Root + endpoint)
}

func GetRead(endpoint string) ([]byte, error) {
	resp, err := Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func GetPEM() (*x509.CertPool, error) {
	const url = "https://static.developer.riotgames.com/docs/lol/riotgames.pem"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pemData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	certs := x509.NewCertPool()
	if ok := certs.AppendCertsFromPEM(pemData); !ok {
		return certs, errors.New("couldn't append")
	}
	return certs, nil
}

func GetEndPoints() ([]string, error) {
	endpoints := make([]string, 30)
	reqData := new(struct {
		Paths map[string]interface{}
	})

	resp, err := Get("/swagger/v3/openapi.json")
	if err != nil {
		return endpoints, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return endpoints, errors.New(resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return endpoints, err
	}

	if err := jsoniter.ConfigFastest.Unmarshal(data, &reqData); err != nil {
		return endpoints, err
	}

	for key := range reqData.Paths {
		endpoints = append(endpoints, key)
	}

	return endpoints, nil
}
