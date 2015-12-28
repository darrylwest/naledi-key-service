package dao

import (
    "gopkg.in/redis.v3"
    // "github.com/darrylwest/cassava-logger/logger"
)

type DataSource struct {
    client *redis.Client
    cache Cache
}

// create a new datasource with optional recis client; cache is always created...
func NewDataSource(client *redis.Client) DataSource {
    ds := DataSource{}

    ds.client = client
    ds.cache = NewCache() // NewNullCache

    return ds
}

func NewCachedDataSource(client *redis.Client) DataSource {
    ds := DataSource{}

    ds.client = client
    ds.cache = NewCache()

    return ds
}

func (ds *DataSource) Get(key string) (interface{}, error) {
    var value interface{}

    value = ds.cache.Get(key)

    // if value == nil, try to pull from redis
    log.Info( "get: %s=%v", key, value )

    return value, nil
}

func (ds *DataSource) Set(key, value string) error {
    ds.cache.Set(key, value)

    log.Info("set: %s=%s", key, value)

    return nil
}

func (ds *DataSource) Delete(key string) interface{} {
    var value interface{}

    value = ds.cache.Delete( key )

    // TODO : remove from redis

    return value
}

func (ds *DataSource) GetCache() Cache {
    return ds.cache
}

func (ds *DataSource) GetCacheLen() int {
    return ds.cache.Len()
}
