package dao

import (
	"keyservice"
	// "gopkg.in/redis.v3"
	"github.com/darrylwest/cassava-logger/logger"
)

var (
	log *logger.Logger
	ctx *keyservice.Context
)

func InitializeDao(context *keyservice.Context, logger *logger.Logger) {
	ctx = context
	log = logger
}
