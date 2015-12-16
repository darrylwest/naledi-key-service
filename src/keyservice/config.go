package keyservice

import (
	"gopkg.in/redis.v3"
    "errors"
)

type Config struct {
	appkey  string
	baseURI string

	primaryRedis   *redis.Options
	secondaryRedis *redis.Options
}

func ReadConfig(file string) (*Config, error) {
	log.Info("read the configuration file: ", file)

    return nil, errors.New("not implemented yet")
}
