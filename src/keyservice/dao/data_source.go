package dao

import (
	"gopkg.in/redis.v3"
	"keyservice/models"
	// "github.com/darrylwest/cassava-logger/logger"
)

type DataSource struct {
	client *redis.Client
	cache  *DataModelCache
}

// create a new datasource with optional recis client; cache is always created...
func NewDataSource(client *redis.Client) DataSource {
	ds := DataSource{}

	ds.client = client

	// make this a no-op cache
	ds.cache = NewDataModelCache()

	return ds
}

func NewCachedDataSource(client *redis.Client) DataSource {
	ds := DataSource{}

	ds.client = client
	ds.cache = NewDataModelCache()

	return ds
}

func (ds *DataSource) Get(key string) (models.DataModelType, error) {
	value := ds.cache.Get(key)

	// if value == nil, try to pull from redis
	// if value found in redis, then push to cache
	log.Info("get: %s=%v", key, value)

	return value, nil
}

func (ds *DataSource) Set(key string, value models.DataModelType) error {
	ds.cache.Set(key, value)

	log.Info("set: %s=%v", key, value)
	// use go routine or queue to save data to redis

	return nil
}

func (ds *DataSource) Delete(key string) interface{} {
	value := ds.cache.Delete(key)

	// TODO : remove from redis

	return value
}

func (ds *DataSource) GetCache() *DataModelCache {
	return ds.cache
}

func (ds *DataSource) GetCacheLen() int {
	return ds.cache.Len()
}
