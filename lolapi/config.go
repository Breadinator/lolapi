package lolapi

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
)

type Config struct {
	APIKey string
	Limit  uint16
}

func GetConfig() (*Config, error) {
	p, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	c := new(Config)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	exists := CheckPathExists(p)

	if exists {
		contents, err := ioutil.ReadFile(p)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(contents, c)
	} else {
		GenerateNewConfig(c)
		err = WriteNewConfig(c, p, json)
	}

	return c, err
}

func GetConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Abs(path.Join(dir, "gololapi-config.json"))
}

func CheckPathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func GenerateNewConfig(c *Config) {
	*c = *new(Config)
	c.APIKey = "Insert API key here"
	c.Limit = 30
}

func WriteNewConfig(c *Config, path string, json jsoniter.API) error {
	marsh, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(marsh)
	return err

}
