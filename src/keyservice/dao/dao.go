package dao

import (
	"keyservice"
	"gopkg.in/redis.v3"
	"github.com/darrylwest/cassava-logger/logger"
)

const (
	NotImplementedYet = "%s not implemented yet"
	NotFound          = "%s not found for %s"
)

var (
	log *logger.Logger
	ctx *keyservice.Context

	primaryClient *redis.Client
	secondaryClient *redis.Client
)


func InitializeDao(context *keyservice.Context, logger *logger.Logger) {
	ctx = context
	log = logger
}

func GetPrimaryClient() *redis.Client {
	if primaryClient == nil {
		conf := ctx.GetConfig()
		opts := conf.GetPrimaryRedisOptions()
		log.Info("create the primary client: %s", opts)
		primaryClient = redis.NewClient(opts)

		pong, err := primaryClient.Ping().Result()
		if err != nil {
			panic( err )
		}

		log.Info("ping->%s", pong)
	}

	return primaryClient
}

func GetSecondaryClient() *redis.Client {
	if secondaryClient == nil {
		conf := ctx.GetConfig()
		opts := conf.GetSecondaryRedisOptions()
		log.Info("create the secondary client: %s", opts)
		secondaryClient = redis.NewClient(opts)

		pong, err := secondaryClient.Ping().Result()
		if err != nil {
			panic( err )
		}

		log.Info("ping->%s", pong)
	}

	return secondaryClient
}
