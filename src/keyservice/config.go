package keyservice

import (
	"encoding/json"
	"encoding/hex"
	"gopkg.in/redis.v3"
	"io/ioutil"
)

type Config struct {
	name    string
	appkey  string
	baseURI string

	primaryRedisOptions   *redis.Options
	secondaryRedisOptions *redis.Options

	privateLocalKey *[]byte
}

func (c Config) ToMap() map[string]interface{} {
	hash := make(map[string]interface{})

	hash["name"] = c.name
	hash["appkey"] = c.appkey
	hash["baseURI"] = c.baseURI

	hash["primaryRedisOptions"] = c.primaryRedisOptions
	hash["secondaryRedisOptions"] = c.secondaryRedisOptions

	pk := make([]byte, KeySize)
	copy(pk[:], *c.privateLocalKey)
	hash["privateLocalKey"] = pk

	return hash
}

func ReadConfig(path string) (*Config, error) {
	log.Info("read the configuration file: %s", path)

	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Error("config file read error: ", err)
		return nil, err
	}

	return ParseConfig(data)
}

func ParseRedisOptions(hash map[string]interface{}) *redis.Options {
	opts := new(redis.Options)

	opts.Addr = hash["addr"].(string)
	opts.Password = hash["password"].(string)
	opts.DB = int64(hash["db"].(float64))

	return opts
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

	config.primaryRedisOptions = ParseRedisOptions(hash["primaryRedisOptions"].(map[string]interface{}))
	config.secondaryRedisOptions = ParseRedisOptions(hash["secondaryRedisOptions"].(map[string]interface{}))

	if key, ok := hash["privateLocalKey"].(string); ok == true {
		log.Debug("key: %s", key)

		decoded, err := hex.DecodeString(key)
		if err != nil {
			log.Error("error decoding private local key")
		} else {
			config.privateLocalKey = &decoded
			log.Debug("private local key: %v", config.privateLocalKey)
		}

	}

	return config, nil
}
