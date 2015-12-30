package dao

import (
	"keyservice"
	// "gopkg.in/redis.v3"
	"github.com/darrylwest/cassava-logger/logger"
)

const (
	NotImplementedYet = "%s not implemented yet"
	NotFound          = "%s not found for %s"
)

var (
	log *logger.Logger
	ctx *keyservice.Context
)


func InitializeDao(context *keyservice.Context, logger *logger.Logger) {
	ctx = context
	log = logger
}

// func getPrimaryClient() *redis.Client
// func getSecondaryClient() *.redis.Client
