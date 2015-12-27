package dao

import (
    "keyservice"
	"gopkg.in/redis.v3"
)

type UserDao struct {
    redisOptions **redis.Options
}

// the context contains keys to access primary and secondary data sources
func CreateUserDao(ctx *keyservice.Context) *UserDao {
    dao := new(UserDao)

    return dao
}
