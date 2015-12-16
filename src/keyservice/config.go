package keyservice

import (
	"gopkg.in/redis.v3"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	name    string
	appkey  string
	baseURI string

	primaryRedis   *redis.Options
	secondaryRedis *redis.Options
}

func ReadConfig(path string) (*Config, error) {
	log.Info("read the configuration file: ", path)

	data, err := ioutil.ReadFile( path )

	if err != nil {
		log.Error("config file read error: ", err)
		return nil, err
	}

	return ParseConfig(data)
}

func ParseConfig(data []byte) (*Config, error) {
	var hash map[string]interface{}

	if err := json.Unmarshal(data, &hash); err != nil {
		log.Error("parse error: ", err)
		return nil, err
	}

	config := new(Config)

	config.name = hash["name"].(string)
	config.appkey = hash["appkey"].(string)
	config.baseURI = hash["baseURI"].(string)

	

	return config, nil
}
