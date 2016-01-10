package dao

import (
	"gopkg.in/redis.v3"
	"keyservice"
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

// returns an interface from cache or database; if item is not found, a nil is returned
func (ds *DataSource) Get(key string) (interface{}, error) {
	value := ds.cache.Get(key)

	log.Info("get: %s=%v", key, value)

	if value == nil && ds.client != nil {
		log.Info("not cached, find from database: %s", key)
		return ds.client.Get(key).Result()
	}

	return value, nil
}

// sets cache and database; returns error or nil on no errors
func (ds *DataSource) Set(key string, value keyservice.DataModelType) error {
	ds.cache.Set(key, value)

	log.Info("set: %s=%v", key, value)
	if ds.client != nil {
		json, err := value.ToJSON()
		if err != nil {
			return err
		}

		log.Info("set db: %s", json)

		err = ds.client.Set(key, string(json), 0).Err()

		return err
	}

	return nil
}

// delete from cache and database and return the item (if in cache)
func (ds *DataSource) Delete(key string) interface{} {
	value := ds.cache.Delete(key)

	// TODO : remove from redis
	if ds.client != nil {
		ds.client.Del(key)
	}

	return value
}

func (ds *DataSource) GetCache() *DataModelCache {
	return ds.cache
}

func (ds *DataSource) GetCacheLen() int {
	return ds.cache.Len()
}
